package core

// DevlinkParamTypeEnum - datatype used in a given devlink parameter
type DevlinkParamTypeEnum uint8

const (
	// DevlinkParamTypeU8 - DEVLINK_PARAM_TYPE_U8 - uint8
	DevlinkParamTypeU8 DevlinkParamTypeEnum = 1

	// DevlinkParamTypeU16 - DEVLINK_PARAM_TYPE_U16 - uint16
	DevlinkParamTypeU16 DevlinkParamTypeEnum = 2

	// DevlinkParamTypeU32 - DEVLINK_PARAM_TYPE_U32 - uint32
	DevlinkParamTypeU32 DevlinkParamTypeEnum = 3

	// DevlinkParamTypeString - DEVLINK_PARAM_TYPE_STRING - string
	DevlinkParamTypeString DevlinkParamTypeEnum = 5

	// DevlinkParamTypeBool - DEVLINK_PARAM_TYPE_BOOL - bool
	DevlinkParamTypeBool DevlinkParamTypeEnum = 6
)
