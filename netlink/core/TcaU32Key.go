package core

// TcU32Key - tc_u32_key - Represents a key for the U32 filter in Traffic Control (TC).
//
// This structure is used to specify a key in the U32 classifier for matching packets
// based on a bitmask and value. It is part of the U32 filter which allows for flexible
// filtering based on packet fields.
//
// Fields:
// - Mask: The bitmask applied to the packet field to extract specific bits. This is stored in big-endian format.
// - Val: The value to match after the mask is applied. This is stored in big-endian format.
// - Off: The offset where the field should be matched within the packet.
// - OffMask: A mask applied to the offset, allowing for flexible field matching.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
type TcU32Key struct {
	// Mask - The bitmask used to extract the relevant bits from the packet.
	Mask uint32 // big endian

	// Val - The value that should match after applying the mask.
	Val uint32 // big endian

	// Off - The offset in the packet where the match is performed.
	Off int32

	// OffMask - The mask applied to the offset for flexible matching.
	OffMask int32
}
