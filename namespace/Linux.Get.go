//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
	"os"
)

// Get - return a handle to the current thread's network namespace.
func Get() (Handle, error) {
	return GetFromThread(os.Getpid(), unix.Gettid())
}
