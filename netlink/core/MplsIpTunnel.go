package core

// MplsIptunnel represents the MPLS IP Tunnel attributes for netlink communication.
// MPLS (Multiprotocol Label Switching) allows for forwarding packets based on labels rather than network addresses.
// This enumeration corresponds to the MPLS IP tunnel attributes that are used to manage and configure tunnels in the Linux kernel.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls_iptunnel.h

type MplsIptunnel uint8

const (
	// MplsIptunnelUnspec - MPLS_IPTUNNEL_UNSPEC -
	// Unspecified MPLS IP tunnel attribute.
	MplsIptunnelUnspec MplsIptunnel = iota

	// MplsIptunnelDst - MPLS_IPTUNNEL_DST -
	// Destination address for the MPLS IP tunnel.
	MplsIptunnelDst
)
