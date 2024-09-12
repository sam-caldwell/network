package namespace

import (
	"golang.org/x/sys/unix"
)

// Equal - compare two network handles and return true if they refer to the same network namespace.
//
// This works by comparing the device and its file descriptor inode.
func (h *Handle) Equal(other Handle) bool {
	var a, b unix.Stat_t
	return *h == other ||
		(unix.Fstat(int(*h), &a) == nil &&
			unix.Fstat(int(other), &b) == nil &&
			a.Dev == b.Dev &&
			a.Ino == b.Ino)
}
