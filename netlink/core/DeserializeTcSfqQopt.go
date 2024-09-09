package core

import (
	"errors"
)

// DeserializeTcSfqQopt - Safely converts a byte slice into a TcSfqQopt structure.
// This method checks the length of the input byte slice to ensure it matches the expected size of TcSfqQopt.
// It manually reads each field using the appropriate endianness.
//
// If the input byte slice is too short, it returns an error.
func DeserializeTcSfqQopt(b []byte) (*TcSfqQopt, error) {
	if len(b) < SizeOfTcSfqQopt {
		return nil, errors.New("input too short")
	}

	msg := &TcSfqQopt{}

	// Deserialize each field manually, using the correct offsets.
	msg.Quantum = b[0]
	msg.Perturb = int32(NativeEndian.Uint32(b[1:5]))
	msg.Limit = NativeEndian.Uint32(b[5:9])
	msg.Divisor = b[9]
	msg.Flows = b[10]

	return msg, nil
}
