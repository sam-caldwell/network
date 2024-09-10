package core

import (
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeIfInfoMsg - Deserialize byte slice into IfInfoMsg
//
//	type IfInfoMsg struct {
//	   Family uint8
//	   _      uint8
//	   Type   uint16
//	   Index  int32
//	   Flags  uint32
//	   Change uint32
//	}
func DeserializeIfInfoMsg(b []byte) (*IfInfoMsg, error) {
	if len(b) < SizeOfIfInfoMsg {
		return nil, errors.New("IfInfoMsg to short")
	}
	result := IfInfoMsg{
		unix.IfInfomsg{
			// byte 0
			Family: b[0],
			// byte 1 is skipped
			// bytes 2-3: Type
			Type: NativeEndian.Uint16(b[2:4]),
			// bytes 4-7: Index
			Index: int32(NativeEndian.Uint32(b[4:8])),
			// bytes 8-11: Flags
			Flags: NativeEndian.Uint32(b[8:12]),
			// bytes 12-15: Change
			Change: NativeEndian.Uint32(b[12:16]),
		},
	}
	return &result, nil
}
