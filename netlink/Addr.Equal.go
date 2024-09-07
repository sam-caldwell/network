package netlink

// Equal - returns true if both Addresses have the same net.IPNet value. Ignore label for comparison.
func (addr Addr) Equal(rhs Addr) bool {

	return addr.IP.Equal(rhs.IP) && addr.Mask.Size() == rhs.Mask.Size()

}
