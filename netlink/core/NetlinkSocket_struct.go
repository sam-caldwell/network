package core

import (
	"golang.org/x/sys/unix"
	"sync"
)

// NetlinkSocket - sockaddr_nl - represents a netlink socket with an associated file descriptor (fd) and sockaddr
// structure (lsa). The socket is protected by a mutex to ensure thread-safe operations, preventing concurrent access
// issues during send, receive, or other socket operations.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
type NetlinkSocket struct {

	// Mutex to ensure thread-safe operations on the socket.
	sync.Mutex

	// fd - File descriptor for the netlink socket.
	fd int32

	// lsa - Socket address information specific to netlink.
	lsa unix.SockaddrNetlink
}
