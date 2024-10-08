//go:build linux

package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
)

// String - return the file descriptor, its dev and inode, as a string.
func (h *Handle) String() string {

	if *h == -1 {
		return "NS(none)"
	}

	var s unix.Stat_t

	if err := unix.Fstat(int(*h), &s); err != nil {
		return fmt.Sprintf("NS(%d: unknown)", h)
	}

	return fmt.Sprintf("NS(%d: %d, %d)", h, s.Dev, s.Ino)

}
