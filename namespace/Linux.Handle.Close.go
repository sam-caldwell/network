//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
)

// Close - close the NamespaceHandle and reset its file descriptor to -1 (closedHandle).
//
// WARNING: DO NOT USE AFTER Close() is called.  Bad things will happen.
func (ns *Handle) Close() error {
	if err := unix.Close(int(*ns)); err != nil {
		return err
	}
	*ns = closedHandle
	return nil
}
