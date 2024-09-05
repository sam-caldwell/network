package core

import (
	"encoding/binary"
)

// Serialize - Converts the TcPolice structure into a byte slice safely.
// This method manually encodes each field in the TcPolice struct to a byte slice,
// ensuring portability and safety by avoiding the use of unsafe pointers.
func (msg *TcPolice) Serialize() []byte {
	buf := make([]byte, SizeofTcPolice)

	// Manually serialize each field using the appropriate endianness
	binary.NativeEndian.PutUint32(buf[0:4], msg.Index)
	binary.NativeEndian.PutUint32(buf[4:8], uint32(msg.Action))
	binary.NativeEndian.PutUint32(buf[8:12], msg.Limit)
	binary.NativeEndian.PutUint32(buf[12:16], msg.Burst)
	binary.NativeEndian.PutUint32(buf[16:20], msg.Mtu)
	r1 := msg.Rate.Serialize()
	copy(buf[20:20+len(r1)], r1)
	r2 := msg.PeakRate.Serialize()
	copy(buf[20+len(r2):20+2*len(r2)], r2)
	binary.NativeEndian.PutUint32(buf[36:40], uint32(msg.Refcnt))
	binary.NativeEndian.PutUint32(buf[40:44], uint32(msg.Bindcnt))
	binary.NativeEndian.PutUint32(buf[44:48], msg.Capab)

	return buf
}
