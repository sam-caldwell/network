package core

// Serialize - serialize the RtAttr into a byte array and iterate through its children.
func (a *RtAttr) Serialize() []byte {

	length := a.Len()
	buf := make([]byte, rtaAlignOf(length))
	next := 4

	if a.Data != nil {
		copy(buf[next:], a.Data)
		next += rtaAlignOf(len(a.Data))
	}

	if len(a.children) > 0 {
		for _, child := range a.children {
			childBuf := child.Serialize()
			copy(buf[next:], childBuf)
			next += rtaAlignOf(len(childBuf))
		}
	}

	if l := uint16(length); l != 0 {
		nativeEndian.PutUint16(buf[0:2], l)
	}

	nativeEndian.PutUint16(buf[2:4], a.Type)
	return buf

}
