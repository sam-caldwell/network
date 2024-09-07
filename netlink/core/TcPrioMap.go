package core

// TcPriorityMap represents the traffic control priority queueing option.
//
// This struct mirrors the `tc_prio_qopt` structure in the Linux kernel's
// traffic control (tc) system. It is used to configure priority queueing
// settings for traffic classes.
//
// The `Bands` field specifies the number of priority bands, while the `Priomap`
// array maps priorities to those bands.
//
// References:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcPriorityMap struct {
	Bands   int32                // Number of priority bands.
	Priomap [TcPrioMax + 1]uint8 // Priority mapping to bands (0-15).
}
