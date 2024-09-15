package core

import (
	"errors"
	"golang.org/x/sys/unix"
	"sync/atomic"
)

// Receive - read a message from the netlink socket and returns the parsed NetlinkMessage slice, the sender's
// SockaddrNetlink, and an error if any occurred during the process.
func (s *NetlinkSocket) Receive() ([]NetlinkMessage, *unix.SockaddrNetlink, error) {

	var (
		err             error
		ok              bool
		numberBytesRecd int
		from            unix.Sockaddr
		fromAddr        *unix.SockaddrNetlink
		nl              []NetlinkMessage
		rb              [ReceiveBufferSize]byte
	)

	if socketFd := int(atomic.LoadInt32(&s.fd)); socketFd < 0 {
		return nil, nil, errors.New("cannot receive. socket closed")
	} else {
		numberBytesRecd, from, err = unix.Recvfrom(socketFd, rb[:], 0)
	}

	if err != nil {
		return nil, nil, err
	}

	if fromAddr, ok = from.(*unix.SockaddrNetlink); !ok {
		return nil, nil, errors.New("error (converting to netlink sockaddr)")
	}

	if numberBytesRecd < NetlinkMessageHeaderSize {
		return nil, nil, errors.New("error (short response from netlink)")
	}

	msgLength := nlmAlignOf(numberBytesRecd)

	rb2 := make([]byte, msgLength)

	copy(rb2, rb[:msgLength])

	if nl, err = DeserializeToList(rb2); err != nil {
		return nil, nil, err
	}

	return nl, fromAddr, nil
}
