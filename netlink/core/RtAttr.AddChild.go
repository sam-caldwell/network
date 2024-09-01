package core

// AddChild adds an existing NetlinkRequestData as a child.
func (a *RtAttr) AddChild(attr NetlinkRequestData) {
	a.children = append(a.children, attr)
}
