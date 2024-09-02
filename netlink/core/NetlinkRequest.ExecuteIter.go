package core

import (
	"fmt"
	"golang.org/x/sys/unix"
	"sync/atomic"
	"syscall"
	"unsafe"
)

// nullExecuteIterFunc - a noop function we will use when draining the queue
func nullExecuteIterFunc(msg []byte) bool {
	return true
}

// ExecuteIter - execute the request against the given sockType. Calls the provided callback func once for each
// netlink message. If the callback returns false, it is not called again, but the remaining messages are
// consumed/discarded.
//
// Thread safety: ExecuteIter holds a lock on the socket until
// it finishes iteration so the callback must not call back into
// the netlink API.
func (req *NetlinkRequest) ExecuteIter(sockType int, resType uint16, iterFunction func(msg []byte) bool) error {
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

		if err := s.SetSendTimeout(&SocketTimeoutTv); err != nil {
			return err
		}
		if err := s.SetReceiveTimeout(&SocketTimeoutTv); err != nil {
			return err
		}
		if enableErrorMessageReporting {
			if err := s.SetExtAck(true); err != nil {
				return err
			}
		}

		defer s.Close()
	} else {
		s.Lock()
		defer s.Unlock()
	}

	if err := s.Send(req); err != nil {
		return err
	}

	pid, err := s.GetPid()
	if err != nil {
		return err
	}

done:
	for {
		messages, from, err := s.Receive()
		if err != nil {
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
				return syscall.Errno(unix.EINTR)
			}

			if m.Header.Type == unix.NLMSG_DONE || m.Header.Type == unix.NLMSG_ERROR {
				// NLMSG_DONE might have no payload, if so assume no error.
				if m.Header.Type == unix.NLMSG_DONE && len(m.Data) == 0 {
					break done
				}

				errno := int32(nativeEndian.Uint32(m.Data[0:4]))
				if errno == 0 {
					break done
				}
				var err error
				err = syscall.Errno(-errno)

				unreadData := m.Data[4:]
				if m.Header.Flags&unix.NLM_F_ACK_TLVS != 0 && len(unreadData) > syscall.SizeofNlMsghdr {
					// Skip the echoed request message.
					echoReqH := (*syscall.NlMsghdr)(unsafe.Pointer(&unreadData[0]))
					unreadData = unreadData[nlmAlignOf(int(echoReqH.Len)):]

					// Annotate `err` using nlmsgerr attributes.
					for len(unreadData) >= syscall.SizeofRtAttr {
						attr := (*syscall.RtAttr)(unsafe.Pointer(&unreadData[0]))
						attrData := unreadData[syscall.SizeofRtAttr:attr.Len]

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
				// switch to the nullExecuteIterFunc() which will allow us to drain the remaining messages from
				// the kernel without passing them to the real iterFunction().
				iterFunction = nullExecuteIterFunc
			}
			if m.Header.Flags&unix.NLM_F_MULTI == 0 {
				break done
			}
		}
	}
	return nil
}
