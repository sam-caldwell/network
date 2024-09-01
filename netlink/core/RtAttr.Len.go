package core

import (
	"golang.org/x/sys/unix"
)

// Len - return the byte length in memory of the RtAttr object and its children
func (a *RtAttr) Len() int {

	if len(a.children) == 0 {
		return unix.SizeofRtAttr + len(a.Data)
	}

	l := 0

	for _, child := range a.children {
		l += rtaAlignOf(child.Len())
	}

	l += unix.SizeofRtAttr

	return rtaAlignOf(l + len(a.Data))

}
