package core

// BpfNetlinkPolicy - bpf_nl_policy - defines BPF policy.
type BpfNetlinkPolicy uint8

const (

	// LwtBpfUnspec - LWT_BPF_UNSPEC - constant in the Linux kernel's networking subsystem that represents an
	// unspecified BPF (Berkeley Packet Filter) program type associated with Lightweight Tunnels (LWT). This
	// constant is used as an identifier when no specific BPF program is attached to a lightweight tunnel, or when
	// the BPF program type is undefined or not applicable.
	LwtBpfUnspec BpfNetlinkPolicy = 0

	// LwtBpfIn - LWT_BPF_IN - constant in the Linux kernel that represents a BPF (Berkeley Packet Filter) program
	// that is applied to incoming packets in the context of Lightweight Tunnels (LWT). This constant is used to
	// specify that the associated BPF program should be executed when a packet is received (ingress) through the
	// lightweight tunnel.
	LwtBpfIn BpfNetlinkPolicy = 1

	// LwtBpfOut - LWT_BPF_OUT -constant in the Linux kernel that represents a BPF program applied to outgoing packets
	// (egress) in the context of Lightweight Tunnels (LWT).  This constant is used to specify that the associated BPF
	// program should be executed when a packet is transmitted (egress) through the lightweight tunnel.
	LwtBpfOut BpfNetlinkPolicy = 2

	// LwtBpfXmit - LWT_BPF_XMIT - constant in the Linux kernel that represents a BPF program applied during the
	// transmission phase of packets within a Lightweight Tunnel (LWT).
	LwtBpfXmit BpfNetlinkPolicy = 3

	// LwtBpfXmitHeadroom - LWT_BPF_XMIT_HEADROOM - constant in the Linux kernel that represents the additional
	// headroom reserved for a BPF program during the packet transmission phase within a Lightweight Tunnel (LWT),
	// allowing the BPF program to modify or add headers without reallocating the packet buffer.
	LwtBpfXmitHeadroom BpfNetlinkPolicy = 4

	// LwtBpfMax - LWT_BPF_MAX - constant in the Linux kernel that represents the highest valid value for BPF program
	// types associated with Lightweight Tunnels (LWT). It is derived from __LWT_BPF_MAX and is typically used in the
	// code to refer to the last permissible BPF program type in the enumeration.
	LwtBpfMax BpfNetlinkPolicy = iota - 1
)
