//go:build linux

package core

// Deserialize - Safely deserialize a byte slice into a BridgeVlanInfo structure by calling the
// DeserializeBridgeVlanInfo() helper function.
//
// This method allows you to populate an existing BridgeVlanInfo instance with the deserialized data
// from the provided byte slice. If the input byte slice is too short or invalid, it returns an error.
//
// The BridgeVlanInfo struct represents the VLAN configuration associated with a network bridge,
// with fields such as Flags and Vid (VLAN ID).
//
// Parameters:
// - b: A byte slice that contains the serialized BridgeVlanInfo data. It should be at least 4 bytes long.
//
// Returns:
// - err: An error if the deserialization fails (e.g., if the byte slice is too short), otherwise nil.
//
// Example:
// ```go
// var bridge BridgeVlanInfo
// err := bridge.Deserialize([]byte{0x01, 0x00, 0x05, 0x00})
//
//	if err != nil {
//	    fmt.Println("Deserialization failed:", err)
//	} else {
//
//	    fmt.Printf("Deserialized BridgeVlanInfo: Flags=%d, VID=%d\n", bridge.Flags, bridge.Vid)
//	}
//
// ```
//
// Reference:
// The BridgeVlanInfo structure corresponds to the `bridge_vlan_info` structure in the Linux kernel's
// `if_bridge.h` header file, commonly used for managing VLAN configurations in network bridges.
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
func (bridge *BridgeVlanInfo) Deserialize(b []byte) (err error) {
	result, err := DeserializeBridgeVlanInfo(b)
	if err != nil {
		return err
	}
	*bridge = *result
	return nil
}
