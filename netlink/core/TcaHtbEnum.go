package core

// TcbHtbEnum - Enumeration for HTB (Hierarchical Token Bucket) message attributes.
// This enum represents various attributes related to HTB qdisc operations in Linux Traffic Control (TC).
//
// HTB is used to shape traffic by enforcing bandwidth limits on packets.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcbHtbEnum uint8

const (
	// TcaHtbUnspec - TCA_HTB_UNSPEC - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute.
	TcaHtbUnspec TcbHtbEnum = iota

	// TcaHtbParms - TCA_HTB_PARMS - Parameters for HTB class.
	// This attribute defines the basic parameters for the HTB class, such as rate and ceil.
	TcaHtbParms TcbHtbEnum = iota

	// TcaHtbInit - TCA_HTB_INIT - Initial settings for HTB class.
	// This attribute represents the initial configuration settings for the HTB class.
	TcaHtbInit TcbHtbEnum = iota

	// TcaHtbCtab - TCA_HTB_CTAB - Controlling table for HTB.
	// This attribute refers to the controlling table used by HTB.
	TcaHtbCtab TcbHtbEnum = iota

	// TcaHtbRtab - TCA_HTB_RTAB - Rate table for HTB.
	// The rate table contains the data required for shaping the traffic at specific rates.
	TcaHtbRtab TcbHtbEnum = iota

	// TcaHtbDirectQlen - TCA_HTB_DIRECT_QLEN - Direct queue length for HTB.
	// This attribute specifies the length of the direct queue in HTB.
	TcaHtbDirectQlen TcbHtbEnum = iota

	// TcaHtbRate64 - TCA_HTB_RATE64 - 64-bit rate for HTB.
	// This attribute provides support for 64-bit rate values for bandwidth enforcement.
	TcaHtbRate64 TcbHtbEnum = iota

	// TcaHtbCeil64 - TCA_HTB_CEIL64 - 64-bit ceil rate for HTB.
	// This attribute provides support for 64-bit ceil rate values for limiting traffic peaks.
	TcaHtbCeil64 TcbHtbEnum = iota

	// TcaHtbMax - TCA_HTB_MAX - Maximum attribute value for HTB.
	// This constant represents the maximum valid value for the HTB attributes.
	TcaHtbMax = iota - 1
)
