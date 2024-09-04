package core

import (
	"golang.org/x/sys/unix"
)

// DeserializeRtGenMsg - given a slice of bytes, deserialize the input into an RtGenmsg struct
func DeserializeRtGenMsg(b []byte) *RtGenMsg {

	return &RtGenMsg{
		RtGenmsg: unix.RtGenmsg{
			Family: b[0],
		},
	}

}
