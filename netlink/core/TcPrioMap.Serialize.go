package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize serializes the TcPrioMap struct into a byte slice.
func (x *TcPrioMap) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize the Bands field (int32).
	err := binary.Write(buf, binary.LittleEndian, x.Bands)
	if err != nil {
		return nil, err
	}

	// Serialize the Priomap field (array of uint8).
	err = binary.Write(buf, binary.LittleEndian, x.Priomap)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
