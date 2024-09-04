package core

import "encoding/binary"

// Uint32 - return the uint32 value respecting the NET_BYTEORDER flag
func (attr *Attribute) Uint32() uint32 {

	if attr.Type&NlaFNetByteorder != 0 {

		return binary.BigEndian.Uint32(attr.Value)

	} else {

		return NativeEndian.Uint32(attr.Value)

	}

}
