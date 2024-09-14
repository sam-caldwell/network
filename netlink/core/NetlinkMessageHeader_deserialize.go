package core

import (
	"golang.org/x/sys/unix"
)

// DeserializeNetlinkMessageHeader deserializes a byte slice into a unix.NlMsghdr structure using native endian.
//
//	type NlMsghdr struct {
//	   Len   uint32
//	   Type  uint16
//	   Flags uint16
//	   Seq   uint32
//	   Pid   uint32
//	}
func DeserializeNetlinkMessageHeader(data []byte) (*unix.NlMsghdr, error) {

	if err := checkInputSize(data, NetlinkMessageHeaderSize, disableSizeCheck); err != nil {
		return nil, err
	}

	return &(unix.NlMsghdr{
		Len:   NativeEndian.Uint32(data[0:4]),
		Type:  NativeEndian.Uint16(data[4:6]),
		Flags: NativeEndian.Uint16(data[6:8]),
		Seq:   NativeEndian.Uint32(data[8:12]),
		Pid:   NativeEndian.Uint32(data[12:16]),
	}), nil

}
