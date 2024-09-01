package core

import (
	"golang.org/x/sys/unix"
)

// nlmAlignOf - Round the length of a netlink message up to align it properly.
// Taken from syscall/netlink_linux.go by The Go Authors under BSD-style license.
func nlmAlignOf(msglen int) int {
	return (msglen + unix.NLMSG_ALIGNTO - 1) & ^(unix.NLMSG_ALIGNTO - 1)
}
