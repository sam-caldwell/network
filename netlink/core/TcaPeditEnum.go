package core

// TcaPeditEnum represents the `TCA_PEDIT` attributes used in the Linux kernel's traffic control (TC) pedit action.
// The pedit action is used to modify packet headers, and these enum values define various pedit attributes.
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type TcaPeditEnum uint16

const (
	// TcaPeditUnspec - TCA_PEDIT_UNSPEC - represents an unspecified pedit action.
	TcaPeditUnspec TcaPeditEnum = iota

	// TcaPeditTm - TCA_PEDIT_TM - represents a pedit time action.
	TcaPeditTm TcaPeditEnum = iota

	// TcaPeditParms - TCA_PEDIT_PARMS - represents the pedit parameters.
	TcaPeditParms TcaPeditEnum = iota

	// TcaPeditPad - TCA_PEDIT_PAD - represents the padding for pedit.
	TcaPeditPad TcaPeditEnum = iota

	// TcaPeditParmsEx - TcaPeditParmsEx - represents extended pedit parameters.
	TcaPeditParmsEx TcaPeditEnum = iota

	// TcaPeditKeysEx - TcaPeditKeysEx - represents extended pedit keys.
	TcaPeditKeysEx TcaPeditEnum = iota

	// TcaPeditKeyEx - TcaPeditKeyEx - represents a single extended pedit key.
	TcaPeditKeyEx TcaPeditEnum = iota
)
