package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Serialize - safely convert the XfrmId struct into a byte slice
// by serializing each field using the encoding/binary package.
// This ensures portability and avoids the use of unsafe pointers.
func (msg *XfrmId) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize the Daddr field (assuming XfrmAddress is a byte array or struct).
	// Adjust for the size of Daddr if necessary.
	if err := binary.Write(buf, binary.BigEndian, msg.Daddr); err != nil {
		return nil, fmt.Errorf("failed to serialize Daddr: %w", err)
	}

	// Serialize the Spi (Security Parameter Index), which is a uint32 (big-endian).
	if err := binary.Write(buf, binary.BigEndian, msg.Spi); err != nil {
		return nil, fmt.Errorf("failed to serialize Spi: %w", err)
	}

	// Serialize the Proto field (uint8).
	if err := binary.Write(buf, binary.BigEndian, msg.Proto); err != nil {
		return nil, fmt.Errorf("failed to serialize Proto: %w", err)
	}

	// Serialize the Pad field (3-byte array).
	if err := binary.Write(buf, binary.BigEndian, msg.Pad); err != nil {
		return nil, fmt.Errorf("failed to serialize Pad: %w", err)
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
