package core

// AddChild adds an existing NetlinkRequestData as a child.
func (attr *RtAttr) AddChild(attr NetlinkRequestData) {
	attr.children = append(attr.children, attr)
}
