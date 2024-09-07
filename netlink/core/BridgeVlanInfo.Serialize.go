//go:build linux

package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Serialize the BridgeVlanInfo to a byte slice
func (bridge *BridgeVlanInfo) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Convert each field to the appropriate byte representation
	// Assuming little-endian, change to big-endian if needed
	if err := binary.Write(buf, NativeEndian, bridge.Flags); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, NativeEndian, bridge.Vid); err != nil {
		return nil, err
	}

	// Return the serialized byte slice
	return buf.Bytes(), nil
}
