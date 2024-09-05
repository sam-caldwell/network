package core

import (
	"errors"
)

// DeserializeTcSfqRedStats - Safely deserializes a byte slice into a TcSfqRedStats struct.
// This function ensures the byte slice is of the correct length and manually decodes each field.
func DeserializeTcSfqRedStats(b []byte) (*TcSfqRedStats, error) {
	if len(b) < SizeofTcSfqRedStats {
		return nil, errors.New("DeserializeTcSfqRedStats: byte slice too short")
	}

	// Deserialize each field using the appropriate endianness
	return &TcSfqRedStats{
		ProbDrop:       NativeEndian.Uint32(b[0:4]),
		ForcedDrop:     NativeEndian.Uint32(b[4:8]),
		ProbMark:       NativeEndian.Uint32(b[8:12]),
		ForcedMark:     NativeEndian.Uint32(b[12:16]),
		ProbMarkHead:   NativeEndian.Uint32(b[16:20]),
		ForcedMarkHead: NativeEndian.Uint32(b[20:24]),
	}, nil
}
