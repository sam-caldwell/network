package core

import (
	"golang.org/x/sys/unix"
)

// NewNetlinkSocket - create and return a new NetlinkSocket instance
func NewNetlinkSocket(fd int, sa *unix.SockaddrNetlink) *NetlinkSocket {
	return &NetlinkSocket{
		fd:  int32(fd),
		lsa: *sa,
	}
}
