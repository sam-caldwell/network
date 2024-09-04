package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Serialize - Serialize the TcActionMsg into a byte slice.
// This function uses binary encoding to safely serialize the structure into a byte slice.
//
// Reference:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
func (msg *TcActionMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	// Safely write the TcActionMsg struct into the byte buffer using binary encoding
	err := binary.Write(buf, binary.LittleEndian, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize TcActionMsg: %v", err)
	}
	return buf.Bytes(), nil
}
