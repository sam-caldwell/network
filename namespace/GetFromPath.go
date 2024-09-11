//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
)

// GetFromPath - return a handle to the network namespace identified by the given path.
func GetFromPath(path string) (Handle, error) {
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_CLOEXEC, 0)
	return Handle(fd), err
}
