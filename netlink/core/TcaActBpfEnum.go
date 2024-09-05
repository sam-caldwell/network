package core

// TcaActBpfEnum - Enum for BPF-related actions in Traffic Control (TC) actions.
//
// This enumeration defines constants used for BPF (Berkeley Packet Filter) actions in Linux Traffic Control (TC).
// These actions help manage eBPF programs in the context of traffic control, allowing packet filtering and manipulation
// at various stages of the traffic control pipeline.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_bpf.h
type TcaActBpfEnum uint8

const (
	// TcaActBpfUnspec - TCA_ACT_BPF_UNSPEC - unspecified action.
	// Used as a placeholder for an unspecified or undefined BPF action.
	TcaActBpfUnspec TcaActBpfEnum = iota

	// TcaActBpfTm - TCA_ACT_BPF_TM - timing related information.
	// This attribute holds timing information related to BPF actions.
	TcaActBpfTm TcaActBpfEnum = iota

	// TcaActBpfParms - TCA_ACT_BPF_PARMS - parameters for BPF programs.
	// This attribute contains parameters associated with the BPF program used in TC actions.
	TcaActBpfParms TcaActBpfEnum = iota

	// TcaActBpfOpsLen - TCA_ACT_BPF_OPS_LEN - length of BPF ops.
	// This attribute provides the length of BPF operations.
	TcaActBpfOpsLen TcaActBpfEnum = iota

	// TcaActBpfOps - TCA_ACT_BPF_OPS - BPF operations.
	// This attribute represents the BPF operations that will be executed in the action.
	TcaActBpfOps TcaActBpfEnum = iota

	// TcaActBpfFd - TCA_ACT_BPF_FD - file descriptor.
	// This attribute holds the file descriptor for the BPF program.
	TcaActBpfFd TcaActBpfEnum = iota

	// TcaActBpfName - TCA_ACT_BPF_NAME - BPF program name.
	// This attribute contains the name of the BPF program used in the action.
	TcaActBpfName TcaActBpfEnum = iota

	// TcaActBpfMax - TCA_ACT_BPF_MAX - maximum BPF action.
	// Defines the maximum valid value for this enumeration.
	TcaActBpfMax = iota - 1
)

// TcaActBpf - TCA_ACT_BPF - constant representing BPF action type.
// This constant is used to identify BPF actions in Traffic Control (TC).
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_bpf.h
const (
	TcaActBpf = 13
)
