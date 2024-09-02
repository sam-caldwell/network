package core

import (
	"errors"
)

// DeserializeGenlmsg creates a Genlmsg structure from a byte slice.
// It returns an error if the byte slice is of incorrect length.
func DeserializeGenlmsg(b []byte) (*Genlmsg, error) {
	if len(b) != 2 {
		return nil, errors.New("byte slice size is incorrect for valid Genlmsg")
	}
	return &Genlmsg{
		Command: b[0],
		Version: b[1],
	}, nil
}
