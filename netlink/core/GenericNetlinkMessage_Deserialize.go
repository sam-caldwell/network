package core

// DeserializeGenericNetlinkMessage - creates a GenericNetlinkMessage structure from a byte slice.  It returns an
// error if the byte slice is of incorrect length.
func DeserializeGenericNetlinkMessage(b []byte) (*GenericNetlinkMessage, error) {

	if err := checkInputSize(b, GenericNetlinkMessageSize, GenericNetlinkMessageSize); err != nil {
		return nil, err
	}

	return &GenericNetlinkMessage{
		Command: b[0],
		Version: b[1],
	}, nil

}
