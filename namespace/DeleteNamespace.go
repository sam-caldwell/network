//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
	"os"
	"path"
)

// DeleteNamespace deletes a named network namespace
func DeleteNamespace(name string) error {
	namedPath := path.Join(bindMountPath, name)

	err := unix.Unmount(namedPath, unix.MNT_DETACH)
	if err != nil {
		return err
	}

	return os.Remove(namedPath)
}
