package core

// XfrmUserSpiInfo represents the `xfrm_userspi_info` structure used in the Linux kernel's XFRM subsystem.
// This structure is used to manage the Security Parameter Index (SPI) range and the associated Security Association (SA).
// The SPI is a unique identifier for an IPsec Security Association.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmUserSpiInfo struct {
	// XfrmUsersaInfo contains the details about the Security Association (SA), such as traffic selectors,
	// source address, lifetime configuration, current lifetime, and related statistics.
	XfrmUsersaInfo XfrmUsersaInfo

	// Min represents the minimum value in the SPI range. SPIs are 32-bit values, and this field indicates
	// the lower bound for the range of SPIs that may be assigned to Security Associations.
	Min uint32

	// Max represents the maximum value in the SPI range. SPIs are 32-bit values, and this field indicates
	// the upper bound for the range of SPIs that may be assigned to Security Associations.
	Max uint32
}
