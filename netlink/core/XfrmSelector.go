package core

// XfrmSelector - Represents a traffic selector for IPsec Security Associations (SAs).
//
// This struct corresponds to the xfrm_selector structure in the Linux kernel, which is used to define
// the traffic matching rules for IPsec SAs. The traffic is matched using the source and destination IP addresses,
// source and destination ports, protocol, and other fields.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmSelector struct {
	// Destination IP address (IPv4/IPv6)
	Daddr XfrmAddress
	// Source IP address (IPv4/IPv6)
	Saddr XfrmAddress
	// Destination port (big-endian)
	Dport uint16
	// Destination port mask for matching ranges
	DportMask uint16
	// Source port (big-endian)
	Sport uint16
	// Source port mask for matching ranges
	SportMask uint16
	// Address family (AF_INET for IPv4, AF_INET6 for IPv6)
	Family uint16
	// Prefix length for destination address
	PrefixlenD uint8
	// Prefix length for source address
	PrefixlenS uint8
	// Protocol identifier (e.g., TCP, UDP)
	Proto uint8
	// Padding for 8-byte alignment
	Pad [3]byte
	// Interface index
	Ifindex int32
	// UID of the user who owns the SA
	User uint32
}
