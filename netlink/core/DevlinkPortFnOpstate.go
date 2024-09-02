package core

// DevlinkPortFnOpstate - devlink_port_fn_opstate - indicates operational state of the function
//
// @DEVLINK_PORT_FN_OPSTATE_ATTACHED: Driver is attached to the function.
//
//	For graceful tear down of the function, after inactivation of the
//	function, user should wait for operational state to turn DETACHED.
//
// @DEVLINK_PORT_FN_OPSTATE_DETACHED: Driver is detached from the function.
//
//	It is safe to delete the port.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkPortFnOpstate uint8

const (
	// DevlinkPortFnOpstateDetached - DEVLINK_PORT_FN_OPSTATE_DETACHED
	DevlinkPortFnOpstateDetached = 0

	// DevlinkPortFnOpstateAttached - DEVLINK_PORT_FN_OPSTATE_ATTACHED
	DevlinkPortFnOpstateAttached = 1
)
