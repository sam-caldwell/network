//go:build linux

package core

import (
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeIfAddressMessage - deserialize the interface address message
func DeserializeIfAddressMessage(b []byte) (result *IfAddressMessage, err error) {
	if len(b) < SizeOfIfAddressMessage {
		return nil, errors.New("IfAddressMessage too short")
	}
	result = &IfAddressMessage{
		unix.IfAddrmsg{
			Family:    b[0],
			Prefixlen: b[1],
			Flags:     b[2],
			Scope:     b[3],
		},
	}

	NativeEndian.PutUint32(b[4:8], result.Index)

	return result, nil
}
