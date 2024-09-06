package core

// XfrmLifetimeCur represents the current state of an IPsec Security Association's (SA) lifetime.
// It corresponds to the xfrm_lifetime_cur structure in the Linux kernel, which keeps track of the
// current byte and packet counts, as well as the time when the SA was added and last used.
//
// This struct is used in the context of monitoring the usage and lifetime of an SA to determine
// when it should be re-keyed or terminated based on the configured limits.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h

type XfrmLifetimeCur struct {
	// Bytes represents the number of bytes that have been processed by the SA.
	Bytes uint64

	// Packets represents the number of packets that have been processed by the SA.
	Packets uint64

	// AddTime is the time (in seconds) since the SA was added.
	// This helps determine the SA's age relative to the configured lifetime limits.
	AddTime uint64

	// UseTime is the time (in seconds) since the SA was last used.
	// This helps track how recently the SA was used to determine if it should be re-keyed or terminated.
	UseTime uint64
}
