package core

// TcaSkbEditEnum - Enumeration of SKB Edit Action Types
//
// This enumeration represents different action types related to editing SKB (Socket Buffer) fields in the Linux Traffic Control (tc) subsystem.
//
// The SKB Edit actions allow modifications to various packet-level fields like priority, mark, queue mapping, and packet type.
//
// For more information, refer to:
// - Linux Kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_skbedit.h
type TcaSkbEditEnum uint8

const (
	// TcaSkbEditUnspec - TCA_SKBEDIT_UNSPEC - Unspecified SKB Edit action.
	// Used as a placeholder for unspecified SKB edit actions.
	TcaSkbEditUnspec TcaSkbEditEnum = iota

	// TcaSkbEditTm - TCA_SKBEDIT_TM - Timestamp modification.
	// Used to edit the timestamp of the SKB.
	TcaSkbEditTm TcaSkbEditEnum = iota

	// TcaSkbEditParms - TCA_SKBEDIT_PARMS - Edit SKB parameters.
	// Specifies general parameters for SKB modification.
	TcaSkbEditParms TcaSkbEditEnum = iota

	// TcaSkbEditPriority - TCA_SKBEDIT_PRIORITY - Modify packet priority.
	// Allows changing the priority of the SKB.
	TcaSkbEditPriority TcaSkbEditEnum = iota

	// TcaSkbEditQueueMapping - TCA_SKBEDIT_QUEUE_MAPPING - Modify queue mapping.
	// Used to edit the queue mapping of the SKB for packet scheduling.
	TcaSkbEditQueueMapping TcaSkbEditEnum = iota

	// TcaSkbEditMark - TCA_SKBEDIT_MARK - Modify packet mark.
	// Allows changing the mark value of the SKB for firewall or routing purposes.
	TcaSkbEditMark TcaSkbEditEnum = iota

	// TcaSkbEditPad - TCA_SKBEDIT_PAD - Padding for alignment.
	// Adds padding to ensure the alignment of SKB data.
	TcaSkbEditPad TcaSkbEditEnum = iota

	// TcaSkbEditPtype - TCA_SKBEDIT_PTYPE - Modify packet type.
	// Changes the packet type, for example, from IPv4 to IPv6.
	TcaSkbEditPtype TcaSkbEditEnum = iota

	// TcaSkbEditMask - TCA_SKBEDIT_MASK - Apply a mask to modify specific fields.
	// Used to apply masks on SKB fields for selective modification.
	TcaSkbEditMask TcaSkbEditEnum = iota

	// TcaSkbEditMax - TCA_SKBEDIT_MAX - Maximum value for SKB Edit actions.
	TcaSkbEditMax = iota - 1
)
