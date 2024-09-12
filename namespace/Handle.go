//go:build linux

package namespace

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
// - A `Handle` can be obtained via functions like `Get()` or `New()`, which retrieve or create a namespace.
// - Once a handle is obtained, it can be passed to system calls (e.g., `setns(fd, CLONE_NEWNET)`).
type Handle int
