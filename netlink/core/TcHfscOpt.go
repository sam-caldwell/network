package core

// TcHfscOpt - Represents HFSC (Hierarchical Fair Service Curve) options in Linux Traffic Control.
//
// The HFSC scheduler allows precise control of bandwidth allocation using curves to specify
// bandwidth over time. This struct defines the default class for scheduling under the HFSC system.
//
// Reference:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcHfscOpt struct {
	// Defcls - Default class for scheduling.
	// Represents the default traffic class used by the HFSC scheduler.
	Defcls uint16
}
