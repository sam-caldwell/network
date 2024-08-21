package core

// Uint32Bitfield - represent a 32-bit field with individual bits or groups of bits used to
// encode various flags or options.
type Uint32Bitfield struct {
	Value    uint32
	Selector uint32
}
