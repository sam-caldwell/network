package core

// XfrmMsg -  XFRM messages used in IPsec configuration and management.
// These message types are part of the Linux kernel's XFRM subsystem, which is responsible for managing
// Security Associations (SA), Security Policies (SP), and related components in IPsec.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmMsg interface {
	Type() XfrmMsgType
}
