//go:build linux

package core

// BridgeVlanInfo
//
//	struct bridge_vlan_info {
//	  __u16 flags;
//	  __u16 vid;
//	};
type BridgeVlanInfo struct {
	Flags uint16
	Vid   uint16
}
