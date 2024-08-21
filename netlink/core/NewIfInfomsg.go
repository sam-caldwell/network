package core

import "golang.org/x/sys/unix"

// NewIfInfomsg - Create an IfInfoMsg with family specified.
func NewIfInfomsg(family IfFamily) *IfInfoMsg {
	return &IfInfoMsg{
		IfInfomsg: unix.IfInfomsg{
			Family: uint8(family),
		},
	}
}
