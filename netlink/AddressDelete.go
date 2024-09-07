package netlink

// AddressDelete - delete an IP address from the link device
//
// Equivalent to: `ip addr del $addr dev $link`
//
// If `addr` is an IPv4 address and a broadcast address is not given, it will be calculated from the CIDR addr where
// the addr is at least /30.
func AddressDelete(link Link, addr *Addr) error {

	return pkgHandle.AddrDel(link, addr)

}
