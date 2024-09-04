package core

// Seg6Local - Parameters for SRv6 Local Actions.
// These constants define various parameters that can be used with SRv6 local actions in the Linux kernel.
//
// References:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
type Seg6Local struct{}

const (
	// Seg6LocalUnspec - SEG6_LOCAL_UNSPEC - unspecified local action parameter.
	Seg6LocalUnspec = iota

	// Seg6LocalAction - SEG6_LOCAL_ACTION - action to be executed.
	Seg6LocalAction

	// Seg6LocalSrh - SEG6_LOCAL_SRH - SRH (Segment Routing Header) for encapsulation.
	Seg6LocalSrh

	// Seg6LocalTable - SEG6_LOCAL_TABLE - lookup table for next-hop selection.
	Seg6LocalTable

	// Seg6LocalNh4 - SEG6_LOCAL_NH4 - next-hop IPv4 address.
	Seg6LocalNh4

	// Seg6LocalNh6 - SEG6_LOCAL_NH6 - next-hop IPv6 address.
	Seg6LocalNh6

	// Seg6LocalIif - SEG6_LOCAL_IIF - input interface index.
	Seg6LocalIif

	// Seg6LocalOif - SEG6_LOCAL_OIF - output interface index.
	Seg6LocalOif

	// Seg6LocalBpf - SEG6_LOCAL_BPF - BPF program used for custom SRv6 action.
	Seg6LocalBpf

	// Seg6LocalMax - SEG6_LOCAL_MAX - maximum value for SRv6 local parameters.
	Seg6LocalMax = iota - 1
)
