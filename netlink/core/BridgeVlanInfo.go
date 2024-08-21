//go:build linux

package core

// BridgeVlanInfo - represents a specific VLAN (Virtual Local Area Network) configuration associated with
// a network bridge. This structure is typically used in low-level networking operations on Linux systems,
// particularly when interacting with network interfaces and managing VLANs.
//
//	struct bridge_vlan_info {
//	  __u16 flags;
//	  __u16 vid;
//	};
type BridgeVlanInfo struct {
	Flags uint16
	Vid   uint16
}
