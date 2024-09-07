package netlink

import "net"

// Addr represents an IP address in netlink. Netlink IP addresses include a mask,
// which is why the address is stored as a net.IPNet. This struct is commonly used
// for IP address management in Linux, such as adding, deleting, or modifying
// IP addresses associated with network interfaces.
//
// This struct corresponds to the way IP addresses are managed via RTNETLINK in the
// Linux kernel, particularly in the IP address management functions. The primary
// C source code reference is:
// - https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c (RTNETLINK interface)
// - https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c (IPv4 address management)
// - https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c (IPv6 address management)
type Addr struct {

	// IPNet - represents the IP address and network mask. It's a pointer to net.IPNet,
	// which contains the IP address and the associated subnet mask (e.g., 192.168.0.1/24).
	*net.IPNet

	// Label - is the label associated with the address. It can be used for distinguishing between
	// multiple addresses associated with the same interface.
	Label string

	// Flags - represent the flags associated with the IP address (e.g., IFA_F_SECONDARY).
	// These flags can indicate special properties of the address, such as whether it's secondary.
	Flags int

	// Scope - defines the scope of the address (e.g., link-local, global, site-local).
	// This value determines the reachability of the address in the network.
	Scope int

	// Peer - represents the peer IP address for a point-to-point interface.
	// It's used for interfaces that have a point-to-point connection (e.g., tunnels).
	Peer *net.IPNet

	// Broadcast - represents the broadcast address associated with this IP.
	// It is used for IPv4 addresses that support broadcasting (e.g., 192.168.0.255).
	Broadcast net.IP

	// PreferredLft - represents the preferred lifetime of the IP address, in seconds.
	// After this time, the IP address is no longer preferred for new connections.
	PreferredLft int

	// ValidLft - represents the valid lifetime of the IP address, in seconds.
	// After this time, the address is considered invalid and is typically removed.
	ValidLft int

	// LinkIndex - represents the index of the network interface this address is associated with.
	// It's used to identify the specific network interface on which the address is configured.
	LinkIndex int
}
