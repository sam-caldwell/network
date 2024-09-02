package core

// DevlinkParamType - datatype used in a given devlink parameter
type DevlinkParamType uint8

const (
	// DevlinkParamTypeU8 - DEVLINK_PARAM_TYPE_U8 - uint8
	DevlinkParamTypeU8 = 1

	// DevlinkParamTypeU16 - DEVLINK_PARAM_TYPE_U16 - uint16
	DevlinkParamTypeU16 = 2

	// DevlinkParamTypeU32 - DEVLINK_PARAM_TYPE_U32 - uint32
	DevlinkParamTypeU32 = 3

	// DevlinkParamTypeString - DEVLINK_PARAM_TYPE_STRING - string
	DevlinkParamTypeString = 5

	// DevlinkParamTypeBool - DEVLINK_PARAM_TYPE_BOOL - bool
	DevlinkParamTypeBool = 6
)
