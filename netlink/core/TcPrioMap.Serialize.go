package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize serializes the TcPriorityMap struct into a byte slice.
func (x *TcPriorityMap) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize the Bands field (int32).
	err := binary.Write(buf, NativeEndian, x.Bands)
	if err != nil {
		return nil, err
	}

	// Serialize the Priomap field (array of uint8).
	err = binary.Write(buf, NativeEndian, x.Priomap)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
