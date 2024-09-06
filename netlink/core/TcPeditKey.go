package core

// TcPeditKey represents the key structure used in the pedit action within the Linux kernel's traffic control (TC).
// The pedit action is used to modify packet headers, and each key defines how specific portions of a packet
// should be manipulated using masks, offsets, and shift operations.
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type TcPeditKey struct {
	// Mask represents the bitmask to apply to the packet data.
	Mask uint32

	// Val represents the value to write after applying the mask.
	Val uint32

	// Off represents the offset in the packet where the modification will be applied.
	Off uint32

	// At represents the position within the offset where the value will be written.
	At uint32

	// OffMask represents a mask to apply to the offset, which allows for more precise targeting.
	OffMask uint32

	// Shift represents the number of bits to shift the value by before writing it to the packet.
	Shift uint32
}
