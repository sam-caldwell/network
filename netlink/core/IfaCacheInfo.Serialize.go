package core

import (
	"golang.org/x/sys/unix"
)

// Serialize - ifa_cacheinfo - serialize IfaCacheInfo to byte slice
func (msg *IfaCacheInfo) Serialize() []byte {

	length := unix.SizeofIfaCacheinfo

	b := make([]byte, length)

	NativeEndian.PutUint32(b[0:4], msg.Prefered)

	NativeEndian.PutUint32(b[4:8], msg.Valid)

	NativeEndian.PutUint32(b[8:12], msg.Cstamp)

	NativeEndian.PutUint32(b[12:16], msg.Tstamp)

	return b

}
