//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
)

// Equal - compare two network handles and return true if they refer to the same network namespace.
//
// This works by comparing the device and its file descriptor inode.
func (ns *Handle) Equal(other Handle) bool {
	if *ns == other {
		return true
	}
	var a, b unix.Stat_t
	if err := unix.Fstat(int(*ns), &a); err != nil {
		return false
	}
	if err := unix.Fstat(int(other), &b); err != nil {
		return false
	}
	return (a.Dev == b.Dev) && (a.Ino == b.Ino)
}
