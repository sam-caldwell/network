package core

// DeserializeNetlinkMessage - deserialize a message header and data from a byte slice.
func DeserializeNetlinkMessage(b []byte) (header *NlMsghdr, remainingData []byte, messageLength int, err error) {

	if err = checkInputSize(b, NetlinkMessageHeaderLength, disableSizeCheck); err != nil {
		return nil, nil, 0, err
	}

	if header, err = DeserializeNetlinkMessageHeader(b); err != nil {
		return nil, nil, 0, err
	}

	messageLength = nlmAlignOf(int(header.Len))

	if int(header.Len) < NetlinkMessageHeaderLength || len(b)%4 != 0 {
		return nil, nil, 0, EINVAL
	}

	return header, b[NetlinkMessageHeaderLength:], messageLength, err
}
