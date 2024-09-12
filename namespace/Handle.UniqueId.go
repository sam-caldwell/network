package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
)

// UniqueId return string uniquely identifying the associated namespace
func (h *Handle) UniqueId() string {

	if *h == closedHandle {
		return "NS(none)"
	}

	var s unix.Stat_t
	if err := unix.Fstat(int(*h), &s); err != nil {
		return "NS(unknown)"
	}

	return fmt.Sprintf("NS(%d:%d)", s.Dev, s.Ino)

}
