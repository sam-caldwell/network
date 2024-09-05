package core

import "errors"

// DeserializeXfrmUserExpire safely converts a byte slice into an XfrmUserExpire struct.
// Ensures that the input byte slice is of the correct length, avoiding the use of unsafe pointers.
func DeserializeXfrmUserExpire(b []byte) (*XfrmUserExpire, error) {
	// Check if the input byte slice has the correct length
	if len(b) < SizeofXfrmUserExpire {
		return nil, errors.New("DeserializeXfrmUserExpire: byte slice too short")
	}

	msg := &XfrmUserExpire{}

	// Manually copy each field from the byte slice to the struct
	copy(msg.XfrmUsersaInfo.Serialize(), b[:SizeofXfrmUsersaInfo])
	msg.Hard = b[SizeofXfrmUsersaInfo]
	// Ensure remaining bytes are appropriately set (padding)
	copy(msg.Pad[:], b[SizeofXfrmUsersaInfo+1:SizeofXfrmUserExpire])

	return msg, nil
}
