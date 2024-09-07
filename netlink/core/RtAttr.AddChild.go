package core

// AddChild adds an existing NetlinkRequestData as a child.
func (attr *RtAttr) AddChild(rhs NetlinkRequestData) {
	attr.children = append(attr.children, rhs)
}
