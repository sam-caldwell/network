//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
	"os"
	"path"
)

// Delete - deletes a named network namespace from both the in-memory map and filesystem.
func Delete(name string) error {

	namedPath := path.Join(bindMountPath, name)

	err := unix.Unmount(namedPath, unix.MNT_DETACH)
	if err != nil {
		return err
	}

	return os.Remove(namedPath)

}
