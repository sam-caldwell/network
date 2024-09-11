package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"runtime"
)

// Handle represents a handle to a network namespace in the Linux kernel.
// It can be cast directly to an int and used as a file descriptor for system calls such as `setns()`.
//
// This type directly corresponds to a file descriptor that points to a namespace in the Linux kernel.
// Network namespaces are represented in `/proc/[pid]/ns/net`, which can be accessed or manipulated using
// system calls like `unshare()` and `setns()`, allowing processes to join, leave, or isolate their network stack.
//
// References:
// - Network Namespace Documentation: https://www.kernel.org/doc/Documentation/networking/netns.txt
// - Linux Kernel Source (Netlink and Namespaces): https://github.com/torvalds/linux/tree/master/net/core
//
// Usage:
// - A `Handle` can be obtained via functions like `Get()` or `NewHandle()`, which retrieve or create a namespace.
// - Once a handle is obtained, it can be passed to system calls (e.g., `setns(fd, CLONE_NEWNET)`).
type Handle int

// NewHandle - creates a new network namespace using the `unshare` system call.
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
func NewHandle() (Handle, error) {

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Call unshare to create a new network namespace (CLONE_NEWNET isolates the network namespace).
	if err := unix.Unshare(unix.CLONE_NEWNET); err != nil {
		return closedHandle, err
	}

	// Retrieve and return a handle to the newly created network namespace.
	return Opsys.GetFromThread(os.Getpid(), unix.Gettid())
}

// String - return the file descriptor, its dev and inode, as a string.
func (ns *Handle) String() string {
	if *ns == -1 {
		return "NS(none)"
	}
	var s unix.Stat_t
	if err := unix.Fstat(int(*ns), &s); err != nil {
		return fmt.Sprintf("NS(%d: unknown)", ns)
	}
	return fmt.Sprintf("NS(%d: %d, %d)", ns, s.Dev, s.Ino)
}

// UniqueId return string uniquely identifying the associated namespace
func (ns *Handle) UniqueId() string {
	if *ns == -1 {
		return "NS(none)"
	}
	var s unix.Stat_t
	if err := unix.Fstat(int(*ns), &s); err != nil {
		return "NS(unknown)"
	}
	return fmt.Sprintf("NS(%d:%d)", s.Dev, s.Ino)
}

// Set - set current network namespace to the namespace represented by NamespaceHandle.
func Set(ns Handle) error {
	return unix.Setns(int(ns), unix.CLONE_NEWNET)
}

// SetNamespace - set the namespace (using golang.org/x/sys/unix.Setns)
//
// We keep this wrapper around unix.Setns() because SetNamespace() is more readable
// and less likely to be confused with a few other things we are building (e.g. name server).
// ...and there's the golang love of changing things under us that make wrappers appealing.
//
// See https://man7.org/linux/man-pages/man2/setns.2.html
func SetNamespace(namespace Handle, namespaceType NsTypes) error {
	return unix.Setns(int(namespace), int(namespaceType))
}

// IsOpen returns true if Close() the handle has not been closed (-1)
func (ns *Handle) IsOpen() bool {
	return *ns != -1
}

// None - return an empty (-1: closed) namespace handle.
func None() Handle {
	return Handle(closedHandle)
}

// Equal - compare two network handles and return true if they refer to the same network namespace.
//
// This works by comparing the device and its file descriptor inode.
func (ns *Handle) Equal(other Handle) bool {
	var a, b unix.Stat_t
	return (*ns == other) || (unix.Fstat(int(*ns), &a) == nil) &&
		(unix.Fstat(int(other), &b) == nil) &&
		(a.Dev == b.Dev) &&
		(a.Ino == b.Ino)
}

// Close - close the NamespaceHandle and reset its file descriptor to -1 (closedHandle).
//
// WARNING: DO NOT USE AFTER Close() is called.  Bad things will happen.
func (ns *Handle) Close() error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if *ns == closedHandle {
		return fmt.Errorf("handle is already closed")
	}
	if err := unix.Close(int(*ns)); err != nil {
		return err
	}
	*ns = closedHandle
	return nil
}
