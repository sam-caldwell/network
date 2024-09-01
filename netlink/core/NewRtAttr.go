//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// NewRtAttr - Create a new Extended RtAttr object
func NewRtAttr(attrType int, data []byte) *RtAttr {

	return &RtAttr{

		RtAttr: unix.RtAttr{

			Type: uint16(attrType),
		},

		children: []NetlinkRequestData{},

		Data: data,
	}

}
