package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Serialize - safely convert the XfrmLifetimeCfg struct into a byte slice by serializing each field in the correct
// order using the encoding/binary package.
func (msg *XfrmLifetimeCfg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize each field in big-endian order.
	if err := binary.Write(buf, binary.BigEndian, msg.SoftByteLimit); err != nil {
		return nil, fmt.Errorf("failed to serialize SoftByteLimit: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, msg.HardByteLimit); err != nil {
		return nil, fmt.Errorf("failed to serialize HardByteLimit: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, msg.SoftPacketLimit); err != nil {
		return nil, fmt.Errorf("failed to serialize SoftPacketLimit: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, msg.HardPacketLimit); err != nil {
		return nil, fmt.Errorf("failed to serialize HardPacketLimit: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, msg.SoftAddExpiresSeconds); err != nil {
		return nil, fmt.Errorf("failed to serialize SoftAddExpiresSeconds: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, msg.HardAddExpiresSeconds); err != nil {
		return nil, fmt.Errorf("failed to serialize HardAddExpiresSeconds: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, msg.SoftUseExpiresSeconds); err != nil {
		return nil, fmt.Errorf("failed to serialize SoftUseExpiresSeconds: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, msg.HardUseExpiresSeconds); err != nil {
		return nil, fmt.Errorf("failed to serialize HardUseExpiresSeconds: %w", err)
	}

	// Return the serialized byte slice
	return buf.Bytes(), nil
}
