package core

// TcHtbGlob - Global settings for the HTB (Hierarchical Token Bucket) qdisc.
// This structure holds configuration data for the HTB qdisc, which controls traffic shaping and prioritization.
//
// HTB is a popular qdisc in the Linux kernel for controlling traffic and bandwidth allocation.
//
// The "quantum" in HTB represents the smallest unit of traffic (in bytes) that can be sent in one round of traffic
// scheduling. It is typically derived from the data rate and determines how much data can be dequeued in a single
// transmission opportunity.
//
// The quantum ensures fair distribution of bandwidth among different traffic classes by ensuring that each class
// gets a minimum amount of data transmitted in each round, based on its assigned rate.
//
// For more details, refer to the Linux kernel source code:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
// - https://man7.org/linux/man-pages/man8/tc-htb.8.html
type TcHtbGlob struct {
	// Version - HTB version.
	Version uint32
	// Rate2Quantum - Conversion factor from rate (bps) to quantum (unit of traffic control).
	Rate2Quantum uint32
	// Defcls - Default traffic class.
	Defcls uint32
	// Debug - Debugging flags.
	Debug uint32
	// DirectPkts - Counter for direct packets processed.
	DirectPkts uint32
}
