package core

import (
	"fmt"
)

// DeserializeTcGen - converts a byte slice into a TcGen object using safe binary decoding.
func DeserializeTcGen(b []byte) (*TcGen, error) {
	if len(b) < SizeofTcGen {
		return nil, fmt.Errorf("invalid byte slice size for TcGen")
	}

	// Create a new TcGen object
	msg := &TcGen{}

	// Deserialize fields manually using binary encoding
	msg.Index = NativeEndian.Uint32(b[0:4])
	msg.Capab = NativeEndian.Uint32(b[4:8])
	msg.Action = int32(NativeEndian.Uint32(b[8:12]))
	msg.Refcnt = int32(NativeEndian.Uint32(b[12:16]))
	msg.Bindcnt = int32(NativeEndian.Uint32(b[16:20]))

	return msg, nil
}
