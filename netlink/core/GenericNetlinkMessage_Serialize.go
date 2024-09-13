package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize converts the GenericNetlinkMessage structure into a byte slice.
func (msg *GenericNetlinkMessage) Serialize() ([]byte, error) {

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, NativeEndian, msg.Command); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, NativeEndian, msg.Version); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
