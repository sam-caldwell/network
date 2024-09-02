package core

// DevlinkPortFunctionAttr - devlink_port_function_attr
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkPortFunctionAttr uint8

const (

	// DevlinkPortFunctionAttrUnspec - DEVLINK_PORT_FUNCTION_ATTR_UNSPEC
	DevlinkPortFunctionAttrUnspec DevlinkPortFunctionAttr = 0

	// DevlinkPortFunctionAttrHwAddr - DEVLINK_PORT_FUNCTION_ATTR_HW_ADDR
	DevlinkPortFunctionAttrHwAddr DevlinkPortFunctionAttr = 1 /* binary */

	// DevlinkPortFnAttrState - DEVLINK_PORT_FN_ATTR_STATE
	DevlinkPortFnAttrState DevlinkPortFunctionAttr = 2 /* u8 */

	// DevlinkPortFnAttrOpstate - DEVLINK_PORT_FN_ATTR_OPSTATE
	DevlinkPortFnAttrOpstate DevlinkPortFunctionAttr = 3 /* u8 */

	// DevlinkPortFnAttrCaps - DEVLINK_PORT_FN_ATTR_CAPS
	DevlinkPortFnAttrCaps DevlinkPortFunctionAttr = 4 /* bitfield32 */

	// DevlinkPortFnAttrDevlink - DEVLINK_PORT_FN_ATTR_DEVLINK
	DevlinkPortFnAttrDevlink DevlinkPortFunctionAttr = 5 /* nested */

	// DevlinkPortFnAttrMaxIoEqs - DEVLINK_PORT_FN_ATTR_MAX_IO_EQS
	DevlinkPortFnAttrMaxIoEqs DevlinkPortFunctionAttr = 6 /* u32 */

	/*
		add more values here.
	*/

	// __devlinkPortFunctionAttrMax - __DEVLINK_PORT_FUNCTION_ATTR_MAX
	__devlinkPortFunctionAttrMax = 7

	// DevlinkPortFunctionAttrMax - DEVLINK_PORT_FUNCTION_ATTR_MAX
	DevlinkPortFunctionAttrMax = __devlinkPortFunctionAttrMax - 1
)
