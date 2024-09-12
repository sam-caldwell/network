package core

import "encoding/binary"

// MplsStackEncode - given a list of MplsLabels, return a serialized byte slice.
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
func MplsStackEncode(labels ...uint32) []byte {

	const sBit = 1 << MplsLsSShift

	output := make([]byte, 4*len(labels)) // pre-allocate 4 bytes per uint32 element in the input.

	for idx, label := range labels {
		shiftedLabel := label << MplsLsLabelShift

		// Set the S bit (Bottom of Stack) for the last label in the third byte (8th bit of 3rd byte)
		if idx == len(labels)-1 {
			shiftedLabel |= sBit
		}

		// Set the value in the byte slice
		binary.BigEndian.PutUint32(output[idx*4:], shiftedLabel)
	}

	return output
}
