package core

// DevlinkEswitchEncapMode - devlink_eswitch_encap_mode
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkEswitchEncapMode uint8

const (

	// DevlinkEswitchEncapModeNone - DEVLINK_ESWITCH_ENCAP_MODE_NONE
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchEncapModeNone DevlinkEswitchEncapMode = 0

	// DevlinkEswitchEncapModeBasic - DEVLINK_ESWITCH_ENCAP_MODE_BASIC
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchEncapModeBasic DevlinkEswitchEncapMode = 1
)
