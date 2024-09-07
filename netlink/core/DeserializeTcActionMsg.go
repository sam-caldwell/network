package core

import (
	"bytes"
	"encoding/binary"
	"log"
)

// DeserializeTcActionMsg safely deserializes a byte slice into a TcActionMsg object.
func DeserializeTcActionMsg(b []byte) *TcActionMsg {
	if len(b) < SizeOfTcActionMsg {
		log.Fatalf("Byte slice too short to deserialize TcActionMsg")
	}

	msg := &TcActionMsg{}
	buf := bytes.NewReader(b[:SizeOfTcActionMsg])
	err := binary.Read(buf, NativeEndian, msg)
	if err != nil {
		log.Fatalf("Failed to deserialize TcActionMsg: %v", err)
	}

	return msg
}
