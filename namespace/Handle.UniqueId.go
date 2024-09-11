package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
)

// UniqueId return string uniquely identifying the associated namespace
func (ns *Handle) UniqueId() string {

	if *ns == closedHandle {
		return "NS(closedHandle)"
	}

	var s unix.Stat_t
	if err := unix.Fstat(int(*ns), &s); err != nil {
		return "NS(unknown)"
	}

	return fmt.Sprintf("NS(%d:%d)", s.Dev, s.Ino)

}
