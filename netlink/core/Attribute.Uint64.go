package core

import "encoding/binary"

// Uint64 - return the uint64 value respecting the NET_BYTEORDER flag
func (attr *Attribute) Uint64() uint64 {

	if attr.Type&NlaFNetByteorder != 0 {

		return binary.BigEndian.Uint64(attr.Value)

	} else {

		return nativeEndian.Uint64(attr.Value)

	}

}
