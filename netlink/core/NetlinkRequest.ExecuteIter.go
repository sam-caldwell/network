package core

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"sync/atomic"
	"time"
)

// ExecuteIter - execute the request against the given sockType. Calls the provided callback func once for each
// netlink message. If the callback returns false, it is not called again, but the remaining messages are
// consumed/discarded.
//
// Thread safety:
//
//	ExecuteIter holds lock on socket until it finishes iteration. The callback must	not call back into the netlink API.
func (req *NetlinkRequest) ExecuteIter(socketType int, resType uint16, iterFunction func(msg []byte) bool) error {
	var (
		s   *NetlinkSocket
		err error
	)

	// If NetlinkRequest.Socket map not nil, look up the socket type
	if req.Sockets != nil {
		if sh, ok := req.Sockets[IpProtocol(socketType)]; ok {
			s = sh.Socket
			req.Seq = atomic.AddUint32(&sh.Seq, 1)
		}
	}

	sharedSocket := s != nil
	if sharedSocket {
		log.Println("sharedSocket: true")
		s.Lock()
		defer s.Unlock()
	} else {
		log.Println("sharedSocket: false")
		s, err = GetNetlinkSocket(IpProtocol(socketType))
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
	}

	log.Println("sending")
	if err = s.Send(req); err != nil {
		return err
	}
	log.Println("send with no error")
	var pid uint32
	if pid, err = s.GetPid(); err != nil {
		return err
	}
	log.Printf("pid: %v", pid)

	const maxRetries = 3
	const baseDelay = 50 * time.Millisecond

	for retries := 0; retries < maxRetries; retries++ {
		var (
			from     *unix.SockaddrNetlink
			messages []NetlinkMessage
		)

		if messages, from, err = s.Receive(); err != nil {
			if errors.Is(err, unix.EAGAIN) || errors.Is(err, unix.EWOULDBLOCK) || errors.Is(err, unix.EINTR) {
				backoff := baseDelay * time.Duration(1<<retries) // Exponential backoff
				log.Printf("s.Receive() retry (%d/%d) due to: %v. Retrying in %v\n", retries+1, maxRetries, err, backoff)
				time.Sleep(backoff)
				continue
			}
			log.Printf("s.Receive() error: %v\n", err)
			return err
		}
		log.Printf("s.Receive() success\n"+
			"  messages: %v\n"+
			"      from: %v", messages, from)

		if from.Pid != PortIdKernel {
			return fmt.Errorf("wrong sender portid %d, expected %d", from.Pid, PortIdKernel)
		}

		// Handle each message
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
				// NLMSG_DONE might have no payload, assume no error.
				if m.Header.Type == unix.NLMSG_DONE && len(m.Data) == 0 {
					return nil
				}

				if errno := int32(NativeEndian.Uint32(m.Data[0:4])); errno == 0 {
					return nil
				} else {
					err = unix.Errno(-errno)
				}

				return err
			}
			if resType != 0 && m.Header.Type != resType {
				continue
			}
			if cont := iterFunction(m.Data); !cont {
				iterFunction = nullExecuteIterFunc
			}
			if m.Header.Flags&unix.NLM_F_MULTI == 0 {
				return nil
			}
		}
	}
	return fmt.Errorf("max retries reached")
}
