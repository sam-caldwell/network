//go:build linux

package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
	"runtime"
)

// Close - close the NamespaceHandle and reset its file descriptor to -1 (closedHandle).
//
// WARNING: DO NOT USE AFTER Close() is called.  Bad things will happen.
func (ns *Handle) Close() error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if *ns == closedHandle {
		return fmt.Errorf("handle is already closed")
	}
	if err := unix.Close(int(*ns)); err != nil {
		return err
	}
	*ns = closedHandle
	return nil
}
