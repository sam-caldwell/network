//go:build linux

package core

func (b *BridgeVlanInfo) IngressUntag() bool {

	return b.Flags&BRIDGE_VLAN_INFO_UNTAGGED > 0
	
}
