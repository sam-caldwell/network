package core

import (
	"errors"
)

// Deserialize - Safely converts a byte slice into a TcSfqQopt struct.
// This method ensures that the length of the input byte slice matches the expected size of TcSfqQopt.
//
// If the length of the input byte slice is less than the size of TcSfqQopt, it returns an error.
func (msg *TcSfqQopt) Deserialize(b []byte) error {
	if len(b) < SizeofTcSfqQopt {
		return errors.New("Deserialize: byte slice too short")
	}

	// Copy the content of the byte slice into the TcSfqQopt struct.
	msg.Quantum = b[0] // Quantum is a single byte
	msg.Perturb = int32(NativeEndian.Uint32(b[1:5]))
	msg.Limit = NativeEndian.Uint32(b[5:9])
	msg.Divisor = b[9] // Divisor is a single byte
	msg.Flows = b[10]  // Flows is a single byte

	return nil
}
