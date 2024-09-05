package core

const (
	// TcaBpfFlagActDirect - TCA_BPF_FLAG_ACT_DIRECT -
	// This flag is used to specify a direct action in the BPF (Berkeley Packet Filter) traffic control action.
	// When set, it indicates that the action should be applied directly to the packet without further processing.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_bpf.h
	TcaBpfFlagActDirect uint32 = 1 << iota
)
