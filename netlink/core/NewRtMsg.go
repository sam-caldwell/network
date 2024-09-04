package core

import (
	"golang.org/x/sys/unix"
)

// NewRtMsg - Create, configure and return a new RtLink
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
func NewRtMsg() *RtMsg {
	return &RtMsg{
		RtMsg: unix.RtMsg{
			Table:    unix.RT_TABLE_MAIN,
			Scope:    unix.RT_SCOPE_UNIVERSE,
			Protocol: unix.RTPROT_BOOT,
			Type:     unix.RTN_UNICAST,
		},
	}
}
