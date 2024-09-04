package core

// Len - Return the length of the RtNexthop struct in bytes
func (msg *RtNexthop) Len() int {

	if len(msg.Children) == 0 {
		return SizeofRtNexthop
	}

	l := 0

	for _, child := range msg.Children {
		l += rtaAlignOf(child.Len())
	}

	l += SizeofRtNexthop

	return rtaAlignOf(l)

}
