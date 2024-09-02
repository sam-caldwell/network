package core

// DevlinkEswitchMode - devlink_eswitch_mode
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkEswitchMode uint8

const (
	// DevlinkEswitchModeLegacy - DEVLINK_ESWITCH_MODE_LEGACY
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchModeLegacy = 0

	// DevlinkEswitchModeSwitchdev - DEVLINK_ESWITCH_MODE_SWITCHDEV
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchModeSwitchdev = 1
)
