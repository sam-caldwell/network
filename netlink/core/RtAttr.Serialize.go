package core

// Serialize - serialize the RtAttr into a byte array and iterate through its children.
func (attr *RtAttr) Serialize() ([]byte, error) {

	length := attr.Len()
	buf := make([]byte, rtaAlignOf(length))
	next := 4

	if attr.Data != nil {
		copy(buf[next:], attr.Data)
		next += rtaAlignOf(len(attr.Data))
	}

	if len(attr.children) > 0 {
		for _, child := range attr.children {
			childBuf, err := child.Serialize()
			if err != nil {
				return nil, err
			}
			copy(buf[next:], childBuf)
			next += rtaAlignOf(len(childBuf))
		}
	}

	if l := uint16(length); l != 0 {
		NativeEndian.PutUint16(buf[0:2], l)
	}

	NativeEndian.PutUint16(buf[2:4], attr.Type)
	return buf, nil

}
