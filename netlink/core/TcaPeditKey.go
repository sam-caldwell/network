package core

// TcaPeditKey represents the `TCA_PEDIT_KEY_EX` attribute used in traffic control (TC) in the Linux kernel.
// The pedit action is used to edit packet headers in the kernel. The `TCA_PEDIT_KEY_EX_HTYPE` and
// `TCA_PEDIT_KEY_EX_CMD` constants refer to specific types of key actions within the pedit functionality.
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type TcaPeditKey uint8

const (
	// TcaPeditKeyExHtype - TCA_PEDIT_KEY_EX_HTYPE - represents the header type key for pedit.
	TcaPeditKeyExHtype TcaPeditKey = 1

	// TcaPeditKeyExCmd - TCA_PEDIT_KEY_EX_CMD - represents the command key for pedit.
	TcaPeditKeyExCmd TcaPeditKey = 2
)
