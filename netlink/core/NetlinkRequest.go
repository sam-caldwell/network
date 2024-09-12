package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
	"sync/atomic"
)

// NetlinkRequest - Represents a netlink request message that can be sent via netlink sockets.
//
// This structure is used to build and send netlink messages to the kernel. It wraps the netlink message header (NlMsghdr)
// and includes the data payload (Data), raw data (RawData), and associated socket handles (Sockets).
//
// References:
// - Linux kernel netlink source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
// - Netlink man page: https://man7.org/linux/man-pages/man7/netlink.7.html
type NetlinkRequest struct {

	// NlMsghdr - Represents the netlink message header (struct nlmsghdr).
	//
	// This header defines the type, flags, sequence number, and PID for the message.
	//
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L101-L114
	unix.NlMsghdr

	// Data - The payload data to include in the message.
	//
	// This field represents the body or payload of the netlink message. It typically contains
	// information such as network device attributes, routing information, etc.
	//
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L85
	Data []NetlinkRequestData

	// RawData - A byte slice representing raw data that can be added to the message.
	//
	// This field allows you to include additional data that may not conform to standard netlink
	// message structures, giving flexibility for custom implementations.
	//
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L85
	RawData []byte

	// Sockets - A map of IP protocol to socket handles.
	//
	// This field holds a mapping of IP protocol types (e.g., TCP, UDP) to corresponding socket
	// handles used for communication. Socket handles facilitate interactions with the kernel over netlink.
	//
	// Reference: https://man7.org/linux/man-pages/man7/netlink.7.html
	Sockets map[IpProtocol]*SocketHandle
}

// AddData - append the NetlinkRequest data with the given NetlinkRequestData
func (req *NetlinkRequest) AddData(data NetlinkRequestData) {

	req.Data = append(req.Data, data)

}

// AddRawData - add raw bytes to the end of the NetlinkRequest object during serialization
func (req *NetlinkRequest) AddRawData(data []byte) {
	req.RawData = append(req.RawData, data...)
}

