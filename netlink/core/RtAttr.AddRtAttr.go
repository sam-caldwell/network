package core

// AddRtAttr adds an RtAttr as a child and returns the new attribute
func (a *RtAttr) AddRtAttr(attrType int, data []byte) *RtAttr {

	attr := NewRtAttr(attrType, data)

	a.children = append(a.children, attr)

	return attr

}
