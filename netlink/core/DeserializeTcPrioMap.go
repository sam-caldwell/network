package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeTcPrioMap safely deserializes a byte slice into a TcPriorityMap structure.
// The function assumes that the input byte slice is at least as long as the size of the TcPriorityMap.
//
// If the byte slice is shorter than expected, an error is returned.
//
// TcPriorityMap is used to map priorities to traffic classes in the Linux traffic control system.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
func DeserializeTcPrioMap(b []byte) (*TcPriorityMap, error) {
	if len(b) < SizeOfTcPriorityMap {
		return nil, fmt.Errorf("DeserializeTcPrioMap: input byte slice too short (%d bytes)", len(b))
	}

	msg := &TcPriorityMap{}
	buffer := bytes.NewReader(b)

	err := binary.Read(buffer, NativeEndian, msg)
	if err != nil {
		return nil, fmt.Errorf("DeserializeTcPrioMap: failed to read TcPriorityMap: %v", err)
	}

	return msg, nil
}
