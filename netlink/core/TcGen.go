package core

// TcGen - Represents the generic part of traffic control actions (tc_gen).
//
// This structure holds metadata for traffic control actions used in the Linux kernel's traffic control subsystem.
// These fields help track the status and capabilities of specific traffic control actions.
//
// For more details, see:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_common.h
type TcGen struct {
	// Index - Action's index number
	Index uint32
	// Capab - Capability flags for the action.
	Capab uint32
	// Action - type (e.g., TC_ACT_* constants)
	Action int32
	// Refcnt - Reference count (how many places use this action)
	Refcnt int32
	// Bindcnt - Count of filters bound to this action. The count of how many filters are bound to this action.
	Bindcnt int32
}
