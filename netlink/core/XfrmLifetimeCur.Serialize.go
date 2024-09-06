package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Serialize - safely convert the XfrmLifetimeCur struct into a byte slice by serializing each field in the correct
// order using the encoding/binary package.
func (msg *XfrmLifetimeCur) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize each field in big-endian order to ensure compatibility across different architectures.

	// Serialize Bytes (the number of bytes processed by the SA).
	if err := binary.Write(buf, binary.BigEndian, msg.Bytes); err != nil {
		return nil, fmt.Errorf("failed to serialize Bytes: %w", err)
	}

	// Serialize Packets (the number of packets processed by the SA).
	if err := binary.Write(buf, binary.BigEndian, msg.Packets); err != nil {
		return nil, fmt.Errorf("failed to serialize Packets: %w", err)
	}

	// Serialize AddTime (time since the SA was added).
	if err := binary.Write(buf, binary.BigEndian, msg.AddTime); err != nil {
		return nil, fmt.Errorf("failed to serialize AddTime: %w", err)
	}

	// Serialize UseTime (time since the SA was last used).
	if err := binary.Write(buf, binary.BigEndian, msg.UseTime); err != nil {
		return nil, fmt.Errorf("failed to serialize UseTime: %w", err)
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
