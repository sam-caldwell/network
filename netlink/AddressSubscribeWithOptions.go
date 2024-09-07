package netlink

import (
	"github.com/sam-caldwell/network/namespace"
)

// AddressSubscribeWithOptions -  subscribes to address updates in the system, similar to AddrSubscribe, but allows
// additional options to modify its behavior. This function is particularly useful for monitoring IP address
// changes across network interfaces in Linux. You can provide a specific network namespace, an error callback,
// and other options like buffer sizes and timeouts.
//
// It leverages the RTNETLINK protocol to monitor changes in IP addresses, which is implemented in the Linux kernel.
// Relevant C source code for RTNETLINK and address management:
// - RTNETLINK subscription handling: https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c
// - IP address management (IPv4): https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c
// - IP address management (IPv6): https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c
//
// Parameters:
//   - `ch`: Channel where `AddrUpdate` events will be sent when IP address changes are detected. These updates can
//     include address additions, deletions, or modifications.
//   - `done`: Channel used to signal when the subscription should stop. Once this channel receives a signal, the
//     subscription will be terminated, and no more updates will be sent.
//   - `options`: A set of options that modify the behavior of the subscription. This includes:
//   - `Namespace`: Allows specifying a network namespace for the subscription.
//   - `ErrorCallback`: A callback function that is invoked when an error occurs while processing netlink messages.
//   - `ListExisting`: If true, existing IP addresses will be listed before subscribing to changes.
//   - `ReceiveBufferSize`: Sets the size of the receive buffer for the netlink socket.
//   - `ReceiveBufferForceSize`: Forces the buffer size, overriding system limits if necessary.
//   - `ReceiveTimeout`: Sets a timeout for receiving netlink messages.
//
// The C equivalent for subscribing to address updates using RTNETLINK can be seen in functions like:
//   - `rtnl_addr_notify()` in https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c
//     This handles the actual sending of netlink notifications when address changes occur.
//   - `inet_rtm_newaddr()` in https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c (for IPv4 address changes).
//   - `inet6_rtm_newaddr()` in https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c (for IPv6 address changes).
//
// Returns:
//   - An error if the subscription could not be started, such as an invalid namespace, buffer allocation issue,
//     or failure in creating the netlink socket.
func AddressSubscribeWithOptions(ch chan<- AddrUpdate, done <-chan struct{}, options AddressSubscribeOptions) error {
	// If no namespace is provided, use the default (None).
	if options.Namespace == nil {
		// namespace.None() returns the default namespace (or an uninitialized namespace).
		// Network namespaces allow isolating network interfaces and routing tables for different processes.
		// - Reference for namespaces in C: https://github.com/torvalds/linux/blob/master/fs/nsfs.c
		none := namespace.None()
		options.Namespace = &none
	}

	// Start the subscription by calling addrSubscribeAt with the provided options.
	// This function subscribes to IP address changes using RTNETLINK within the specified namespace.
	// - The 'namespace' package provides the ability to switch to a specific network namespace.
	// - 'addrSubscribeAt()' handles the actual interaction with RTNETLINK to listen for address changes.
	return addrSubscribeAt(
		*options.Namespace,             // The network namespace for the subscription
		namespace.None(),               // Fallback to default namespace if none is provided
		ch,                             // Channel for sending address updates
		done,                           // Channel to signal subscription termination
		options.ErrorCallback,          // Callback for error handling
		options.ListExisting,           // Flag to list existing addresses before subscribing
		options.ReceiveBufferSize,      // Size of the receive buffer for netlink messages
		options.ReceiveTimeout,         // Timeout for receiving messages
		options.ReceiveBufferForceSize, // Force the buffer size, even if it exceeds system limits
	)
}
