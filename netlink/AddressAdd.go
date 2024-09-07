package netlink

// AddressAdd adds an IP address to a network interface (link).
//
// This function mirrors the behavior of the shell command `ip addr add $addr dev ${link}`, which is used to add
// an IP address to a network interface on Linux. It uses the Netlink protocol to communicate with the kernel
// and request the addition of the IP address to the specified interface.
//
// If the provided address (`addr`) is an IPv4 address, and a broadcast address is not given, it will be calculated
// from the CIDR notation (e.g., if the address is at least /30, the broadcast address will be auto-generated).
//
// The `pkgHandle` is a global handle that provides the interface to Netlink. It allows for creating and
// sending Netlink messages to the kernel.
//
// The `AddressAdd()` function under the hood uses the RTM_NEWADDR Netlink message type to request
// the addition of the IP address. This type is used to notify the kernel to add an IP address to the network
// interface.
//
// Relevant C sources:
//   - The kernel handles IP address additions via RTNETLINK. The function `inet_rtm_newaddr()` in the Linux kernel
//     handles the addition of IPv4 addresses:
//     https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c
//   - Similarly, IPv6 addresses are handled by `inet6_rtm_newaddr()` in the Linux kernel:
//     https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c
//
// Parameters:
//   - `link`: The network interface (link) to which the address should be added. This represents an existing network
//     interface on the system (e.g., `eth0`, `wlan0`, etc.).
//   - `addr`: The address to be added to the interface. It contains both the IP address and other optional fields such as
//     broadcast address, label, scope, etc.
//
// The function uses the `pkgHandle.AddressAdd()` to interface with Netlink and send the request to the kernel.
//
// Returns:
// - An error if the address cannot be added (e.g., invalid address, interface not found, permission denied).
func AddressAdd(link Link, addr *Addr) error {
	// These kernel functions are responsible for checking if the provided address is valid, ensuring that the
	// interface supports it, and then updating the system's routing tables and address lists.
	return pkgHandle.AddressAdd(link, addr)
}
