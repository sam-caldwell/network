//go:build linux

package core

// BridgeVlanInfo represents a VLAN (Virtual Local Area Network) configuration associated with
// a network bridge in Linux systems. This structure corresponds to the `bridge_vlan_info`
// struct found in the Linux kernel headers.
//
// It is used primarily for managing VLANs on network bridges and is part of low-level network
// interface configuration through netlink. The Flags field provides additional VLAN options such
// as whether the VLAN is tagged, untagged, or the PVID (Primary VLAN ID). The Vid field represents
// the VLAN ID.
//
// In the Linux kernel, the structure can be found at:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
//
// C Struct Reference:
//
//	struct bridge_vlan_info {
//	  __u16 flags; // VLAN-specific flags (e.g., VLAN is untagged, PVID, etc.)
//	  __u16 vid;   // VLAN ID
//	};
//
// Fields:
//   - Flags: Represents VLAN-specific options. This could include whether the VLAN is tagged,
//     untagged, or set as the PVID.
//   - Vid: The VLAN ID (12 bits in the 802.1Q VLAN header).
type BridgeVlanInfo struct {
	Flags uint16 // VLAN-specific flags (e.g., untagged, PVID)
	Vid   uint16 // VLAN ID (range: 1-4094)
}
