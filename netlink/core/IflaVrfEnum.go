package core

// IflaVrfEnum -
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaVrfEnum uint8

const (
	// IflaVrfUnspec - IFLA_VRF_UNSPEC -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVrfUnspec = iota

	// IflaVrfTable - IFLA_VRF_TABLE -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVrfTable
)
