package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcNetemQopt - Deserialize a byte slice into a TcNetemQopt structure using the binary package.
func DeserializeTcNetemQopt(b []byte) (*TcNetemQopt, error) {
	if len(b) < SizeOfTcNetemQopt {
		return nil, errors.New("DeserializeTcNetemQopt: buffer too small")
	}

	buf := bytes.NewReader(b)
	msg := &TcNetemQopt{}
	err := binary.Read(buf, binary.LittleEndian, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
