package core

import (
	"errors"
	"golang.org/x/sys/unix"
)

// NetlinkMessage - Full netlink message (with header).
//
// This had to be created in this project because it does not exist in the unix package yet,
// and we are trying to avoid use of syscall which is being deprecated.
type NetlinkMessage struct {
	Header unix.NlMsghdr
	Data   []byte
}

// DeserializeNetlinkMessage - deserialize a message header and data from a byte slice.
func DeserializeNetlinkMessage(b []byte) (header *unix.NlMsghdr, remainingData []byte,
	messageLength int, err error) {

	if len(b) < unix.NLMSG_HDRLEN {
		return nil, nil, 0, errors.New("input too short")
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
