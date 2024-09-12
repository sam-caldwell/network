package core

import "encoding/binary"

// MPLSStackDecode - given a byte slice (inputBuffer) return a list of MplsLabels
func MPLSStackDecode(inputBuffer []byte) []MplsLabels {
	if len(inputBuffer)%4 != 0 {
		return nil
	}
	stack := make([]MplsLabels, 0, len(inputBuffer)/4)
	for len(inputBuffer) > 0 {
		l := binary.BigEndian.Uint32(inputBuffer[:4])
		inputBuffer = inputBuffer[4:]
		stack = append(stack, MplsLabels(l)>>MplsLsLabelShift)
		if (l>>MplsLsSShift)&1 > 0 {
			break
		}
	}
	return stack
}
