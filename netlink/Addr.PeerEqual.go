package netlink

// PeerEqual - Compare two peer CIDR addresses but ignore any label comparison
func (addr Addr) PeerEqual(rhs Addr) bool {

	return addr.Peer.IP.Equal(rhs.Peer.IP) && addr.Peer.Mask.Size() == rhs.Peer.Mask.Size()

}