// Execute - Execute the request against the given sockType. Returns a list of netlink messages in serialized format,
// optionally filtered by resType.
func (req *NetlinkRequest) Execute(sockType IpProtocol, resType uint16) ([][]byte, error) {
	var res [][]byte
	err := req.ExecuteIter(sockType, resType, func(msg []byte) bool {
		res = append(res, msg)
		return true
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ExecuteIter - execute the request against the given sockType. Calls the provided callback func once for each
// netlink message. If the callback returns false, it is not called again, but the remaining messages are
// consumed/discarded.
//
// Thread safety:
//
//	ExecuteIter holds a lock on the socket until it finishes iteration so the callback must
//	not call back into the netlink API.
func (req *NetlinkRequest) ExecuteIter(sockType IpProtocol, resType uint16, iterFunction func(msg []byte) bool) error {
	var (
		s   *NetlinkSocket
		err error
	)

	if req.Sockets != nil {
		if sh, ok := req.Sockets[sockType]; ok {
			s = sh.Socket
			req.Seq = atomic.AddUint32(&sh.Seq, 1)
		}
	}
	sharedSocket := s != nil

	if s == nil {
		s, err = getNetlinkSocket(sockType)
		if err != nil {
			return err
		}

		if err = s.SetSendTimeout(&SocketTimeoutTv); err != nil {
			return err
		}
		if err = s.SetReceiveTimeout(&SocketTimeoutTv); err != nil {
			return err
		}
		if enableErrorMessageReporting {
			if err = s.SetExtAck(true); err != nil {
				return err
			}
		}

		defer func() { _ = s.Close() }()
	} else {
		s.Lock()
		defer s.Unlock()
	}

	if err = s.Send(req); err != nil {
		return err
	}

	var pid uint32
	if pid, err = s.GetPid(); err != nil {
		return err
	}

done:
	for {
		var (
			from     *unix.SockaddrNetlink
			messages []NetlinkMessage
		)
		if messages, from, err = s.Receive(); err != nil {
			return err
		}
		if from.Pid != PortIdKernel {
			return fmt.Errorf("wrong sender portid %d, expected %d", from.Pid, PortIdKernel)
		}
		for _, m := range messages {
			if m.Header.Seq != req.Seq {
				if sharedSocket {
					continue
				}
				return fmt.Errorf("wrong Seq nr %d, expected %d", m.Header.Seq, req.Seq)
			}
			if m.Header.Pid != pid {
				continue
			}
			if m.Header.Flags&unix.NLM_F_DUMP_INTR != 0 {
				return unix.EINTR
			}
			if m.Header.Type == unix.NLMSG_DONE || m.Header.Type == unix.NLMSG_ERROR {
				// NLMSG_DONE might have no payload, if so assume no error.
				if m.Header.Type == unix.NLMSG_DONE && len(m.Data) == 0 {
					break done
				}

				if errno := int32(NativeEndian.Uint32(m.Data[0:4])); errno == 0 {
					break done
				} else {
					err = unix.Errno(-errno)
				}
				unreadData := m.Data[4:]
				if m.Header.Flags&unix.NLM_F_ACK_TLVS != 0 && len(unreadData) > SizeOfNlMsgHdr {
					// Skip the echoed request message.
					var echoReqH *unix.NlMsghdr
					if echoReqH, err = DeserializeNlMsgHdr(unreadData); err != nil {
						return err
					}
					unreadData = unreadData[nlmAlignOf(int(echoReqH.Len)):]

					// Annotate `err` using nlmsgerr attributes.
					for len(unreadData) >= SizeOfUnixRtAttr {
						var attr *unix.RtAttr
						if attr, err = DeserializeUnixRtAttr(unreadData); err != nil {
							return err
						}
						attrData := unreadData[SizeOfUnixRtAttr:attr.Len]

						switch v := NlmsgerrAttr(attr.Type); v {
						case NlmsgerrAttrMsg:
							err = fmt.Errorf("%w: %s", err, unix.ByteSliceToString(attrData))
						default:
							// TODO: handle other NLMSGERR_ATTR types
						}

						unreadData = unreadData[rtaAlignOf(int(attr.Len)):]
					}
				}

				return err
			}
			if resType != 0 && m.Header.Type != resType {
				continue
			}
			if cont := iterFunction(m.Data); !cont {
				//
				// switch to the nullExecuteIterFunc() which will allow us to drain the remaining messages from
				// the kernel without passing them to the real iterFunction().
				//
				iterFunction = nullExecuteIterFunc
			}
			if m.Header.Flags&unix.NLM_F_MULTI == 0 {
				break done
			}
		}
	}
	return nil
}

// Serialize outputs a serialized []byte from the NetlinkRequest struct.
func (req *NetlinkRequest) Serialize() (out []byte, err error) {
	// Calculate the total length of the netlink message.
	length := SizeOfNlMsgHdr
	for _, data := range req.Data {
		s, err := data.Serialize()
		if err != nil {
			return nil, err
		}
		length += len(s)
	}
	length += len(req.RawData)

	// Update the message length in the header.
	req.Len = uint32(length)

	// Create a buffer to hold the serialized data.
	buf := bytes.NewBuffer(make([]byte, 0, length))

	// Serialize the netlink message header.
	if err = binary.Write(buf, NativeEndian, req.Len); err != nil {
		return nil, err
	}
	if err = binary.Write(buf, NativeEndian, req.Type); err != nil {
		return nil, err
	}
	if err = binary.Write(buf, NativeEndian, req.Flags); err != nil {
		return nil, err
	}
	if err = binary.Write(buf, NativeEndian, req.Seq); err != nil {
		return nil, err
	}
	if err = binary.Write(buf, NativeEndian, req.Pid); err != nil {
		return nil, err
	}

	// Serialize the payload data.
	for _, data := range req.Data {
		s, err := data.Serialize()
		if err != nil {
			return nil, err
		}
		buf.Write(s)
	}

	// Add the raw data, if any.
	if len(req.RawData) > 0 {
		buf.Write(req.RawData)
	}

	return buf.Bytes(), nil
}

// NetlinkRequestData - interface that abstracts the data included in a netlink request, allowing different types of
// data to be encapsulated and sent as part of a netlink message.
type NetlinkRequestData interface {
	Len() int
	Serialize() ([]byte, error)
}
