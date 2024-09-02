package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

// deserializeUnixRtAttr deserializes a byte slice into a unix.RtAttr structure.
func deserializeUnixRtAttr(unreadData []byte) (*unix.RtAttr, error) {
	if len(unreadData) < unix.SizeofRtAttr {
		return nil, errors.New("byte slice is too short to contain RtAttr")
	}

	var attr unix.RtAttr
	buf := bytes.NewReader(unreadData)

	if err := binary.Read(buf, binary.LittleEndian, &attr); err != nil {
		return nil, err
	}

	return &attr, nil
}
