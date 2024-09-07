package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeNlMsgHdr deserializes a byte slice into a unix.NlMsghdr structure.
func DeserializeNlMsgHdr(unreadData []byte) (*unix.NlMsghdr, error) {
	if len(unreadData) < SizeOfNlMsgHdr {
		return nil, errors.New("byte slice is too short")
	}

	var hdr unix.NlMsghdr
	buf := bytes.NewReader(unreadData)

	if err := binary.Read(buf, NativeEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
