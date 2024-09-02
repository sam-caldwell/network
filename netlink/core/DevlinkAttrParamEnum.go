package core

// DevlinkAttrParamEnum - enumerated type for devlink attribute parameters
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkAttrParamEnum uint8

const (
	// DevlinkAttrParam - DEVLINK_ATTR_PARAM - nested
	DevlinkAttrParam DevlinkAttrParamEnum = 80

	// DevlinkAttrParamName - DEVLINK_ATTR_PARAM_NAME - string
	DevlinkAttrParamName DevlinkAttrParamEnum = 81

	// DevlinkAttrParamGeneric - DEVLINK_ATTR_PARAM_GENERIC - flag
	DevlinkAttrParamGeneric DevlinkAttrParamEnum = 82

	// DevlinkAttrParamType - DEVLINK_ATTR_PARAM_TYPE - u8
	DevlinkAttrParamType DevlinkAttrParamEnum = 83

	// DevlinkAttrParamValuesList - DEVLINK_ATTR_PARAM_VALUES_LIST - nested
	DevlinkAttrParamValuesList DevlinkAttrParamEnum = 84

	// DevlinkAttrParamValue - DEVLINK_ATTR_PARAM_VALUE - nested
	DevlinkAttrParamValue DevlinkAttrParamEnum = 85

	// DevlinkAttrParamValueData - DEVLINK_ATTR_PARAM_VALUE_DATA - dynamic
	DevlinkAttrParamValueData DevlinkAttrParamEnum = 86

	// DevlinkAttrParamValueCmode - DEVLINK_ATTR_PARAM_VALUE_CMODE - u8
	DevlinkAttrParamValueCmode DevlinkAttrParamEnum = 87
)
