package core

import (
	"golang.org/x/sys/unix"
)

// DeserializeRtMsg - Converts a byte slice into an RtMsg struct.
//
// This function manually parses the byte slice to safely construct
// the RtMsg struct, avoiding unsafe pointer usage.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
func DeserializeRtMsg(b []byte) *RtMsg {
	if len(b) < SizeOfUnixRtMsg {
		return nil // Error handling for insufficient byte length
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
			Flags:    uint32(b[8]) | uint32(b[9])<<8 | uint32(b[10])<<16 | uint32(b[11])<<24,
		},
	}
}
