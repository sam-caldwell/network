package core

import (
	"errors"
)

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

	// Ensure that the byte slice is at least the size of the BridgeVlanInfo structure
	if len(b) < SizeOfBridgeVlanInfo {
		return nil, errors.New("input byte slice is too short")
	}

	// Deserialize the byte slice into a BridgeVlanInfo structure
	return &BridgeVlanInfo{
		Flags: BridgeVlanInfoFlags(NativeEndian.Uint16(b[0:2])), // Deserialize the Flags field
		Vid:   VlanIdType(NativeEndian.Uint16(b[2:4])),          // Deserialize the VLAN ID field
	}, nil
}
