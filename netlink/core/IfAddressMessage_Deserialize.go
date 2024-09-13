package core

import (
	"golang.org/x/sys/unix"
)

// Deserialize - deserialize the interface address message
func (msg *IfAddressMessage) Deserialize(b []byte) (err error) {
	var result *IfAddressMessage
	if result, err = DeserializeIfAddressMessage(b); err != nil {
		return err
	}
	*msg = *result
	return nil
}

// DeserializeIfAddressMessage - deserialize the interface address message
func DeserializeIfAddressMessage(b []byte) (result *IfAddressMessage, err error) {
	if err := checkInputSize(b, IfAddressMessageSize, IfAddressMessageSize); err != nil {
		return nil, err
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
