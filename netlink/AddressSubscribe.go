package netlink

import "github.com/sam-caldwell/network/namespace"

// AddressSubscribe listens for changes in network addresses (e.g., IP address additions or deletions) and sends
// notifications through the provided channel. This function allows users to subscribe to address change events.
//
// The subscription will remain active until the `done` channel is closed. The `done` channel should be used to
// signal that the subscription should end, and the function will clean up the Netlink subscription.
//
// The underlying implementation uses Netlink to listen for address changes and forwards them as `AddressUpdate`
// events through the provided channel. Netlink messages, such as `RTM_NEWADDR` and `RTM_DELADDR`, are used by the
// kernel to notify about address changes.
//
// Relevant C Sources:
// - Address change events are handled in the kernel by RTNETLINK in `rtnetlink.c`.
//   - IPv4 addresses: https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c
//   - IPv6 addresses: https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c
//
// Parameters:
// - `ch`: A channel to which address update notifications (`AddressUpdate` structs) will be sent.
//   - These notifications will contain information about IP address changes such as whether an address was added or removed.
//
// - `done`: A channel to signal when the subscription should stop. Closing this channel will stop the subscription and clean up resources.
//
// Returns:
// - `error`: If there is an error setting up the Netlink subscription or managing the Netlink communication, an error is returned.
//
// Example Usage:
// ```go
// updates := make(chan AddressUpdate)
// done := make(chan struct{})
//
//	go func() {
//	    for update := range updates {
//	        fmt.Println("Address updated:", update)
//	    }
//	}()
//
//	if err := AddressSubscribe(updates, done); err != nil {
//	    log.Fatal(err)
//	}
//
// ```
//
// This example sets up a subscription to address changes, printing each update until `done` is closed.
func AddressSubscribe(ch chan<- AddressUpdate, done <-chan struct{}) error {
	// Calls the addressSubscribeAt function, which is responsible for setting up the Netlink subscription.
	// It subscribes to address changes in the default network namespace (namespace.None()) and forwards the updates.
	// The following parameters are passed:
	// - `namespace.None()`: This indicates that the subscription will take place in the default network namespace.
	// - `ch`: The channel that will receive the address updates.
	// - `done`: The channel used to signal when the subscription should stop.
	// - `nil`: No specific error callback is provided here.
	// - `false`: Indicates that existing addresses should not be listed.
	// - `0`: No custom buffer size is specified.
	// - `nil`: No receive timeout is set.
	// - `false`: No force buffer resizing is requested.
	return addressSubscribeAt(
		namespace.None(), // The current network namespace (default).
		namespace.None(), // No custom namespace is provided.
		ch,               // Channel to send address updates.
		done,             // Channel to close when done.
		nil,              // No error callback provided.
		false,            // Do not list existing addresses.
		0,                // Default receive buffer size.
		nil,              // No custom receive timeout.
		false,            // Do not force buffer resizing.
	)
}
