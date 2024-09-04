package core

// Seg6IpTunnel - Enumeration for SEG6 IP Tunneling options.
//
// This type defines the available tunneling options for Segment Routing over IPv6 (SRv6).
// SEG6 tunneling allows encapsulating IPv6 traffic using Segment Routing Headers (SRH),
// providing a method to route traffic through specific segments in a network.
//
// For further reference, see the related Linux kernel source code:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_iptunnel.h
type Seg6IpTunnel uint8

// SEG6_IPTUNNEL_* Constants - Segment Routing tunnel attributes.
//
// These constants represent the attributes used in the SRv6 tunnel.
// The main attribute, `SEG6_IPTUNNEL_SRH`, defines the Segment Routing Header (SRH) used to encapsulate IPv6 packets.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_iptunnel.h
const (
	// Seg6IptunnelUnspec - SEG6_IPTUNNEL_UNSPEC - Unspecified attribute (0).
	// This is a placeholder for an unspecified or unknown attribute.
	Seg6IptunnelUnspec Seg6IpTunnel = 0

	// Seg6IptunnelSrh - SEG6_IPTUNNEL_SRH - Specifies the Segment Routing Header (SRH).
	// This attribute defines the SRH used to encapsulate the original IPv6 packet.
	Seg6IptunnelSrh Seg6IpTunnel = 1

	// __seg6IptunnelMax - __SEG6_IPTUNNEL_MAX - Maximum internal value for segment routing attributes.
	// This constant is used internally to calculate the maximum valid attribute.
	__seg6IptunnelMax Seg6IpTunnel = 2

	// Seg6IptunnelMax - SEG6_IPTUNNEL_MAX - Maximum value for valid SEG6 tunnel attributes.
	// This is the highest value for SEG6 tunneling attributes, excluding internal placeholders.
	Seg6IptunnelMax = __seg6IptunnelMax - 1
)
