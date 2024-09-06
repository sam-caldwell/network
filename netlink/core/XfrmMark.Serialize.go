package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Serialize - safely convert the XfrmMark struct into a byte slice
// by serializing each field using the encoding/binary package.
// This ensures portability and avoids the use of unsafe pointers.
func (msg *XfrmMark) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize the Value field (uint32) using big-endian order.
	if err := binary.Write(buf, binary.BigEndian, msg.Value); err != nil {
		return nil, fmt.Errorf("failed to serialize Value: %w", err)
	}

	// Serialize the Mask field (uint32) using big-endian order.
	if err := binary.Write(buf, binary.BigEndian, msg.Mask); err != nil {
		return nil, fmt.Errorf("failed to serialize Mask: %w", err)
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
