package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeTcPrioMap safely deserializes a byte slice into a TcPrioMap structure.
// The function assumes that the input byte slice is at least as long as the size of the TcPrioMap.
//
// If the byte slice is shorter than expected, an error is returned.
//
// TcPrioMap is used to map priorities to traffic classes in the Linux traffic control system.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
func DeserializeTcPrioMap(b []byte) (*TcPrioMap, error) {
	if len(b) < SizeofTcPrioMap {
		return nil, fmt.Errorf("DeserializeTcPrioMap: input byte slice too short (%d bytes)", len(b))
	}

	msg := &TcPrioMap{}
	buffer := bytes.NewReader(b)

	err := binary.Read(buffer, NativeEndian, msg)
	if err != nil {
		return nil, fmt.Errorf("DeserializeTcPrioMap: failed to read TcPrioMap: %v", err)
	}

	return msg, nil
}
