package core

// VethInfo - Port of the C veth.h Linux kernel file for the Virtual Ethernet (veth) network device.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/veth.h
type VethInfo uint8

const (

	// VethInfoUnspec - VETH_INFO_UNSPEC - unspecified enum value (0)
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/veth.h
	VethInfoUnspec VethInfo = 0

	// VethInfoPeer - VETH_INFO_PEER - Represents an attribute related to the "peer" of a veth interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/veth.h
	VethInfoPeer VethInfo = 1

	// VethInfoMax - VETH_INFO_MAX - A helper constant that is used to define the maximum value in the enumeration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/veth.h
	VethInfoMax = 1
)
