package core

import (
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeNetlinkMessageHeaderAndData - deserialize a message header and data from a byte slice.
func DeserializeNetlinkMessageHeaderAndData(b []byte) (header *unix.NlMsghdr, remainingData []byte,
	messageLength int, err error) {

	if len(b) < unix.NLMSG_HDRLEN {
		return nil, nil, 0, errors.New("input too short to contain NlMsghdr")
	}

	header, err = DeserializeNlMsgHdr(b)
	if err != nil {
		return nil, nil, 0, err
	}

	messageLength = nlmAlignOf(int(header.Len))

	if int(header.Len) < unix.NLMSG_HDRLEN || messageLength > len(b) {
		return nil, nil, 0, unix.EINVAL
	}

	return header, b[unix.NLMSG_HDRLEN:], messageLength, nil
}
