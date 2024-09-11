//go:build linux

package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
	"unsafe"
)

// IfAddressMessage - wrapper around unix.IfAddrmsg in case golang will
// eventually want to make changes here, and I need to minimize impact.
//
// // unix.IfAddrmsg looks like this...
//
//	type IfAddrmsg struct {
//		Family    uint8
//		Prefixlen uint8
//		Flags     uint8
//		Scope     uint8
//		Index     uint32
//	}
//
// Note:
//
//	We have some stricter types in this package (e.g. InterfaceFamily) to help
//	developers quickly identify relevant constants without having to resort
//	to Linux man pages.
type IfAddressMessage struct {
	unix.IfAddrmsg
}

const (
	// SizeOfIfAddressMessage     = 0x8 // bytes as derived from unix.SizeOfIfAddressMessage
	SizeOfIfAddressMessage = int(unsafe.Sizeof(IfAddressMessage{}))
)

// Len - Return the size of the interface address message
func (msg *IfAddressMessage) Len() int {
	return SizeOfIfAddressMessage
}

// Serialize - serialize an interface address message as a byte slice.
func (msg *IfAddressMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Write each field to the buffer in the correct order and format
	if err := binary.Write(buf, NativeEndian, msg.Family); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Prefixlen); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Flags); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Scope); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Index); err != nil {
		return nil, err
	}

	// Return the byte slice
	return buf.Bytes(), nil
}

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
	if len(b) < SizeOfIfAddressMessage {
		return nil, errors.New(ErrInputTooShort)
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
