package core

// DevlinkEswitchInlineModeEnum - devlink_eswitch_inline_mode
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkEswitchInlineModeEnum uint8

const (
	// DevlinkEswitchInlineModeNone - DEVLINK_ESWITCH_INLINE_MODE_NONE
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchInlineModeNone DevlinkEswitchInlineModeEnum = 0

	// DevlinkEswitchInlineModeLink - DEVLINK_ESWITCH_INLINE_MODE_LINK
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchInlineModeLink DevlinkEswitchInlineModeEnum = 1

	// DevlinkEswitchInlineModeNetwork - DEVLINK_ESWITCH_INLINE_MODE_NETWORK
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchInlineModeNetwork DevlinkEswitchInlineModeEnum = 2

	// DevlinkEswitchInlineModeTransport - DEVLINK_ESWITCH_INLINE_MODE_TRANSPORT
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchInlineModeTransport DevlinkEswitchInlineModeEnum = 3
)
