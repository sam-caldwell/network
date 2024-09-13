package core

// HfscCopt - HFSC control options
//
// This struct defines control options for the Hierarchical Fair Service Curve (HFSC) qdisc,
// which is used to manage and shape network traffic with advanced service curve control.
// It consists of three service curves: Real-time Service Curve (Rsc), Fair Service Curve (Fsc),
// and Upper Service Curve (Usc). Each service curve is represented by a `Curve` structure,
// which defines the slope and time-based components of bandwidth control.
//
// Related sources:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
// - https://man7.org/linux/man-pages/man8/tc-hfsc.8.html
type HfscCopt struct {
	// Rsc: Real-time Service Curve, specifying guaranteed real-time bandwidth.
	Rsc Curve
	// Fsc: Fair Service Curve, providing fair bandwidth allocation among traffic flows.
	Fsc Curve
	// Usc: Upper Service Curve, limiting the maximum rate of bandwidth.
	Usc Curve
}
