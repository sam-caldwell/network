package core

// IflaVfLinkStateEnum - Enumeration for Virtual Function (VF) link state attributes
//
// This enumeration defines various attributes related to the link state of Virtual Functions (VFs).
// The link state controls whether the VF's link is considered up, down, or follows the state of the physical uplink.
type IflaVfLinkStateEnum uint8

const (
	// IflaVfLinkStateAuto - IFLA_VF_LINK_STATE_AUTO -
	// This attribute sets the VF's link state to follow the state of the uplink.
	// The VF's link will be considered up or down based on the state of the physical network interface it is attached to.
	IflaVfLinkStateAuto IflaVfLinkStateEnum = 0

	// IflaVfLinkStateEnable - IFLA_VF_LINK_STATE_ENABLE -
	// This attribute forces the VF's link to always be up, regardless of the state of the physical uplink.
	// This is useful in scenarios where the VF must always be operational.
	IflaVfLinkStateEnable IflaVfLinkStateEnum = 1

	// IflaVfLinkStateDisable - IFLA_VF_LINK_STATE_DISABLE -
	// This attribute forces the VF's link to always be down, regardless of the state of the physical uplink.
	// This is useful in scenarios where the VF should not be allowed to transmit or receive traffic.
	IflaVfLinkStateDisable IflaVfLinkStateEnum = 2

	// IflaVfLinkStateMax - IFLA_VF_LINK_STATE_MAX -
	// This constant represents the maximum valid value for VF link state attributes.
	// It is set to the last valid link state (IflaVfLinkStateDisable) to serve as a boundary marker.
	IflaVfLinkStateMax IflaVfLinkStateEnum = IflaVfLinkStateDisable
)
