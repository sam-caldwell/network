package netlink

// AddressReplace adds or replaces an IP address on a specified network link (interface).
// This operation is an "upsert" (update or insert) of an IP address on the device. If the IP address
// exists, it will be replaced with the new one; if it doesn't exist, it will be added.
//
// The equivalent shell command is: `ip addr replace $addr dev $link`
//
// This function uses Netlink to communicate with the kernel, allowing the addition or replacement of IP addresses
// on network interfaces. If the provided address is an IPv4 address and no broadcast address is specified, the
// broadcast address will be calculated automatically for subnets with a mask of at least /30.
//
// Relevant C sources:
//   - IPv4 address addition/replacement is handled in the Linux kernel by `inet_rtm_newaddr()` in `devinet.c`:
//     https://github.com/torvalds/linux/blob/master/net/ipv4/devinet.c
//   - IPv6 address addition/replacement is handled by `inet6_rtm_newaddr()` in `addrconf.c`:
//     https://github.com/torvalds/linux/blob/master/net/ipv6/addrconf.c
//
// Parameters:
//   - `link`: The network link (interface) to which the address will be added or replaced. This is analogous to specifying
//     the `dev $link` part of the `ip` command, where `$link` is the network interface (e.g., eth0, wlan0).
//   - `addr`: The address (including the subnet mask) that should be added or replaced on the link. This structure contains
//     the IP address, subnet mask, and other optional settings like broadcast, label, flags, and scope.
//
// Returns:
//   - `error`: If there is an issue with the operation, such as a Netlink communication error or invalid parameters,
//     an error will be returned.
//
// Example Usage:
// ```go
// link, err := LinkByName("eth0")
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	addr := &Addr{
//	    IPNet: &net.IPNet{
//	        IP:   net.ParseIP("192.168.1.100"),
//	        Mask: net.CIDRMask(24, 32),
//	    },
//	    Broadcast: net.ParseIP("192.168.1.255"),
//	}
//
//	if err := AddressReplace(link, addr); err != nil {
//	    log.Fatal(err)
//	}
//
// ```
//
// In this example, the function replaces or adds the IP address `192.168.1.100/24` on the `eth0` interface with
// the broadcast address `192.168.1.255`.
func AddressReplace(link Link, addr *Addr) error {
	// Use the package handle to invoke the address replacement operation via Netlink.
	// The pkgHandle.AddrReplace function is responsible for sending the correct Netlink message to the kernel.
	// This replaces the existing IP address on the specified link, or adds it if it does not already exist.
	return pkgHandle.AddressReplace(link, addr)
}
