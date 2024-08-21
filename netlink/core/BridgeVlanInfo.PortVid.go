//go:build linux

package core

// PortVID - returns the VLAN ID assigned to a port in the bridge,
// determining which VLAN traffic is tagged for that port.
func (bridge *BridgeVlanInfo) PortVID() bool {

	return bridge.Flags&BRIDGE_VLAN_INFO_PVID > 0

}
