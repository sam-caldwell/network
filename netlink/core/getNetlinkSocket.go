package core

import (
	"golang.org/x/sys/unix"
)

// getNetlinkSocket - return the NetlinkSocket struct by reference
func getNetlinkSocket(protocol int) (*NetlinkSocket, error) {

	fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW|unix.SOCK_CLOEXEC, protocol)

	if err != nil {
		return nil, err
	}

	s := &NetlinkSocket{
		fd: int32(fd),
	}

	s.lsa.Family = unix.AF_NETLINK

	if err := unix.Bind(fd, &s.lsa); err != nil {
		_ = unix.Close(fd)
		return nil, err
	}

	return s, nil
}
