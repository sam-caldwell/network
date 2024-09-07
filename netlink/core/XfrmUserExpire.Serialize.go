package core

// Serialize safely converts the XfrmUserExpire structure into a byte slice.
// This method manually encodes each field into the byte slice, ensuring no unsafe pointers are used.
func (msg *XfrmUserExpire) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfXfrmUserExpire)

	// Serialize the XfrmUserSaInfo structure
	xfrmInfoBytes, err := msg.XfrmUsersaInfo.Serialize()
	if err != nil {
		return nil, err
	}
	copy(buf[:SizeOfXfrmUserSaInfo], xfrmInfoBytes)

	// Serialize the Hard field
	buf[SizeOfXfrmUserSaInfo] = msg.Hard

	// Serialize the Pad field (7 bytes)
	copy(buf[SizeOfXfrmUserSaInfo+1:], msg.Pad[:])

	return buf, nil
}
