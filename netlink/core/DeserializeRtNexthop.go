package core

import (
	"golang.org/x/sys/unix"
)

// DeserializeRtNexthop - Converts a byte slice into an RtNexthop struct.
func DeserializeRtNexthop(b []byte) *RtNexthop {
	if len(b) < unix.SizeofRtNexthop {
		return nil // Error handling for insufficient byte length
	}
	return &RtNexthop{
		RtNexthop: unix.RtNexthop{
			Len:     NativeEndian.Uint16(b[:2]),
			Flags:   b[2],
			Hops:    b[3],
			Ifindex: int32(NativeEndian.Uint32(b[4:8])),
		},
	}
}
