package netlink

import "github.com/sam-caldwell/network/namespace"

// AddressSubscribeAt allows the caller to subscribe to address change events in a specific network namespace.
// This function works like AddressSubscribe but provides an option to specify the network namespace (`ns`)
// in which the subscription should be established.
//
// Netlink events related to IP address changes (such as addition or removal of addresses) are monitored in the
// specified namespace. These updates are sent to the provided channel (`ch`). The subscription remains active until
// the `done` channel is closed, at which point the Netlink subscription is cleaned up.
//
// Relevant C Sources:
// - Address changes for both IPv4 and IPv6 are handled in the Linux kernel's RTNETLINK subsystem.
//   - IPv4 address changes: https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c
//   - IPv6 address changes: https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c
//
// Parameters:
//   - `ns`: The network namespace (`namespace.Handle`) in which to subscribe to address changes. This allows subscribing
//     in network namespaces other than the default one, which is particularly useful in containerized environments where
//     each container has its own network namespace.
//   - `ch`: A channel that will receive `AddressUpdate` notifications when addresses change in the specified namespace.
//     The notifications will contain details about the IP address change, such as whether an address was added or removed.
//   - `done`: A channel to signal when the subscription should stop. Closing this channel will terminate the subscription
//     and clean up the Netlink connection.
//
// Returns:
// - `error`: If an error occurs while setting up the Netlink subscription or handling the namespace, the function returns an error.
//
// Usage Example:
// ```go
// nsHandle, err := namespace.Open("/var/run/netns/my_namespace")
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// defer nsHandle.Close()
//
// updates := make(chan AddressUpdate)
// done := make(chan struct{})
//
//	go func() {
//	    for update := range updates {
//	        fmt.Println("Address updated:", update)
//	    }
//	}()
//
//	if err := AddressSubscribeAt(nsHandle, updates, done); err != nil {
//	    log.Fatal(err)
//	}
//
// ```
//
// In this example, the user subscribes to address changes within a specific network namespace (`my_namespace`),
// and the updates are printed to the console until the `done` channel is closed.
func AddressSubscribeAt(ns namespace.Handle, ch chan<- AddressUpdate, done <-chan struct{}) error {
	// Calls the internal addressSubscribeAt function with the specified network namespace (`ns`).
	// The following parameters are passed:
	// - `ns`: The network namespace to subscribe in.
	// - `namespace.None()`: No fallback namespace is provided.
	// - `ch`: Channel to send address updates.
	// - `done`: Channel to signal when to stop the subscription.
	// - `nil`: No error callback is provided.
	// - `false`: Do not list existing addresses.
	// - `0`: Default receive buffer size.
	// - `nil`: No custom receive timeout is set.
	// - `false`: Do not force buffer resizing.
	return addressSubscribeAt(ns, namespace.None(), ch, done, nil, false, 0, nil, false)
}
