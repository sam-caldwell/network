package core

// IPSetError represents an error type for IP sets in the Netlink interface.
//
// In the context of Netlink communication, IP sets are used to manage sets of IP addresses, ports, and other elements in the kernel.
// Errors in managing these sets are represented using error codes that are often passed as pointers or special values.
//
// This type mirrors the error handling mechanisms in the Linux kernel for IP sets, particularly in the Netfilter subsystem.
//
// References:
// - Linux Kernel IPSet: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
// - Netlink Error Handling: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
type IPSetError uintptr
