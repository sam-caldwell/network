package core

// TcU32Sel - Represents the `tc_u32_sel` structure used in Linux traffic control (tc) for u32 classification.
//
// This structure holds selection criteria for the `u32` classifier, which is a generic, programmable packet
// classifier in Linux used to match on specific fields of a packet (e.g., IP header fields).
//
// The `tc_u32_sel` structure contains details about how to perform the classification, including flags, offsets,
// keys, and masks. These are used to control how specific portions of the packet are matched.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
type TcU32Sel struct {
	// Flags - Flags to control the behavior of the u32 classifier.
	Flags uint8

	// Offshift - Shift value for offsets.
	Offshift uint8

	// Nkeys - Number of keys used for matching.
	Nkeys uint8

	// Pad - Padding for alignment.
	Pad uint8

	// Offmask - Mask for applying offsets, stored in big endian format.
	Offmask uint16 // big endian

	// Off - Offset value to be used in matching.
	Off uint16

	// Offoff - Further offset to be applied.
	Offoff int16

	// Hoff - Hardware offset value for additional matching operations.
	Hoff int16

	// Hmask - Hardware mask for packet classification, stored in big endian format.
	Hmask uint32 // big endian

	// Keys - Array of `TcU32Key` structures used to store individual match keys.
	// Each key represents a specific match rule for the classifier.
	Keys []TcU32Key
}
