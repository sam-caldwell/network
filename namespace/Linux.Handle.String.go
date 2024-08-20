//go:build linux

package namespace

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// String - return the file descriptor, its dev and inode, as a string.
func (ns *Handle) String() string {
	if *ns == -1 {
		return "NS(none)"
	}
	var s unix.Stat_t
	if err := unix.Fstat(int(*ns), &s); err != nil {
		return fmt.Sprintf("NS(%d: unknown)", ns)
	}
	return fmt.Sprintf("NS(%d: %d, %d)", ns, s.Dev, s.Ino)
}
