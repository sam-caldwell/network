package core

import (
	"errors"
)

// DeserializeTunnelKey safely deserializes a byte slice into a TcTunnelKey struct.
// The function ensures that the byte slice is the correct length and uses
// binary encoding to parse the data correctly.
//
// type TcTunnelKey struct {
//    TcGen
//    Action int32
// }
//
// type TcGen struct {
//    Index   uint32
//    Capab   uint32
//    Action  int32
//    Refcnt  int32
//    Bindcnt int32
// }
//
func DeserializeTunnelKey(b []byte) (*TcTunnelKey, error) {
	if len(b) < SizeOfTcTunnelKey {
		return nil, errors.New("input too short")
	}

	msg := &TcTunnelKey{}

	// Deserialize each field in the struct.
	msg.Index = NativeEndian.Uint32(b[0:4])
	msg.Capab = NativeEndian.Uint32(b[4:8])
	msg.Action = int32(NativeEndian.Uint32(b[8:12]))
	msg.Refcnt = int32(NativeEndian.Uint32(b[12:16]))
	msg.Bindcnt = int32(NativeEndian.Uint32(b[16:20]))

	return msg, nil
}
