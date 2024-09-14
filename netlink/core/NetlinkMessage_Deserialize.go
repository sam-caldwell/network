package core

import (
	"golang.org/x/sys/unix"
)

// DeserializeNetlinkMessage - deserialize a message header and data from a byte slice.
//
//	type NlMsghdr struct {
//	   Len   uint32
//	   Type  uint16
//	   Flags uint16
//	   Seq   uint32
//	   Pid   uint32
//	}
func DeserializeNetlinkMessage(b []byte) (header *unix.NlMsghdr, remainingData []byte, messageLength int, err error) {

	if err = checkInputSize(b, unix.NLMSG_HDRLEN, disableSizeCheck); err != nil {

		return nil, nil, 0, err

	}

	header, err = DeserializeNlMsgHdr(b)
	if err != nil {

		return nil, nil, 0, err

	}

	messageLength = nlmAlignOf(int(header.Len))

	if int(header.Len) < unix.NLMSG_HDRLEN || messageLength > len(b) {
		return nil, nil, 0, unix.EINVAL
	}

	return header, b[unix.NLMSG_HDRLEN:], messageLength, err
}
