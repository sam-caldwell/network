package core

// BpfProgPolicy - bpf_prog_policy - represents the BPF program policies associated with Lightweight Tunnels (LWT).
// These constants correspond to values used in the Linux kernel for managing BPF programs in LWT contexts.
// See https://github.com/torvalds/linux/blob/master/net/core/lwt_bpf.c
type BpfProgPolicy uint8

const (

	// LwtBpfProgUnspec - LWT_BPF_PROG_UNSPEC - constant in the Linux kernel, specifically related to Lightweight
	// Tunnels (LWT) and BPF (Berkeley Packet Filter) programs. It is used as an identifier for an unspecified BPF
	// program associated with a lightweight tunnel. In this context, "unspecified" typically means that no specific
	// BPF program is attached or that the BPF program type is undefined.
	LwtBpfProgUnspec BpfProgPolicy = 0

	// LwtBpfProgFd - LWT_BPF_PROG_FD - constant in the Linux kernel used in the context of Lightweight Tunnels (LWT)
	// and Berkeley Packet Filter (BPF) programs. It represents a file descriptor (FD) associated with a BPF program
	// that is attached to a lightweight tunnel. This file descriptor is used to reference and manage the specific
	// BPF program that will be executed in the context of the lightweight tunnel for packet processing tasks such as
	// filtering, classification, or manipulation.
	LwtBpfProgFd BpfProgPolicy = 1

	// LwtBpfProgName - LWT_BPF_PROG_NAME - constant in the Linux kernel that represents the name of a BPF (Berkeley
	// Packet Filter) program associated with a Lightweight Tunnel (LWT). This constant is used to specify or identify
	// the BPF program by its name, allowing the kernel or user-space tools to reference and manage the BPF program
	// associated with the tunnel based on its human-readable identifier rather than a file descriptor or other numeric
	// identifier. This can be particularly useful in scenarios where multiple BPF programs are in use and need to be
	// distinguished by name.
	LwtBpfProgName BpfProgPolicy = 2

	// LwtBpfProgMax - LWT_BPF_PROG_MAX - constant in the Linux kernel that defines the maximum valid value for BPF
	// (Berkeley Packet Filter) program types or attributes associated with Lightweight Tunnels (LWT). It sets the
	// upper boundary for the enumeration of BPF program-related attributes, ensuring that any indexing or iteration
	// over BPF program types stays within this defined range. Unlike __LWT_BPF_PROG_MAX, which is typically an
	// internal marker, LWT_BPF_PROG_MAX might be used in code to refer to the last valid BPF program attribute in
	// this enumeration.
	LwtBpfProgMax BpfProgPolicy = iota - 1
)
