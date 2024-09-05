package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize converts the Genlmsg structure into a byte slice.
func (msg *Genlmsg) Serialize() ([]byte, error) {

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, NativeEndian, msg.Command); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, NativeEndian, msg.Version); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
