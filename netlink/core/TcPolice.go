package core

// TcPolice - tc_police - Structure representing the policing settings in the Linux Traffic Control (tc) subsystem.
//
// This struct is used to define traffic policing rules that specify limits on packet transmission rates.
// It allows setting burst limits, rate limits, and MTU constraints for traffic shaping.
//
// Related kernel source code:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcPolice struct {

	// Index - Unique index for the policing rule
	Index uint32

	// Action - Action to perform (drop, pass, etc.)
	Action int32

	// Limit - Maximum allowed bytes
	Limit uint32

	// Burst - Maximum allowed burst size
	Burst uint32

	// Mtu - Maximum transmission unit (MTU)
	Mtu uint32

	// Rate - Rate specification for allowed traffic
	Rate TcRateSpec

	// PeakRate - Peak rate specification
	PeakRate TcRateSpec

	// Refcnt - Reference count
	Refcnt int32

	// Bindcnt - Bind count for the policy
	Bindcnt int32

	// Capab - Capabilities and flags for policing
	Capab uint32
}
