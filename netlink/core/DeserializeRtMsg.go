package core

import (
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeRtMsg - Converts a byte slice into an RtMsg struct.
//
// This function manually parses the byte slice to safely construct
// the RtMsg struct, avoiding unsafe pointer usage.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
func DeserializeRtMsg(b []byte) (*RtMsg, error) {
	if len(b) < SizeOfUnixRtMsg {
		return nil, errors.New("input too short")
	}

	return &RtMsg{
		RtMsg: unix.RtMsg{
			Family:   b[0],
			Dst_len:  b[1],
			Src_len:  b[2],
			Tos:      b[3],
			Table:    b[4],
			Protocol: b[5],
			Scope:    b[6],
			Type:     b[7],
			// Use little-endian to correctly deserialize the Flags
			Flags: binary.LittleEndian.Uint32(b[8:12]),
		},
	}, nil
}
