package namespace

import (
	"golang.org/x/sys/unix"
	"os"
	"path"
)

// Delete - deletes a named network namespace from both the in-memory map and filesystem.
func Delete(name string) error {
	lock.Lock()
	defer lock.Unlock()

	namedPath := path.Join(bindMountPath, name)

	err := unix.Unmount(namedPath, unix.MNT_DETACH)
	if err != nil {
		return err
	}
	if handle, found := namespaces[name]; found {
		if err := handle.Close(); err != nil {
			return err
		}
		delete(namespaces, name)
	}
	return os.Remove(namedPath)
}
