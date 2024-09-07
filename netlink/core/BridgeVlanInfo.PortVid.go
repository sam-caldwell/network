package core

// PortVID checks if the Port VLAN ID (PVID) flag is set within the BridgeVlanInfo structure.
// The PVID indicates whether a port should associate untagged incoming packets with the VLAN ID
// defined in the VID field of the BridgeVlanInfo struct. This is typically used in network bridges
// that support VLANs to handle how untagged traffic is processed.
//
// Return:
//   - bool: true if the PVID flag is set (meaning the port is configured with a default VLAN ID for untagged packets),
//     false otherwise.
func (bridge *BridgeVlanInfo) PortVID() bool {
	return bridge.Flags&BridgeVlanInfoPvid > 0
}
