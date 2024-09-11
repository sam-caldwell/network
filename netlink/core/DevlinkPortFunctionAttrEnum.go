package core

// DevlinkPortFunctionAttrEnum - devlink_port_function_attr
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkPortFunctionAttrEnum uint8

const (

	// DevlinkPortFunctionAttrUnspec - DEVLINK_PORT_FUNCTION_ATTR_UNSPEC
	DevlinkPortFunctionAttrUnspec DevlinkPortFunctionAttrEnum = 0

	// DevlinkPortFunctionAttrHwAddr - DEVLINK_PORT_FUNCTION_ATTR_HW_ADDR
	DevlinkPortFunctionAttrHwAddr DevlinkPortFunctionAttrEnum = 1 /* binary */

	// DevlinkPortFnAttrState - DEVLINK_PORT_FN_ATTR_STATE
	DevlinkPortFnAttrState DevlinkPortFunctionAttrEnum = 2 /* u8 */

	// DevlinkPortFnAttrOpstate - DEVLINK_PORT_FN_ATTR_OPSTATE
	DevlinkPortFnAttrOpstate DevlinkPortFunctionAttrEnum = 3 /* u8 */

	// DevlinkPortFnAttrCaps - DEVLINK_PORT_FN_ATTR_CAPS
	DevlinkPortFnAttrCaps DevlinkPortFunctionAttrEnum = 4 /* bitfield32 */

	// DevlinkPortFnAttrDevlink - DEVLINK_PORT_FN_ATTR_DEVLINK
	DevlinkPortFnAttrDevlink DevlinkPortFunctionAttrEnum = 5 /* nested */

	// DevlinkPortFnAttrMaxIoEqs - DEVLINK_PORT_FN_ATTR_MAX_IO_EQS
	DevlinkPortFnAttrMaxIoEqs DevlinkPortFunctionAttrEnum = 6 /* u32 */

	/*
		add more values here.
	*/

	// DevlinkPortFunctionAttrMax - DEVLINK_PORT_FUNCTION_ATTR_MAX
	DevlinkPortFunctionAttrMax = iota - 1
)
