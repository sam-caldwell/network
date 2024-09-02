package core

// VfRate - ifla_vf_rate - represents the ifla_vf_rate structure from the Linux kernel.
//
// It contains information about the Virtual Function (VF) and its
// minimum and maximum transmission rates in Mbps.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type VfRate struct {

	// VF - index
	Vf uint32

	// MinTxRate - Minimum transmission rate in Mbps
	MinTxRate uint32

	// MaxTxRate - Maximum transmission rate in Mbps
	MaxTxRate uint32
}
