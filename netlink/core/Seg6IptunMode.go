package core

// Seg6IptunMode - Represents the SEG6 IP tunneling mode in IPv6 Segment Routing Header (SRH).
//
// These modes define how Segment Routing for IPv6 (SRv6) encapsulation is performed, either inline or encapsulated.
// The modes are utilized in network routing and tunneling configurations to manage packet forwarding behavior using segment routing.
//
// For further details, see:
// - Linux Kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6.h
type Seg6IptunMode uint8

const (
	// Seg6IptunModeInline - SEG6_IPTUN_MODE_INLINE -
	// Indicates that the segment routing header is processed inline without encapsulation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6.h
	Seg6IptunModeInline Seg6IptunMode = 0

	// Seg6IptunModeEncap - SEG6_IPTUN_MODE_ENCAP -
	// Indicates that the segment routing header is used for encapsulation of the original packet.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6.h
	Seg6IptunModeEncap Seg6IptunMode = 1
)
