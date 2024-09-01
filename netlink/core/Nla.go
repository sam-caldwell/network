//go:build linux

package core

const (

	// NlaFNested - NLA_F_NESTED (1 << 15)
	NlaFNested uint16 = 32768

	// NlaFNetByteorder - NLA_F_NET_BYTEORDER (1 << 14)
	NlaFNetByteorder uint16 = 16384

	// NlaTypeMask - NLA_TYPE_MASK - Mask bits for the NLA type
	NlaTypeMask = ^(NlaFNested | NlaFNetByteorder)

	// NlaAlignto - NLA_ALIGNTO 4
	NlaAlignto uint16 = 4
)
