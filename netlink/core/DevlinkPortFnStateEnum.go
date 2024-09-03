package core

// DevlinkPortFnStateEnum - devlink_port_fn_state
//
// The devlink_port_fn_state enum is used to indicate the operational state of a function associated with a devlink
// port. This could refer to a network function or virtual function on a physical port of a network device.
//
// It is useful in the management and configuration of network device ports, particularly in environments where
// features like SR-IOV are used.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkPortFnStateEnum uint8

const (

	// DevlinkPortFnStateInactive - DEVLINK_PORT_FN_STATE_INACTIVE - This state indicates that the function is
	// currently active. An active function is operational and can participate in network activities such as sending
	// or receiving packets.
	DevlinkPortFnStateInactive DevlinkPortFnStateEnum = 0

	// DevlinkPortFnStateActive - DEVLINK_PORT_FN_STATE_ACTIVE -
	// This state indicates that the function is currently inactive. In this state, the function is not operational
	// and does not participate in any network activity.
	DevlinkPortFnStateActive DevlinkPortFnStateEnum = 1
)
