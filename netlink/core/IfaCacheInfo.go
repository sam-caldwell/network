package core

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
)

// IfaCacheInfo - Wrapper around unix.IfaCacheinfo with additional functionality.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_addr.h
type IfaCacheInfo struct {
	unix.IfaCacheinfo
}

const (
	// SizeOfIfaCacheInfo - 0x16 // bytes as derived from unix.SizeOfIfaCacheInfo
	SizeOfIfaCacheInfo = int(unsafe.Sizeof(IfaCacheInfo{}))
)

// Len - return the length (in bytes) of the IfaCacheInfo struct.
func (msg *IfaCacheInfo) Len() int {

	return SizeOfIfaCacheInfo

}

// Serialize - ifa_cacheinfo - serialize IfaCacheInfo to byte slice
func (msg *IfaCacheInfo) Serialize() ([]byte, error) {

	length := SizeOfIfaCacheInfo

	b := make([]byte, length)

	NativeEndian.PutUint32(b[0:4], msg.Prefered)
	NativeEndian.PutUint32(b[4:8], msg.Valid)
	NativeEndian.PutUint32(b[8:12], msg.Cstamp)
	NativeEndian.PutUint32(b[12:16], msg.Tstamp)

	return b, nil

}

// String provides a human-readable representation of IfaCacheInfo.
//
// From unix package...
//
//	type IfaCacheinfo struct {
//	   Prefered uint32
//	   Valid    uint32
//	   Cstamp   uint32
//	   Tstamp   uint32
//	}
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_addr.h
func (msg *IfaCacheInfo) String() string {

	return fmt.Sprintf("Valid: %d, Prefered: %d, Created: %d, Updated: %d",
		msg.Valid, msg.Prefered, msg.Cstamp, msg.Tstamp)

}

// Deserialize method converts a byte slice to an IfaCacheInfo struct
//
//	type IfaCacheinfo struct {
//	   Prefered uint32
//	   Valid    uint32
//	   Cstamp   uint32
//	   Tstamp   uint32
//	}
func (msg *IfaCacheInfo) Deserialize(data []byte) error {
	if len(data) < SizeOfIfaCacheInfo {
		return errors.New(ErrInputTooShort)
	}
	msg.Prefered = NativeEndian.Uint32(data[0:4])
	msg.Valid = NativeEndian.Uint32(data[4:8])
	msg.Cstamp = NativeEndian.Uint32(data[8:12])
	msg.Tstamp = NativeEndian.Uint32(data[12:16])
	return nil
}

// DeserializeIfaCacheInfo converts a byte slice to an IfaCacheInfo struct
func DeserializeIfaCacheInfo(data []byte) (*IfaCacheInfo, error) {
	if len(data) < SizeOfIfaCacheInfo {
		return nil, errors.New("data slice too short")
	}
	var info IfaCacheInfo
	if err := info.Deserialize(data); err != nil {
		return nil, err
	}
	return &info, nil
}
