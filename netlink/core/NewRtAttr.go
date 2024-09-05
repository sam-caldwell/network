//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// NewRtAttr - Create a new Extended RtAttr object
func NewRtAttr(attrType uint16, data []byte) *RtAttr {

	return &RtAttr{

		RtAttr: unix.RtAttr{

			Type: attrType,
		},

		children: []NetlinkRequestData{},

		Data: data,
	}

}
