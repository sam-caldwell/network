package core

import (
	"golang.org/x/sys/unix"
)

// Serialize - Convert the RtMsg struct into a byte slice.
//
// This method serializes the RtMsg structure to its byte representation,
// making it ready for netlink message transmission.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
func (msg *RtMsg) Serialize() []byte {
	buf := make([]byte, unix.SizeofRtMsg)
	buf[0] = msg.Family
	buf[1] = msg.Dst_len
	buf[2] = msg.Src_len
	buf[3] = msg.Tos
	buf[4] = msg.Table
	buf[5] = msg.Protocol
	buf[6] = msg.Scope
	buf[7] = msg.Type

	// Break down Flags (uint32) into bytes
	buf[8] = byte(msg.Flags)
	buf[9] = byte(msg.Flags >> 8)
	buf[10] = byte(msg.Flags >> 16)
	buf[11] = byte(msg.Flags >> 24)

	return buf
}
