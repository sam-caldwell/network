package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeUnixRtAttr deserializes a byte slice into a unix.RtAttr structure.
func DeserializeUnixRtAttr(unreadData []byte) (*unix.RtAttr, error) {
	if len(unreadData) < SizeOfUnixRtAttr {
		return nil, errors.New("byte slice is too short to contain RtAttr")
	}

	var attr unix.RtAttr
	buf := bytes.NewReader(unreadData)

	if err := binary.Read(buf, NativeEndian, &attr); err != nil {
		return nil, err
	}

	return &attr, nil
}
