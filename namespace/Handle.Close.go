package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
	"runtime"
)

// Close - close the NamespaceHandle and reset its file descriptor to -1 (closedHandle).
//
// WARNING: DO NOT USE AFTER Close() is called.  Bad things will happen.
func (h *Handle) Close() error {

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if *h == closedHandle {
		return fmt.Errorf("handle is closed")
	}

	if err := unix.Close(int(*h)); err != nil {
		return err
	}

	*h = closedHandle

	return nil

}
