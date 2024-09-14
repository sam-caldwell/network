package core

import (
	"errors"
	"golang.org/x/sys/unix"
	"sync/atomic"
)

// GetPid retrieves the PID (Process ID) associated with the netlink socket.
//
// This function uses the socket's file descriptor to obtain the socket address (Sockaddr)
// associated with the netlink socket. If the socket is of type SockaddrNetlink, it returns
// the PID stored in the socket address structure. Otherwise, it returns an error indicating
// that the socket type is unsupported.
func (s *NetlinkSocket) GetPid() (uint32, error) {

	socketFd := int(atomic.LoadInt32(&s.fd))

	lsa, err := unix.Getsockname(socketFd)

	if err != nil {
		return 0, err
	}

	switch v := lsa.(type) {
	case *unix.SockaddrNetlink:
		// Return the Pid from the SockaddrNetlink structure.
		return v.Pid, nil
	}

	return 0, errors.New("unsupported socket type")
}
