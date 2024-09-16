package core

// DevlinkParamCmode - devlink_param_cmode - enumerated values for DEVLINK_PARAM_CMODE_*
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkParamCmode uint8

const (
	// DevlinkParamCmodeRuntime - DEVLINK_PARAM_CMODE_RUNTIME
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkParamCmodeRuntime DevlinkParamCmode = 0

	// DevlinkParamCmodeDriverinit - DEVLINK_PARAM_CMODE_DRIVERINIT
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkParamCmodeDriverinit DevlinkParamCmode = 1

	// DevlinkParamCmodePermanent - DEVLINK_PARAM_CMODE_PERMANENT
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkParamCmodePermanent DevlinkParamCmode = 2

	/* Add new configuration modes above */

	// DevlinkParamCmodeMax - DEVLINK_PARAM_CMODE_MAX
	DevlinkParamCmodeMax = iota - 1
)
