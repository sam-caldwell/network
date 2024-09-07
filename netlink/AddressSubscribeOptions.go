package netlink

import (
	"github.com/sam-caldwell/network/namespace"
	"golang.org/x/sys/unix"
)

// AddressSubscribeOptions contains a set of options to use with AddressSubscribeWithOptions.
// These options allow fine-tuning the behavior when subscribing to netlink messages related to IP address changes.
// Subscriptions to address changes are typically done using the RTNETLINK protocol.
//
// Relevant C Source code:
// - RTNETLINK subscription: https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c
// - IP Address management in the Linux kernel: https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c (IPv4)
//   and https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c (IPv6)

type AddressSubscribeOptions struct {
	// Namespace allows specifying a network namespace for the subscription. In Linux, each
	// network namespace has its own set of network interfaces, IP addresses, and routing tables.
	// This option allows subscribing to events in a specific network namespace rather than the default one.
	//
	// Corresponds to the Linux kernel's network namespace handling:
	// - `setns()` syscall: https://man7.org/linux/man-pages/man2/setns.2.html
	// - Network namespace switching in the kernel: https://github.com/torvalds/linux/blob/master/fs/nsfs.c
	Namespace *namespace.Handle

	// ErrorCallback is a function that is called when an error occurs while receiving netlink messages.
	// This can be used for logging or handling errors asynchronously during the subscription.
	//
	// Related C handling in RTNETLINK:
	// - `rtnl_notify()`: https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c
	// - Error handling in netlink messages is managed by the kernel’s `netlink_rcv_skb()`.
	ErrorCallback func(error)

	// ListExisting defines whether existing IP addresses should be listed when the subscription
	// starts. If true, the current state of all IP addresses will be provided before the subscription
	// starts receiving new events. This is useful for applications that need an initial snapshot
	// of IP addresses before listening for changes.
	//
	// Corresponds to the initial IP address dump handled by:
	// - `rtnl_dump_all()`: https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c
	ListExisting bool

	// ReceiveBufferSize sets the size of the recv buffer for the netlink socket. This controls
	// how much data can be received at once. If set, the system will attempt to allocate a buffer of this size.
	//
	// Corresponds to setting the socket receive buffer size via `setsockopt()`, specifically:
	// - `SO_RCVBUF`: https://github.com/torvalds/linux/blob/master/net/core/sock.c
	ReceiveBufferSize int

	// ReceiveBufferForceSize, when set to true, forces the recv buffer size to the value
	// specified by ReceiveBufferSize, overriding system defaults or limits.
	//
	// Corresponds to `SO_RCVBUFFORCE` socket option, which allows privileged processes to set
	// a buffer size that exceeds the normal system limits:
	// - `SO_RCVBUFFORCE`: https://github.com/torvalds/linux/blob/master/net/core/sock.c
	ReceiveBufferForceSize bool

	// ReceiveTimeout sets a timeout for receiving netlink messages. If a message is not received
	// within the specified timeout, an error will be raised. This is useful to prevent blocking indefinitely
	// while waiting for events.
	//
	// This corresponds to the `SO_RCVTIMEO` option in `setsockopt()`:
	// - `SO_RCVTIMEO`: https://github.com/torvalds/linux/blob/master/net/core/sock.c
	ReceiveTimeout *unix.Timeval
}
