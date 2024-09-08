package core

import (
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeRtNextHop - Converts a byte slice into an RtNexthop struct.
func DeserializeRtNextHop(b []byte) (*RtNexthop, error) {
	if len(b) < SizeOfRtNextHop {
		return nil, errors.New("input too short")
	}

	return &RtNexthop{
		RtNexthop: unix.RtNexthop{
			Len:     binary.LittleEndian.Uint16(b[:2]),         // Deserialize Len field (2 bytes)
			Flags:   b[2],                                      // Deserialize Flags (1 byte)
			Hops:    b[3],                                      // Deserialize Hops (1 byte)
			Ifindex: int32(binary.LittleEndian.Uint32(b[4:8])), // Deserialize Ifindex (4 bytes, little-endian)
		},
	}, nil
}
