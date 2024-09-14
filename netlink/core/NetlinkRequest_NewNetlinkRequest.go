package core

import (
	"golang.org/x/sys/unix"
	"sync/atomic"
)

// NewNetlinkRequest - Create a new netlink request from proto and flags.
//
// Warning: The Len value is inaccurate until the message is serialized.
func NewNetlinkRequest(proto, flags int) *NetlinkRequest {
	return &NetlinkRequest{
		NlMsghdr: unix.NlMsghdr{
			Len:   uint32(NetlinkMessageHeaderSize),
			Type:  uint16(proto),
			Flags: unix.NLM_F_REQUEST | uint16(flags),
			Seq:   atomic.AddUint32(&nextSequenceNumber, 1),
		},
	}
}
