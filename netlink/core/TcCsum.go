package core

// TcCsum - Represents the Checksum action in Linux Traffic Control (tc).
//
// The TcCsum struct is used for handling checksum updates in packets as part of Linux traffic control actions.
// The `UpdateFlags` field specifies which checksums should be recalculated.
//
// Fields:
// - `TcGen`: A generic structure for traffic control actions, containing index, capabilities, and other fields.
// - `UpdateFlags`: A bitmask specifying which checksum(s) to update.
//
// For more details, refer to:
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
// - Traffic Control (tc) documentation: https://man7.org/linux/man-pages/man8/tc.8.html
type TcCsum struct {

	// General traffic control action parameters
	TcGen

	// UpdateFlags - Bitmask specifying which checksum(s) to update
	UpdateFlags uint32
}
