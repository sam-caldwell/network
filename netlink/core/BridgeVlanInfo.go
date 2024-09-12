package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"unsafe"
)

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
	Flags BridgeVlanInfoFlags // VLAN-specific flags (e.g., untagged, PVID)
	Vid   VlanIdType          // VLAN ID (range: 1-4094)
}

// SizeOfBridgeVlanInfo - size of BridgeVlanInfo
const SizeOfBridgeVlanInfo = int(unsafe.Sizeof(BridgeVlanInfo{}))

// BridgeVlanInfoFlags -
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
type BridgeVlanInfoFlags uint16

const (
	// BridgeVlanInfoFlagsUnset - flags are unset
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeVlanInfoFlagsUnset = 0

	// BridgeVlanInfoMaster - BRIDGE_VLAN_INFO_MASTER - Operate on Bridge device as well
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeVlanInfoMaster BridgeVlanInfoFlags = 1

	// BridgeVlanInfoPvid - BRIDGE_VLAN_INFO_PVID -  VLAN is PVID, ingress untagged
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeVlanInfoPvid BridgeVlanInfoFlags = 2

	// BridgeVlanInfoUntagged - BRIDGE_VLAN_INFO_UNTAGGED - VLAN egresses untagged
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeVlanInfoUntagged BridgeVlanInfoFlags = 4

	// BridgeVlanInfoRangeBegin - BRIDGE_VLAN_INFO_RANGE_BEGIN - VLAN is start of vlan range
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeVlanInfoRangeBegin BridgeVlanInfoFlags = 8

	// BridgeVlanInfoRangeEnd - BRIDGE_VLAN_INFO_RANGE_END - VLAN is end of vlan range
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeVlanInfoRangeEnd BridgeVlanInfoFlags = 16

	// BridgeVlanInfoBrentry - BRIDGE_VLAN_INFO_BRENTRY - Global bridge VLAN entry
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeVlanInfoBrentry BridgeVlanInfoFlags = 32

	// BridgeVlanInfoOnlyOpts - BRIDGE_VLAN_INFO_ONLY_OPTS - Skip create/delete/flags
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeVlanInfoOnlyOpts BridgeVlanInfoFlags = 64
)

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

// IngressUntag checks whether the BridgeVlanInfo structure has the "untagged" flag set.
// This flag indicates whether incoming packets with a matching VLAN ID should have their
// VLAN tags removed (untagged) before being forwarded. This is particularly useful in VLAN
// configurations where you want to manage VLAN tags on ingress into a network bridge.
//
// The "untagged" flag is part of the BridgeVlanInfo's Flags field and is represented by the
// BridgeVlanInfoUntagged constant.
//
// Reference:
//   - BridgeVlanInfo structure as defined in the Linux kernel's "if_link.h":
//     https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
//
// Returns:
//   - true if the "untagged" flag is set in the Flags field, false otherwise.
func (bridge *BridgeVlanInfo) IngressUntag() bool {
	return bridge.Flags&BridgeVlanInfoUntagged > 0
}

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

// Serialize serializes the BridgeVlanInfo struct to a byte slice.
//
// It converts each field in the BridgeVlanInfo struct to its byte representation,
// using the system's native endianness (little-endian or big-endian depending on the system).
//
// Returns:
// - A byte slice containing the serialized data.
// - An error if the serialization process fails.
//
// The `binary.Write` method is used to encode the fields in the buffer.
//
// Fields:
// - `Flags`: A uint16 field representing flags for the VLAN, which is serialized first.
// - `Vid`: A uint16 field representing the VLAN ID, serialized after `Flags`.
//
// Example Usage:
// ```go
//
//	bridge := BridgeVlanInfo{
//	    Flags: 0x0100,
//	    Vid:   100,
//	}
//
// serializedData, err := bridge.Serialize()
//
//	if err != nil {
//	    log.Fatalf("serialization failed: %v", err)
//	}
//
// ```
func (bridge *BridgeVlanInfo) Serialize() ([]byte, error) {
	// Create a new buffer to hold the serialized data
	buf := new(bytes.Buffer)

	// Serialize the Flags field (uint16)
	if err := binary.Write(buf, NativeEndian, bridge.Flags); err != nil {
		return nil, err
	}

	// Serialize the Vid field (uint16)
	if err := binary.Write(buf, NativeEndian, bridge.Vid); err != nil {
		return nil, err
	}

	// Return the byte slice containing the serialized data
	return buf.Bytes(), nil
}

// String - Returns a string representation of the BridgeVlanInfo struct.
// It outputs the field names and their values in a key-value format, making it easier to debug and inspect
// the content of the BridgeVlanInfo struct.
//
// Example output:
// BridgeVlanInfo{Flags: 0x0100, Vid: 100}
//
// This method is useful when logging or printing the contents of a BridgeVlanInfo object.
func (bridge *BridgeVlanInfo) String() string {
	return fmt.Sprintf("BridgeVlanInfo{Flags: 0x%04x, Vid: %d}", bridge.Flags, bridge.Vid)
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
