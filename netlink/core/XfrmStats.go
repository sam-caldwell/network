package core

// XfrmStats represents the `xfrm_stats` structure used in the Linux kernel's XFRM subsystem.
// This structure holds various statistics for IPsec Security Associations (SAs), including replay window,
// replay errors, and integrity check failures.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmStats struct {
	// ReplayWindow tracks the size of the replay window, which is used to prevent replay attacks.
	// This field represents how many recent packets are remembered to detect duplicates.
	ReplayWindow uint32

	// Replay counts the number of replay errors detected for the Security Association (SA).
	// Replay errors occur when a packet is detected as a duplicate or out of order based on the replay window.
	Replay uint32

	// IntegrityFailed counts the number of packets that failed integrity checks.
	// This occurs when the integrity of a packet cannot be verified (e.g., due to a wrong authentication key).
	IntegrityFailed uint32
}
