//go:build linux

package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Serialize the BridgeVlanInfo to a byte slice
func (bridge *BridgeVlanInfo) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Convert each field to the appropriate byte representation
	// Assuming little-endian, change to big-endian if needed
	if err := binary.Write(buf, binary.LittleEndian, bridge.Flags); err != nil {
		// Handle the error according to your needs
		return nil
	}

	if err := binary.Write(buf, binary.LittleEndian, bridge.Vid); err != nil {
		// Handle the error according to your needs
		return nil
	}

	// Return the serialized byte slice
	return buf.Bytes()
}
