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
//
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

// DeserializeBridgeVlanInfo safely deserializes a byte slice into a BridgeVlanInfo object.
//
// This function reads the provided byte slice and interprets it as a serialized `bridge_vlan_info`
// structure, converting the byte slice back into a `BridgeVlanInfo` Go struct. It also validates
// the byte slice length to ensure that it can properly deserialize the structure.
//
// C Reference:
// The `bridge_vlan_info` structure in the Linux kernel is defined as:
//
//	struct bridge_vlan_info {
//	  __u16 flags;  // VLAN-specific flags (e.g., untagged, PVID, etc.)
//	  __u16 vid;    // VLAN ID
//	};
//
// The deserialization process takes into account the system's native byte order (endianness).
//
// Linux Kernel Reference:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
//
// Parameters:
// - b: The byte slice that contains the serialized form of the BridgeVlanInfo object.
//
// Returns:
// - *BridgeVlanInfo: The deserialized BridgeVlanInfo object if successful.
// - error: Returns an error if the byte slice is too short to contain the structure.
//
// Example Usage:
// ```go
// data := []byte{0x01, 0x00, 0x05, 0x00} // Example serialized data
// vlanInfo, err := DeserializeBridgeVlanInfo(data)
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Printf("VLAN ID: %d, Flags: %d\n", vlanInfo.Vid, vlanInfo.Flags)
// ```
func DeserializeBridgeVlanInfo(b []byte) (*BridgeVlanInfo, error) {

	if err := checkInputSize(b, BridgeVlanInfoSize, BridgeVlanInfoSize); err != nil {
		return nil, err
	}

	// Deserialize the byte slice into a BridgeVlanInfo structure
	return &BridgeVlanInfo{
		Flags: BridgeVlanInfoFlags(NativeEndian.Uint16(b[0:2])), // Deserialize the Flags field
		Vid:   VlanIdType(NativeEndian.Uint16(b[2:4])),          // Deserialize the VLAN ID field
	}, nil
}
