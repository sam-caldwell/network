package core

import (
	"golang.org/x/sys/unix"
)

// Len - return the length of the RtGenMsg struct in bytes
func (msg *RtGenMsg) Len() int {

	return rtaAlignOf(unix.SizeofRtGenmsg)

}
