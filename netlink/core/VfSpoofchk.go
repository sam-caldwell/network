package core

// VfSpoofchk represents the ifla_vf_spoofchk structure from the Linux kernel.
//
// It contains the Virtual Function (VF) index and the spoof check setting.
//
//	See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type VfSpoofchk struct {

	// Vf - index
	Vf uint32

	// Setting - spoof-check setting
	Setting uint32
}
