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
	var s1, s2 unix.Stat_t
	if err := unix.Fstat(int(*ns), &s1); err != nil {
		return false
	}
	if err := unix.Fstat(int(other), &s2); err != nil {
		return false
	}
	return (s1.Dev == s2.Dev) && (s1.Ino == s2.Ino)
}
