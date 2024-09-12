package core

import (
	"encoding/binary"
	"log"
)

// MplsStackDecode - given a byte slice (inputBuffer) return a list of MplsLabels
//
// Reference: RFC 5462, RFC 3032
//
//	  0                   1                   2                   3
//	  0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//	 +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	 |                Label                  | TC  |S|       TTL     |
//	 +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
//		Label:  Label Value, 20 bits
//		TC:     Traffic Class field, 3 bits
//		S:      Bottom of Stack, 1 bit
//		TTL:    Time to Live, 8 bits
func MplsStackDecode(inputBuffer []byte) []uint32 {

	if len(inputBuffer)%4 != 0 {
		log.Printf("return nil on %v", inputBuffer)
		return nil
	}

	sz := len(inputBuffer) / 4
	stack := make([]uint32, 0, sz)

	for len(inputBuffer) > 0 {
		l := binary.BigEndian.Uint32(inputBuffer[:4])
		inputBuffer = inputBuffer[4:]

		// Correct extraction of the MPLS label (upper 20 bits)
		label := l >> MplsLsLabelShift
		log.Printf("Label: %04x", label)

		// Append the label to the stack
		stack = append(stack, label)

		// If the bottom-of-stack bit is set, break the loop
		if (l>>8)&1 > 0 {
			break
		}
	}

	return stack
}
