package core

// DevlinkEswitchModeEnum - devlink_eswitch_mode
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkEswitchModeEnum uint8

const (
	// DevlinkEswitchModeLegacy - DEVLINK_ESWITCH_MODE_LEGACY
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchModeLegacy DevlinkEswitchModeEnum = 0

	// DevlinkEswitchModeSwitchdev - DEVLINK_ESWITCH_MODE_SWITCHDEV
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchModeSwitchdev DevlinkEswitchModeEnum = 1
)
