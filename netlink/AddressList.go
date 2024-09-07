package netlink

// AddressList returns a list of IP addresses in the system, filtered by the specified network link (interface)
// and IP family (IPv4 or IPv6). It replicates the functionality of the shell command `ip addr show`, which
// is used to display all IP addresses assigned to the network interfaces on the system.
//
// This function uses the Netlink protocol to retrieve the list of IP addresses from the kernel. Specifically,
// it interacts with the RTNETLINK subsystem to fetch information about IP addresses for the specified link and IP family.
//
// Relevant C sources:
//   - RTNETLINK is responsible for network interface and address management in the Linux kernel.
//     Address listing is handled in functions such as `inet_dump_ifaddr()` for IPv4 and `inet6_dump_addr()` for IPv6.
//   - IPv4 address listing: https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c
//   - IPv6 address listing: https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c
//
// Parameters:
//   - `link`: The network interface (link) for which the IP addresses should be listed. This corresponds to a
//     specific network interface on the system, such as `eth0`, `wlan0`, or `lo` (loopback).
//   - `family`: The address family to filter by. This can be either `AF_INET` (for IPv4) or `AF_INET6` (for IPv6).
//
// The function communicates with the kernel using RTNETLINK messages (such as `RTM_GETADDR`) to fetch the IP addresses.
// It uses the `pkgHandle.AddressList()` function, which builds and sends the appropriate Netlink request.
//
// Returns:
// - A slice of `Addr` objects, each representing an IP address assigned to the specified link.
// - An error if there is an issue retrieving the addresses (e.g., invalid link, permission denied, etc.).
func AddressList(link Link, family int) ([]Addr, error) {
	// The `pkgHandle.AddressList()` function sends an `RTM_GETADDR` Netlink message to request a list of IP addresses
	// from the kernel for the specified network interface and address family. This message is part of the RTNETLINK
	// protocol and is processed in the kernel by functions like:
	//
	// - `inet_dump_ifaddr()` in `devinet.c` for IPv4 addresses:
	//   https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c
	//
	// - `inet6_dump_addr()` in `addrconf.c` for IPv6 addresses:
	//   https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c
	//
	// These kernel functions gather the IP addresses assigned to the specified interface (link) and return them
	// in a structured format (Netlink messages) to the calling user-space process.
	//
	// The `family` parameter allows filtering by IPv4 (`AF_INET`) or IPv6 (`AF_INET6`).
	//
	// The function returns a slice of `Addr` objects, which represent IP addresses along with other metadata such
	// as broadcast addresses, scope, flags, etc.
	return pkgHandle.AddressList(link, family)
}
