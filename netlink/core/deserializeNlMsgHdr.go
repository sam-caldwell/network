package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

// deserializeNlMsgHdr deserializes a byte slice into a unix.NlMsghdr structure.
func deserializeNlMsgHdr(unreadData []byte) (*unix.NlMsghdr, error) {
	if len(unreadData) < unix.SizeofNlMsghdr {
		return nil, errors.New("byte slice is too short to contain NlMsghdr")
	}

	var hdr unix.NlMsghdr
	buf := bytes.NewReader(unreadData)

	if err := binary.Read(buf, NativeEndian, &hdr); err != nil {
		return nil, err
	}

	return &hdr, nil
}
