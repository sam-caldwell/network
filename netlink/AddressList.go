package netlink

// AddressList - return a list of IP addresses in the system and filter by link and ip family.
//
// Shell Command: `ip addr show`.
func AddressList(link Link, family int) ([]Addr, error) {
	return pkgHandle.AddrList(link, family)
}
