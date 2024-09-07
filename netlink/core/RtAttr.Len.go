package core

// Len - return the byte length in memory of the RtAttr object and its children
func (attr *RtAttr) Len() int {

	if len(attr.children) == 0 {
		return SizeOfUnixRtAttr + len(attr.Data)
	}

	l := 0

	for _, child := range attr.children {
		l += rtaAlignOf(child.Len())
	}

	l += SizeOfUnixRtAttr

	return rtaAlignOf(l + len(attr.Data))

}
