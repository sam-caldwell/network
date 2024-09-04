package core

// Ipv6SrCrt - represents IPv6 Segment Routing Header (SRH) types.
//
// Segment Routing with IPv6 (SRv6) uses these routing header types for routing packets through specific paths.
// Some of these types are deprecated in modern networking systems.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/ipv6_route.h
type Ipv6SrCrt uint8

const (
	// Ipv6SrcrtStrict - IPV6_SRCRT_STRICT - Deprecated; will be removed.
	// Specifies a strict source routing option, which is no longer used.
	Ipv6SrcrtStrict Ipv6SrCrt = 0x01

	// Ipv6SrcrtType0 - IPV6_SRCRT_TYPE_0 - Deprecated; will be removed.
	// Type 0 Routing Header was used for loose source routing but is now deprecated.
	Ipv6SrcrtType0 Ipv6SrCrt = 0

	// Ipv6SrcrtType2 - IPV6_SRCRT_TYPE_2 - IPv6 type 2 Routing Header.
	// Used for mobile IPv6 routing, allowing for redirection between home and foreign agents.
	Ipv6SrcrtType2 Ipv6SrCrt = 2

	// Ipv6SrcrtType4 - IPV6_SRCRT_TYPE_4 - Segment Routing with IPv6.
	// Used for Segment Routing with IPv6 to specify segment endpoints in the network.
	Ipv6SrcrtType4 Ipv6SrCrt = 4
)
