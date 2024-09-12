package core

import "encoding/binary"

// MplsStackEncode - given a list of MplsLabels, return a serialized byte slice.
func MplsStackEncode(labels ...MplsLabels) []byte {
	b := make([]byte, 4*len(labels))
	for idx, label := range labels {
		l := label << MplsLsLabelShift
		if idx == len(labels)-1 {
			l |= 1 << MplsLsSShift
		}
		binary.BigEndian.PutUint32(b[idx*4:], uint32(l))
	}
	return b
}
