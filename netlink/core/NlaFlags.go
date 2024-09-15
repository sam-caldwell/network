package core

import "unsafe"

type NlaFlags uint16

const (

	// NlaFNested - NLA_F_NESTED (1 << 15) 0x8000
	NlaFNested NlaFlags = 0b1000000000000000

	// NlaFNetByteorder - NLA_F_NET_BYTEORDER (1 << 14) 0x4000
	NlaFNetByteorder NlaFlags = 0b0100000000000000

	// NlaTypeMask - NLA_TYPE_MASK - Mask bits for the NLA type (0x3FFF)
	//
	//    a   0b1000000000000000 = 0x8000
	//    b   0b0100000000000000 = 0x4000
	//  -------------------------
	//   a|b  0b1100000000000000 = 0xC000
	// ^(a|b) 0b0011111111111111 = 0x3FFF
	NlaTypeMask = ^(NlaFNested | NlaFNetByteorder)

	// NlaAlignto - NLA_ALIGNTO 4
	//
	//   sizeOf(NlaTypeMask) = 2
	// x unused 32bit digits = 2
	// -------------------------
	//            NlaAlignto = 4
	NlaAlignto = NlaFlags(unsafe.Sizeof(NlaTypeMask) * 2)
)
