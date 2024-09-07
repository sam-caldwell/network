package core

import "errors"

// DeserializeXfrmUserExpire safely converts a byte slice into an XfrmUserExpire struct.
// Ensures that the input byte slice is of the correct length, avoiding the use of unsafe pointers.
func DeserializeXfrmUserExpire(b []byte) (*XfrmUserExpire, error) {
	// Check if the input byte slice has the correct length
	if len(b) < SizeOfXfrmUserExpire {
		return nil, errors.New("DeserializeXfrmUserExpire: byte slice too short")
	}

	// Manually copy each field from the byte slice to the struct
	info, err := DeserializeXfrmUserSaInfo(b)
	if err != nil {
		return nil, err
	}
	return &(XfrmUserExpire{
		XfrmUsersaInfo: *info,
		Hard:           b[SizeOfXfrmUserSaInfo],
		Pad:            [7]byte(b[SizeOfXfrmUserSaInfo+1:]),
	}), nil
}
