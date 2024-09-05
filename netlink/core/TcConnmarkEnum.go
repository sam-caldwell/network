package core

// TcConnmarkEnum - Enumeration of possible parameters for tc_connmark actions.
// These constants represent the different options used when configuring a tc_connmark action in the Linux Traffic Control system.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_connmark.h
type TcConnmarkEnum uint8

const (
	// TcaConnmarkUnspec - TCA_CONNMARK_UNSPEC - Unspecified parameter.
	TcaConnmarkUnspec TcConnmarkEnum = iota

	// TcaConnmarkParms - TCA_CONNMARK_PARMS - Parameters for connmark actions.
	// This is used to specify connmark-specific parameters.
	TcaConnmarkParms TcConnmarkEnum = iota

	// TcaConnmarkTm - TCA_CONNMARK_TM - Timing data for connmark actions.
	// This parameter holds timing information associated with the action.
	TcaConnmarkTm TcConnmarkEnum = iota

	// TcaConnmarkMax - TCA_CONNMARK_MAX - Maximum defined connmark parameter.
	// This constant represents the highest value in the enumeration and is used for boundary checking.
	TcaConnmarkMax = iota - 1
)
