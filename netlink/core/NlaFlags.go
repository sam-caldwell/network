package core

import "unsafe"

type NlaFlags uint16

const (

	// NlaFNested - NLA_F_NESTED (1 << 15)
	NlaFNested NlaFlags = 32768

	// NlaFNetByteorder - NLA_F_NET_BYTEORDER (1 << 14)
	NlaFNetByteorder NlaFlags = 16384

	// NlaTypeMask - NLA_TYPE_MASK - Mask bits for the NLA type
	NlaTypeMask = ^(NlaFNested | NlaFNetByteorder)

	// NlaAlignto - NLA_ALIGNTO 4
	NlaAlignto NlaFlags = NlaFlags(unsafe.Sizeof(NlaTypeMask) * 2)
)
