package core

import (
	"sync/atomic"
)

// NewNetlinkRequest - Create a new netlink request from proto and flags.
//
// Warning: The Len value is inaccurate until the message is serialized.
func NewNetlinkRequest(proto, flags int) *NetlinkRequest {
	return &NetlinkRequest{
		NetlinkMessageHeader: NetlinkMessageHeader{
			Len:   uint32(NetlinkMessageHeaderSize),
			Type:  uint16(proto),
			Flags: NlmFRequest | uint16(flags),
			Seq:   atomic.AddUint32(&nextSequenceNumber, 1),
		},
	}
}
