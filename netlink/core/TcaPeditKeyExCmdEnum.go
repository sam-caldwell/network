package core

// PeditCmd represents the command types used in the pedit action within the Linux kernel's traffic control (TC).
// These commands specify how packet headers should be modified using the extended pedit key feature.
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type PeditCmd uint16

const (
	// PeditCmdSet - TCA_PEDIT_KEY_EX_CMD_SET - represents the command to set a header field.
	PeditCmdSet PeditCmd = 0

	// PeditCmdAdd - TCA_PEDIT_KEY_EX_CMD_ADD - represents the command to add to a header field.
	PeditCmdAdd PeditCmd = 1
)
