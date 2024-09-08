package core

import (
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeIfAddressMessage - deserialize the interface address message
func DeserializeIfAddressMessage(b []byte) (result *IfAddressMessage, err error) {
	if len(b) < SizeOfIfAddressMessage {
		return nil, errors.New("input too short")
	}
	result = &IfAddressMessage{
		unix.IfAddrmsg{
			Family:    b[0],
			Prefixlen: b[1],
			Flags:     b[2],
			Scope:     b[3],
		},
	}

	// Read the 4 bytes for the Index field (offset 4 to 8) and assign it
	result.Index = NativeEndian.Uint32(b[4:8])

	return result, err
}
