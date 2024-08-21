package core

import (
	"golang.org/x/sys/unix"
)

// Serialize - serialize IfaCacheInfo to byte slice
//
//	struct ifa_cacheinfo {
//		__u32	ifa_prefered;
//		__u32	ifa_valid;
//		__u32	cstamp; /* created timestamp, hundredths of seconds */
//		__u32	tstamp; /* updated timestamp, hundredths of seconds */
//	};
func (msg *IfaCacheInfo) Serialize() []byte {
	length := unix.SizeofIfaCacheinfo
	b := make([]byte, length)
	nativeEndian.PutUint32(b[0:4], msg.Prefered)
	nativeEndian.PutUint32(b[4:8], msg.Valid)
	nativeEndian.PutUint32(b[8:12], msg.Cstamp)
	nativeEndian.PutUint32(b[12:16], msg.Tstamp)
	return b
}
