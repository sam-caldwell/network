package core

// DevlinkEswitchInlineMode - devlink_eswitch_inline_mode
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkEswitchInlineMode uint8

const (
	// DevlinkEswitchInlineModeNone - DEVLINK_ESWITCH_INLINE_MODE_NONE
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchInlineModeNone DevlinkEswitchInlineMode = 0

	// DevlinkEswitchInlineModeLink - DEVLINK_ESWITCH_INLINE_MODE_LINK
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchInlineModeLink DevlinkEswitchInlineMode = 1

	// DevlinkEswitchInlineModeNetwork - DEVLINK_ESWITCH_INLINE_MODE_NETWORK
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchInlineModeNetwork DevlinkEswitchInlineMode = 2

	// DevlinkEswitchInlineModeTransport - DEVLINK_ESWITCH_INLINE_MODE_TRANSPORT
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkEswitchInlineModeTransport DevlinkEswitchInlineMode = 3
)
