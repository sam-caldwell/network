package core

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"sync/atomic"
	"syscall"
	"time"
)

// ExecuteIter - Execute the request against the given sockType.
func (req *NetlinkRequest) ExecuteIter(socketType int, resType uint16, iterFunction IteratorFunctionPtr) error {
	var (
		err             error
		s               *NetlinkSocket
		SocketTimeoutTv = unix.Timeval{
			Sec:  5,
			Usec: 0,
		}
	)

	// Helper function to check for temporary errors
	isTemporaryError := func(err error) bool {
		if err == nil {
			return false
		}
		var errno syscall.Errno
		if errors.As(err, &errno) {
			if errno == syscall.EAGAIN || errno == syscall.EWOULDBLOCK || errno == syscall.EINTR {
				return true
			}
		}
		return false
	}

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
		// Prepare file descriptor set
		var readFds unix.FdSet
		fd := s.GetFd()
		fdSet(fd, &readFds)

		// Set timeout
		timeout := &unix.Timeval{Sec: 5, Usec: 0} // Set a 5-second timeout

		// Call select()
		ret, err := unix.Select(int(fd+1), &readFds, nil, nil, timeout)
		if err != nil {
			log.Printf("select() error: %v\n", err)
			if isTemporaryError(err) {
				backoff := baseDelay * time.Duration(1<<retries) // Exponential backoff
				log.Printf("Retrying select() (%d/%d) in %v\n", retries+1, maxRetries, backoff)
				time.Sleep(backoff)
				continue
			}
			return err
		}

		if ret == 0 {
			log.Printf("select() timed out, no data available\n")
			continue // Timeout occurred, retry if possible
		}

		var (
			from     *unix.SockaddrNetlink
			messages []NetlinkMessage
		)

		if messages, from, err = s.Receive(); err != nil {
			log.Printf("s.Receive() error: %v\n", err)
			if isTemporaryError(err) {
				backoff := baseDelay * time.Duration(1<<retries) // Exponential backoff
				log.Printf("Retrying receive (%d/%d) in %v\n", retries+1, maxRetries, backoff)
				time.Sleep(backoff)
				continue
			}
			return err
		}

		log.Printf("s.Receive() success\n"+
			"  messages: %v\n"+
			"      from: %v", messages, from)

		if from.Pid != NetlinkPortId {
			return fmt.Errorf("wrong sender portid %d, expected %d", from.Pid, NetlinkPortId)
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
			if m.Header.Flags&NlmFDumpIntr != 0 {
				return EINTR
			}
			if m.Header.Type == NlmsgDone || m.Header.Type == NlmsgError {
				// NLMSG_DONE might have no payload, assume no error.
				if m.Header.Type == NlmsgDone && len(m.Data) == 0 {
					return nil
				}

				if errno := int32(NativeEndian.Uint32(m.Data[0:4])); errno == 0 {
					return nil
				} else {
					err = Errors(-errno)
				}

				return err
			}
			if resType != 0 && m.Header.Type != resType {
				continue
			}
			if cont := iterFunction(m.Data); !cont {
				iterFunction = nullExecuteIterFunc
			}
			if m.Header.Flags&NlmFMulti == 0 {
				return nil
			}
		}
	}
	return fmt.Errorf("max retries reached")
}

// Helper function to set the file descriptor in the FdSet
func fdSet(fd int, set *unix.FdSet) {
	set.Bits[fd/64] |= 1 << (uint(fd) % 64)
}
