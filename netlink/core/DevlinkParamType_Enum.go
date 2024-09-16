package core

// DevlinkParamType - Enum representing the data types used for devlink parameters.
// These types correspond to the parameter types defined in the Linux kernel for devlink
// (See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h).
//
// The devlink parameters provide a mechanism to manage and configure device settings
// dynamically. Each parameter has an associated type that defines the kind of value it holds.
type DevlinkParamType uint8

const (
	// DevlinkParamTypeU8 - DEVLINK_PARAM_TYPE_U8
	// Represents an unsigned 8-bit integer parameter (uint8).
	// Defined in the Linux kernel: https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkParamTypeU8 DevlinkParamType = 1

	// DevlinkParamTypeU16 - DEVLINK_PARAM_TYPE_U16
	// Represents an unsigned 16-bit integer parameter (uint16).
	// Defined in the Linux kernel: https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkParamTypeU16 DevlinkParamType = 2

	// DevlinkParamTypeU32 - DEVLINK_PARAM_TYPE_U32
	// Represents an unsigned 32-bit integer parameter (uint32).
	// Defined in the Linux kernel: https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkParamTypeU32 DevlinkParamType = 3

	// DevlinkParamTypeString - DEVLINK_PARAM_TYPE_STRING
	// Represents a string parameter.
	// Defined in the Linux kernel: https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkParamTypeString DevlinkParamType = 5

	// DevlinkParamTypeBool - DEVLINK_PARAM_TYPE_BOOL
	// Represents a boolean parameter.
	// Defined in the Linux kernel: https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkParamTypeBool DevlinkParamType = 6
)
