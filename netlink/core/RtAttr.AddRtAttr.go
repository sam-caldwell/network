package core

// AddRtAttr adds an RtAttr as a child and returns the new attribute
func (attr *RtAttr) AddRtAttr(attrType int, data []byte) *RtAttr {

	attr := NewRtAttr(attrType, data)

	attr.children = append(attr.children, attr)

	return attr

}
