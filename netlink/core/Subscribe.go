//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// Subscribe - Create a netlink socket with a given protocol (e.g. NETLINK_ROUTE) and subscribe it to multicast groups
// passed in variable argument list.
//
// Returns the netlink socket on which Receive() method can be called to retrieve the messages from the kernel.
func Subscribe(protocol int, groups ...uint) (*NetlinkSocket, error) {

	socketFd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, protocol)

	if err != nil {
		return nil, err
	}

	s := &NetlinkSocket{
		fd: int32(socketFd),
	}

	s.lsa.Family = unix.AF_NETLINK

	for _, g := range groups {
		s.lsa.Groups |= 1 << (g - 1)
	}

	if err := unix.Bind(socketFd, &s.lsa); err != nil {
		_ = unix.Close(socketFd)
		return nil, err
	}

	return s, nil
}
