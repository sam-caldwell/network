package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

// deserializeNetlinkMessageHeaderAndData - deserialize a message header and data from a byte slice.
func deserializeNetlinkMessageHeaderAndData(b []byte) (header *unix.NlMsghdr, remainingData []byte,
	messageLength int, err error) {

	if len(b) < unix.NLMSG_HDRLEN {
		return nil, nil, 0, errors.New("input too short to contain NlMsghdr")
	}

	buf := bytes.NewReader(b)
	header = &(unix.NlMsghdr{})

	if err = binary.Read(buf, binary.LittleEndian, header); err != nil {
		return nil, nil, 0, err
	}

	// Calculate the aligned length of the message and validate the length of the message header and data
	if messageLength = nlmAlignOf(int(header.Len)); int(header.Len) < unix.NLMSG_HDRLEN || messageLength > len(b) {
		return nil, nil, 0, unix.EINVAL
	}

	// Return the deserialized header, the remaining data, the length of the message, and nil error
	return header, b[unix.NLMSG_HDRLEN:], messageLength, nil
}
