package core

// NewIfInfomsgChild - creates and returns a new IfInfomsgChild instance.
func NewIfInfomsgChild(parent *RtAttr, family InterfaceFamily) *IfInfoMsg {
	msg := NewIfInfomsg(family)
	parent.children = append(parent.children, msg)
	return msg
}
