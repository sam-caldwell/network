package core

import (
	"golang.org/x/sys/unix"
)

// NewIfAddressMessage - create, configure and return a new IFAddressMessage object (by reference)
func NewIfAddressMessage(family IfFamily) *IfAddressMessage {
	return &IfAddressMessage{
		IfAddrmsg: unix.IfAddrmsg{
			Family: uint8(family),
		},
	}
}
