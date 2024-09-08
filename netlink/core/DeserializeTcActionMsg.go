package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
)

// DeserializeTcActionMsg safely deserializes a byte slice into a TcActionMsg object.
func DeserializeTcActionMsg(b []byte) (*TcActionMsg, error) {
	if len(b) < SizeOfTcActionMsg {
		return nil, errors.New("input too short")
	}

	msg := &TcActionMsg{}
	buf := bytes.NewReader(b[:SizeOfTcActionMsg])
	err := binary.Read(buf, NativeEndian, msg)
	if err != nil {
		log.Fatalf("Failed to deserialize TcActionMsg: %v", err)
	}

	return msg, nil
}
