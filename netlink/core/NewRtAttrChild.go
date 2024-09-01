package core

// NewRtAttrChild - add RtAttr as child to parent and return this new attribute
//
// ***Deprecated***  Use AddRtAttr() on the parent object
func NewRtAttrChild(parent *RtAttr, attrType int, data []byte) *RtAttr {
	return parent.AddRtAttr(attrType, data)
}
