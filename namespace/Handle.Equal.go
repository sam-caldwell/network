package namespace

import (
	"golang.org/x/sys/unix"
)

// Equal - compare two network handles and return true if they refer to the same network namespace.
//
// This works by comparing the device and its file descriptor inode.
func (ns *Handle) Equal(other Handle) bool {
	var a, b unix.Stat_t
	return (*ns == other) || (unix.Fstat(int(*ns), &a) == nil) &&
		(unix.Fstat(int(other), &b) == nil) &&
		(a.Dev == b.Dev) &&
		(a.Ino == b.Ino)
}
