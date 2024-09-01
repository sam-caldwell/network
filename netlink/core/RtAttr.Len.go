package core

import (
	"golang.org/x/sys/unix"
)

// Len - return the byte length in memory of the RtAttr object and its children
func (attr *RtAttr) Len() int {

	if len(attr.children) == 0 {
		return unix.SizeofRtAttr + len(attr.Data)
	}

	l := 0

	for _, child := range attr.children {
		l += rtaAlignOf(child.Len())
	}

	l += unix.SizeofRtAttr

	return rtaAlignOf(l + len(attr.Data))

}
