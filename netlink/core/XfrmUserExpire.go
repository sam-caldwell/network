package core

// XfrmUserExpire represents the xfrm_user_expire structure used in IPsec state expiration.
//
// This structure contains information about an expired Security Association (SA) state
// and whether the expiration was due to soft or hard limits being reached.
//
// The struct corresponds to the Linux kernel's netlink message for XFRM (IPsec) expiration events.
//
// Fields:
// - XfrmUsersaInfo: Contains details about the expired SA (such as encryption algorithms, SPI, etc.).
// - Hard: A flag indicating if the expiration was due to hard (1) or soft (0) limits.
// - Pad: Padding to align the structure to 64 bits.
//
// Reference:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmUserExpire struct {

	// XfrmUsersaInfo contains information about the expired SA.
	XfrmUsersaInfo XfrmUsersaInfo

	// Hard indicates if the expiration was due to hard limits (1) or soft limits (0).
	Hard uint8

	// Pad ensures 64-bit alignment of the structure.
	Pad [7]byte
}
