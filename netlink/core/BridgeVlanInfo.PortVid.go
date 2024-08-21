//go:build linux

package core

func (b *BridgeVlanInfo) PortVID() bool {

	return b.Flags&BRIDGE_VLAN_INFO_PVID > 0

}
