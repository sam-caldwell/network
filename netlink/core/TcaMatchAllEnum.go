package core

// TcaMatchAllEnum - Enumeration of match-all filter attributes in Linux traffic control.
//
// These enums represent the various attributes for the match-all filter used in Linux traffic control (tc).
// The match-all filter matches all packets, which makes it useful for applying actions such as policing or marking to all traffic.
// These attributes help define the associated class ID, actions, and flags.
//
// For further details, refer to the following Linux kernel sources:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_matchall.h
type TcaMatchAllEnum uint8

const (
	// TcaMatchAllUnspec - TCA_MATCHALL_UNSPEC - Unspecified match-all attribute.
	TcaMatchAllUnspec TcaMatchAllEnum = iota

	// TcaMatchAllClassID - TCA_MATCHALL_CLASSID - Class identifier for associating packets with a class.
	TcaMatchAllClassID

	// TcaMatchAllAct - TCA_MATCHALL_ACT - Specifies the action to take when the filter matches.
	TcaMatchAllAct

	// TcaMatchAllFlags - TCA_MATCHALL_FLAGS - Flags associated with the match-all filter.
	TcaMatchAllFlags
)
