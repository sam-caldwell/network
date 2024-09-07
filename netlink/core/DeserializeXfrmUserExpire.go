package core

import "errors"

// DeserializeXfrmUserExpire safely converts a byte slice into an XfrmUserExpire struct.
// Ensures that the input byte slice is of the correct length, avoiding the use of unsafe pointers.
func DeserializeXfrmUserExpire(b []byte) (*XfrmUserExpire, error) {
	// Check if the input byte slice has the correct length
	if len(b) < SizeofXfrmUserExpire {
		return nil, errors.New("DeserializeXfrmUserExpire: byte slice too short")
	}

	// Manually copy each field from the byte slice to the struct
	info, err := DeserializeXfrmUsersaInfo(b)
	if err != nil {
		return nil, err
	}
	return &(XfrmUserExpire{
		XfrmUsersaInfo: *info,
		Hard:           b[SizeofXfrmUsersaInfo],
		Pad:            [7]byte(b[SizeofXfrmUsersaInfo+1:]),
	}), nil
}
