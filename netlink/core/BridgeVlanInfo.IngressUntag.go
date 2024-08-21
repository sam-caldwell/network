//go:build linux

package core

// IngressUntag - flag within the BridgeVlanInfo structure that indicates whether incoming packets
// with a matching VLAN ID should have their VLAN tags removed (untagged) before being forwarded.
// This is commonly used to manage how VLAN-tagged traffic is handled when it enters a network bridge.
func (bridge *BridgeVlanInfo) IngressUntag() bool {

	return bridge.Flags&BRIDGE_VLAN_INFO_UNTAGGED > 0

}
