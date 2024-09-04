package core

import (
	"golang.org/x/sys/unix"
)

// NewRtGenMsg - Create, configure and return a pointer to a new RtGenMsg struct
func NewRtGenMsg() *RtGenMsg {
	return &RtGenMsg{
		RtGenmsg: unix.RtGenmsg{
			Family: unix.AF_UNSPEC,
		},
	}
}
