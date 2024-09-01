package core

import (
	"golang.org/x/sys/unix"
)

// rtaAlignOf - Round the length of a netlink message up to align it properly.
// Taken from syscall/netlink_linux.go by The Go Authors under BSD-style license.
func rtaAlignOf(attrlen int) int {
	return (attrlen + unix.RTA_ALIGNTO - 1) & ^(unix.RTA_ALIGNTO - 1)
}
