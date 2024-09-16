package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Serialize the IfInfoMsg structure into a byte slice in a safe way.
func (msg IfInfoMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Write each field to the buffer in the correct order and format
	if err := binary.Write(buf, NativeEndian, msg.Family); err != nil {
		return nil, err
	}
	// Placeholder for the second byte, typically padding
	if err := binary.Write(buf, NativeEndian, uint8(0)); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Type); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Flags); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Change); err != nil {
		return nil, err
	}
	// Return the byte slice
	return buf.Bytes(), nil
}
