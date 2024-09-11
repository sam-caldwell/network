//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
)

// New creates a new network namespace using the `unshare` system call.
// This function isolates the network stack of the calling process by creating a new network namespace
// and setting it as the current one.
//
// The `unshare` system call is used with the `CLONE_NEWNET` flag to create a new network namespace.
// After successfully creating the new namespace, the function retrieves a handle to it using the `Get()` function.
//
// Related Linux source files:
//   - The `unshare` system call is defined in the Linux kernel at:
//     https://github.com/torvalds/linux/blob/master/kernel/fork.c
//   - Network namespaces are handled in the `net/` directory of the Linux source tree, particularly in:
//     https://github.com/torvalds/linux/blob/master/net/core/net_namespace.c
//
// Returns:
//   - Handle: A handle to the new network namespace.
//   - error: An error if the system call fails (e.g., insufficient privileges or unsupported kernel).
func New() (Handle, error) {

	// Call unshare to create a new network namespace (CLONE_NEWNET isolates the network namespace).
	if err := unix.Unshare(unix.CLONE_NEWNET); err != nil {
		return closedHandle, err
	}

	// Retrieve and return a handle to the newly created network namespace.
	return Get()
}
