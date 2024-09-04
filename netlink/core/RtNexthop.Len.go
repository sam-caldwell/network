package core

import (
	"golang.org/x/sys/unix"
)

func (msg *RtNexthop) Len() int {

	if len(msg.Children) == 0 {
		return unix.SizeofRtNexthop
	}

	l := 0

	for _, child := range msg.Children {
		l += rtaAlignOf(child.Len())
	}

	l += unix.SizeofRtNexthop

	return rtaAlignOf(l)

}
