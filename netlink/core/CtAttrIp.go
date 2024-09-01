package core

// CtAttrIp - ctattr_ip - structure in the Linux Netfilter framework that stores IP layer-specific attributes,
// such as source and destination IP addresses, for a tracked connection within the connection tracking system.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrIp uint8

const (
	// CtaIpUnspec - CTA_IP_UNSPEC - indicate that no specific IP-related information is provided or applicable
	// within the connection tracking system.
	CtaIpUnspec = 0

	// CtaIpV4Src - CTA_IP_V4_SRC - represent the source IPv4 address for a tracked connection within the
	// connection tracking system.
	CtaIpV4Src = 1

	// CtaIpV4Dst - CTA_IP_V4_DST - represent the destination IPv4 address for a tracked connection within
	// the connection tracking system.
	CtaIpV4Dst = 2

	// CtaIpV6Src - CTA_IP_V6_SRC - represent the source IPv6 address for a tracked connection within the connection
	// tracking system.
	CtaIpV6Src = 3

	// CtaIpV6Dst - CTA_IP_V6_DST - represents the destination IPv6 address for a tracked connection within the
	// connection tracking system.
	CtaIpV6Dst = 4
)
