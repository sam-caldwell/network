package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Converts the TcTbfQopt object into a byte slice in a safe manner.
// This function ensures safe serialization by manually converting each field
// to a byte slice using the binary encoding package.
func (msg *TcTbfQopt) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Use LittleEndian or BigEndian depending on your architecture
	err := binary.Write(buf, binary.LittleEndian, msg.Rate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, msg.Peakrate)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, msg.Limit)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, msg.Buffer)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, msg.Mtu)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
