package core

// XfrmLifetimeCfg represents configuration limits for IPsec Security Associations (SAs) lifetime.
// It corresponds to the xfrm_lifetime_cfg structure in the Linux kernel.
//
// The struct defines the soft and hard limits for bytes, packets, and time (both for adding and using the SA).
//
// For more details, see the Linux kernel code here:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h

type XfrmLifetimeCfg struct {
	// SoftByteLimit is the soft limit for the number of bytes that can be processed by the SA before it needs to be re-keyed.
	SoftByteLimit uint64

	// HardByteLimit is the hard limit for the number of bytes that can be processed by the SA.
	// Once this limit is reached, the SA must be terminated.
	HardByteLimit uint64

	// SoftPacketLimit is the soft limit for the number of packets that can be processed by the SA before it needs to be re-keyed.
	SoftPacketLimit uint64

	// HardPacketLimit is the hard limit for the number of packets that can be processed by the SA.
	// Once this limit is reached, the SA must be terminated.
	HardPacketLimit uint64

	// SoftAddExpiresSeconds is the soft limit for the time (in seconds) since the SA was added,
	// after which it should be re-keyed.
	SoftAddExpiresSeconds uint64

	// HardAddExpiresSeconds is the hard limit for the time (in seconds) since the SA was added.
	// Once this limit is reached, the SA must be terminated.
	HardAddExpiresSeconds uint64

	// SoftUseExpiresSeconds is the soft limit for the time (in seconds) since the SA was last used,
	// after which it should be re-keyed.
	SoftUseExpiresSeconds uint64

	// HardUseExpiresSeconds is the hard limit for the time (in seconds) since the SA was last used.
	// Once this limit is reached, the SA must be terminated.
	HardUseExpiresSeconds uint64
}
