package namespace

import (
	"golang.org/x/sys/unix"
)

// GetFromPath - Return a handle to the network namespace identified by the given path. This will first look up any
// existing namespace file handle from the namespaces map, or if none is found or the found file handle is closed,
// open the path and return its file handle.
func GetFromPath(path string) (Handle, error) {
	lock.Lock()
	defer lock.Unlock()
	if existingHandle, found := namespaces[path]; found {
		return existingHandle, nil
	}
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_CLOEXEC, 0)
	return Handle(fd), err
}
