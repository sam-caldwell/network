package core

// IngressUntag checks whether the BridgeVlanInfo structure has the "untagged" flag set.
// This flag indicates whether incoming packets with a matching VLAN ID should have their
// VLAN tags removed (untagged) before being forwarded. This is particularly useful in VLAN
// configurations where you want to manage VLAN tags on ingress into a network bridge.
//
// The "untagged" flag is part of the BridgeVlanInfo's Flags field and is represented by the
// BridgeVlanInfoUntagged constant.
//
// Reference:
//   - BridgeVlanInfo structure as defined in the Linux kernel's "if_link.h":
//     https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
//
// Returns:
//   - true if the "untagged" flag is set in the Flags field, false otherwise.
func (bridge *BridgeVlanInfo) IngressUntag() bool {
	return bridge.Flags&BridgeVlanInfoUntagged > 0
}
