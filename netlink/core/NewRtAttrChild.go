package core

// NewRtAttrChild - add RtAttr as child to parent and return this new attribute
//
// ***Deprecated***  Use AddRtAttr() on the parent object
func NewRtAttrChild(parent *RtAttr, attrType uint16, data []byte) *RtAttr {
	return parent.AddRtAttr(attrType, data)
}
