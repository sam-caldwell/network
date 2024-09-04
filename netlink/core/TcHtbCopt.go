package core

// TcHtbCopt - Represents the options for an HTB (Hierarchical Token Bucket) class.
// This structure contains parameters used for traffic shaping with HTB qdisc.
//
// HTB allows bandwidth allocation and traffic prioritization. The parameters in this structure help
// define the rate, ceil (upper bandwidth limit), and other aspects of traffic control.
//
// For more information, see:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcHtbCopt struct {
	// Rate - Defines the guaranteed rate for this HTB class (TcRateSpec).
	Rate TcRateSpec

	// Ceil - Defines the upper rate (ceil) for this HTB class (TcRateSpec).
	// This is the maximum bandwidth the class can use.
	Ceil TcRateSpec

	// Buffer - Size of the buffer for the traffic class.
	// It allows burst traffic when the class isn't fully utilizing its rate.
	Buffer uint32

	// Cbuffer - Ceil buffer size.
	// Defines the burst size allowed when reaching the ceil rate.
	Cbuffer uint32

	// Quantum - The quantum of bytes, used in packet scheduling.
	// Larger quantum values tend to favor high throughput flows.
	Quantum uint32

	// Level - Defines the level of the class hierarchy (used for output only).
	// It indicates how deep the class is in the HTB tree structure.
	Level uint32

	// Prio - Defines the priority of the traffic class.
	// Lower values represent higher priority.
	Prio uint32
}
