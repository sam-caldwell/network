//go:build linux

package namespace

// GetFromPath gets a handle to a network namespace
// identified by the path
func GetFromPath(path string) (Handle, error) {
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_CLOEXEC, 0)
	if err != nil {
		return -1, err
	}
	return Handle(fd), nil
}
