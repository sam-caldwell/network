package core

import (
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeRtGenMsg - given a slice of bytes, deserialize the input into an RtGenmsg struct
func DeserializeRtGenMsg(b []byte) (*RtGenMsg, error) {
	if b == nil {
		return nil, errors.New("nil input")
	}
	return &RtGenMsg{
		RtGenmsg: unix.RtGenmsg{
			Family: b[0],
		},
	}, nil

}
