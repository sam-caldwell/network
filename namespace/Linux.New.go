//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
)

// New creates a new network namespace, sets it as current and returns a handle to it.
func New() (Handle, error) {
	if err := unix.Unshare(unix.CLONE_NEWNET); err != nil {
		return -1, err
	}
	return Get()
}
