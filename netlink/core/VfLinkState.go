package core

// VfLinkState represents the ifla_vf_link_state structure from the Linux kernel.
//
// It contains the Virtual Function (VF) index and the link state.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type VfLinkState struct {

	// Vf - index
	Vf uint32

	// LinkState - Link-state setting
	LinkState uint32
}
