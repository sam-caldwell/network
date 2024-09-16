package core

// ConnectorMessageOperation - represents connector message operation, combining message identification,
// sequence, acknowledgment, and operational flags with an additional operation-specific field.
//
// Note that this EXTENDS cn_msg found in the Linux source tree and adds an Operation field.
//
// https://github.com/torvalds/linux/blob/master/include/linux/connector.h,
// https://github.com/vishvananda/netlink/blob/main/nl/nl_linux.go
type ConnectorMessageOperation struct {
	ConnectorMessage
	// here we differ from the C header
	Operation uint32
}
