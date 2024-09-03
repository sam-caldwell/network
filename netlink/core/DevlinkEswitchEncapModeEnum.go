package core

// DevlinkEswitchEncapModeEnum - devlink_eswitch_encap_mode
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkEswitchEncapModeEnum uint8

const (

	// DevlinkEswitchEncapModeNone - DEVLINK_ESWITCH_ENCAP_MODE_NONE
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchEncapModeNone DevlinkEswitchEncapModeEnum = 0

	// DevlinkEswitchEncapModeBasic - DEVLINK_ESWITCH_ENCAP_MODE_BASIC
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchEncapModeBasic DevlinkEswitchEncapModeEnum = 1
)
