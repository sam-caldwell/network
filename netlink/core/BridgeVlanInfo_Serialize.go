package core

import (
	"bytes"
	"encoding/binary"
)

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
