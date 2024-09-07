package netlink

// AddressAdd - add an IP address to a link device.
//
// This is the same as: `ip addr add $addr dev ${link}`
//
// If `addr` is an IPv4 address and a broadcast address is not given, it will be calculated from the CIDR addr where
// the addr is at least /30.
func AddressAdd(link Link, addr *Addr) error {

	return pkgHandle.AddressAdd(link, addr)

}
