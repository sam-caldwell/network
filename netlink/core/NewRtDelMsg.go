package core

import (
	"golang.org/x/sys/unix"
)

// NewRtDelMsg - Returns a new RtMsg struct for deleting routes.
//
// This function creates and returns a new RtMsg structure configured for route deletion.
// It uses default settings, where the routing table is set to `RT_TABLE_MAIN` (main routing table)
// and the scope is set to `RT_SCOPE_NOWHERE` (indicating that no destination is reachable).
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
func NewRtDelMsg() *RtMsg {
	return &RtMsg{
		RtMsg: unix.RtMsg{
			Table: unix.RT_TABLE_MAIN,    // The main routing table
			Scope: unix.RT_SCOPE_NOWHERE, // No destination is reachable
		},
	}
}
