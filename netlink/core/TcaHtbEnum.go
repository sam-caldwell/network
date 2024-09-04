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
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbUnspec TcbHtbEnum = iota

	// TcaHtbParms - TCA_HTB_PARMS - Parameters for HTB class.
	// This attribute defines the basic parameters for the HTB class, such as rate and ceil.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbParms TcbHtbEnum = iota

	// TcaHtbInit - TCA_HTB_INIT - Initial settings for HTB class.
	// This attribute represents the initial configuration settings for the HTB class.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbInit TcbHtbEnum = iota

	// TcaHtbCtab - TCA_HTB_CTAB - Controlling table for HTB.
	// This attribute refers to the controlling table used by HTB.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbCtab TcbHtbEnum = iota

	// TcaHtbRtab - TCA_HTB_RTAB - Rate table for HTB.
	// The rate table contains the data required for shaping the traffic at specific rates.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbRtab TcbHtbEnum = iota

	// TcaHtbDirectQlen - TCA_HTB_DIRECT_QLEN - Direct queue length for HTB.
	// This attribute specifies the length of the direct queue in HTB.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbDirectQlen TcbHtbEnum = iota

	// TcaHtbRate64 - TCA_HTB_RATE64 - 64-bit rate for HTB.
	// This attribute provides support for 64-bit rate values for bandwidth enforcement.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbRate64 TcbHtbEnum = iota

	// TcaHtbCeil64 - TCA_HTB_CEIL64 - 64-bit ceil rate for HTB.
	// This attribute provides support for 64-bit ceil rate values for limiting traffic peaks.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbCeil64 TcbHtbEnum = iota

	// TcaHtbMax - TCA_HTB_MAX - Maximum attribute value for HTB.
	// This constant represents the maximum valid value for the HTB attributes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaHtbMax = iota - 1
)
