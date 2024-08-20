//go:build linux

package namespace

import "os"

// Get gets a handle to the current threads network namespace.
func Get() (Handle, error) {
	return GetFromThread(os.Getpid(), unix.Gettid())
}
