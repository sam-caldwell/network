package namespace

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// UniqueId returns a string which uniquely identifies the namespace
// associated with the network handle.
func (ns *Handle) UniqueId() string {
	if *ns == -1 {
		return "NS(none)"
	}
	var s unix.Stat_t
	if err := unix.Fstat(int(*ns), &s); err != nil {
		return "NS(unknown)"
	}
	return fmt.Sprintf("NS(%d:%d)", s.Dev, s.Ino)
}
