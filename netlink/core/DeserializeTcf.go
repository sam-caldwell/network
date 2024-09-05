package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeTcf - Safely deserialize a byte slice into a Tcf struct.
func DeserializeTcf(b []byte) (*Tcf, error) {
	if len(b) < SizeofTcf {
		return nil, fmt.Errorf("byte slice is too short to be a valid Tcf")
	}

	buf := bytes.NewReader(b)
	tcf := &Tcf{}
	err := binary.Read(buf, NativeEndian, tcf)
	if err != nil {
		return nil, err
	}

	return tcf, nil
}
