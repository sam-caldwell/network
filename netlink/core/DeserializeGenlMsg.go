package core

import (
	"errors"
)

// DeserializeGenlMsg - creates a Genlmsg structure from a byte slice.  It returns an error if the byte slice is
// of incorrect length.
func DeserializeGenlMsg(b []byte) (*Genlmsg, error) {
	if len(b) < 2 {
		return nil, errors.New("byte slice size is too small")
	}
	return &Genlmsg{
		Command: b[0],
		Version: b[1],
	}, nil
}
