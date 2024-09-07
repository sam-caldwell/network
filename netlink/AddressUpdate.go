package netlink

import "net"

// AddressUpdate represents an IP address change event in the system. This structure is typically used to capture
// notifications about IP addresses being added or removed from network interfaces (links).
//
// These updates are received via Netlink, specifically through the RTNETLINK subsystem, which sends notifications
// when there are changes to the IP configuration on the system.
//
// This structure contains details about the address change, such as the IP address itself, the associated network
// interface, and other relevant metadata such as scope and lifetime.
//
// For more details on the underlying C implementation, see:
// - `rtnl_notify()` in the Linux kernel: https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c
// - IPv4 address changes: https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c
// - IPv6 address changes: https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c
type AddressUpdate struct {
	// LinkAddress represents the IP address along with its subnet mask in the form of net.IPNet.
	// It contains both the IP address and the subnet information.
	LinkAddress net.IPNet

	// LinkIndex represents the index of the network interface (link) where the IP address has been added or removed.
	// The index uniquely identifies the interface in the system.
	LinkIndex int

	// Flags represents additional flags associated with the address, such as permanent, temporary, or other state
	// flags. These flags can include interface-specific or address-specific information.
	Flags int

	// Scope represents the scope of the IP address. It determines the visibility and usage of the address.
	// Scope can be one of the following:
	// - `RT_SCOPE_UNIVERSE` (global scope)
	// - `RT_SCOPE_LINK` (link-local scope)
	// - `RT_SCOPE_HOST` (host-local scope)
	// These scopes define where the address is valid and can be routed.
	Scope int

	// PreferredLft represents the preferred lifetime of the address in seconds. This field indicates how long
	// the address is preferred for use in outgoing connections before it becomes deprecated.
	// After this time, the address should not be used to initiate new connections.
	PreferredLft int

	// ValidLft represents the valid lifetime of the address in seconds. This field specifies how long the address
	// remains valid. After this time expires, the address is considered invalid and should not be used.
	ValidLft int

	// NewAddr is a boolean flag indicating whether the IP address was added (`true`) or removed (`false`).
	// This field is useful for distinguishing between new address addition events and address removal events.
	NewAddr bool
}
