package core

import (
	"errors"
)

// DeserializeIfAddressMessage - deserialize the interface address message
func DeserializeIfAddressMessage(b []byte) (result *IfAddressMessage, err error) {
	if len(b) < 4 {
		return nil, errors.New("message too short")
	}
	result = &IfAddressMessage{}

	b = []byte{
		result.Family,
		result.Prefixlen,
		result.Flags,
		result.Scope,
	}

	native := NativeEndian()
	native.PutUint32(b[4:8], result.Index)

	return result, nil
}
