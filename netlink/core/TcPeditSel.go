package core

// TcPeditSel represents the selection structure used in the pedit action within the Linux kernel's traffic control (TC).
// This struct is used to define how many keys are involved in a packet editing operation and any relevant flags.
//
// TcGen is the base structure for general TC actions, which provides common fields used by various actions.
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type TcPeditSel struct {
	// TcGen - provides general fields for the pedit action.
	TcGen

	// NKeys - represents the number of keys used for editing the packet header.
	NKeys uint8

	// Flags - represents additional flags for the pedit operation.
	Flags uint8
}
