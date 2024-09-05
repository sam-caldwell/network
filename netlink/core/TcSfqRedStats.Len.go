package core

// Len - Returns the size of the TcSfqRedStats struct in bytes.
// This function ensures the size is calculated based on the struct definition.
//
// TcSfqRedStats holds statistics for SFQ-RED traffic management, and this function
// provides the size information required for serialization or handling of netlink
// messages in the Linux kernel traffic control system.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
func (msg *TcSfqRedStats) Len() int {
	return SizeofTcSfqRedStats
}
