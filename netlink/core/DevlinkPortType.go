package core

// DevlinkPortType - devlink_port_type
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkPortType uint8

const (

	// DevlinkPortTypeNotset - DEVLINK_PORT_TYPE_NOTSET
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortTypeNotset DevlinkPortType = 0

	// DevlinkPortTypeAuto - DEVLINK_PORT_TYPE_AUTO
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortTypeAuto DevlinkPortType = 1

	// DevlinkPortTypeEth - DEVLINK_PORT_TYPE_ETH
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortTypeEth DevlinkPortType = 2

	// DevlinkPortTypeIb - DEVLINK_PORT_TYPE_IB
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortTypeIb DevlinkPortType = 3
)
