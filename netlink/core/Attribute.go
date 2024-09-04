package core

// Attribute - Netlink message attribute
type Attribute struct {
	Type  NlaFlags
	Value []byte
}
