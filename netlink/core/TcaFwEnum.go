package core

// TcaFwEnum - Enumeration of traffic control filter actions for firewall marks.
//
// These enums represent the different attributes associated with firewall marks (fwmark) in Linux traffic control (tc).
// Firewall marks are often used in packet classification for network shaping and prioritization.
// The enumerations provide options to define class ID, policing actions, and interface device handling.
//
// For more details, refer to the following Linux kernel sources:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_fw.h
type TcaFwEnum uint8

const (
	// TcaFwUnspec - TCA_FW_UNSPEC - Unspecified action for firewall marks.
	TcaFwUnspec TcaFwEnum = iota

	// TcaFwClassID - TCA_FW_CLASSID - Class identifier for associating packets with a class.
	TcaFwClassID

	// TcaFwPolice - TCA_FW_POLICE - Defines the policing (rate-limiting) action on the packets.
	TcaFwPolice

	// TcaFwIndev - TCA_FW_INDEV - Input device associated with the traffic control action.
	TcaFwIndev

	// TcaFwAct - TCA_FW_ACT - Specifies the action to take when the filter matches.
	TcaFwAct

	// TcaFwMask - TCA_FW_MASK - Mask applied to the firewall mark.
	TcaFwMask

	// TcaFwMax - TCA_FW_MAX - Maximum valid value in the firewall mark filter.
	TcaFwMax = iota - 1
)
