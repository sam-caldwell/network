package core

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
)

// DeserializeRtNextHop - Converts a byte slice into an RtNexthop struct.
func DeserializeRtNextHop(b []byte) *RtNexthop {
	// Debugging: Print length of the byte slice
	fmt.Printf("Length of byte slice: %d, Expected size: %d\n", len(b), SizeOfRtNextHop)

	if len(b) < SizeOfRtNextHop {
		fmt.Println("Byte slice too short!")
		return nil // Error handling for insufficient byte length
	}

	// Debugging: Print the slice content for validation
	fmt.Printf("Byte slice content: %v\n", b)

	return &RtNexthop{
		RtNexthop: unix.RtNexthop{
			Len:     binary.LittleEndian.Uint16(b[:2]),         // Deserialize Len field (2 bytes)
			Flags:   b[2],                                      // Deserialize Flags (1 byte)
			Hops:    b[3],                                      // Deserialize Hops (1 byte)
			Ifindex: int32(binary.LittleEndian.Uint32(b[4:8])), // Deserialize Ifindex (4 bytes, little-endian)
		},
	}
}
