package core

import (
	"encoding/binary"
	"unsafe"
)

const (
	SizeOfTcU32Sel = int(unsafe.Sizeof(TcU32Sel{}))
	SizeOfTcU32Key = int(unsafe.Sizeof(TcU32Key{}))
)

// Serialize converts the TcU32Sel object into a byte slice safely.
// It manually converts the fields and the keys into a byte slice using the encoding/binary package.
func (msg *TcU32Sel) Serialize() ([]byte, error) {
	// Allocate a buffer large enough to hold the entire structure, including keys.
	buf := make([]byte, msg.Len())

	// Serialize individual fields into the buffer
	buf[0] = msg.Flags
	buf[1] = msg.Offshift
	buf[2] = msg.Nkeys
	buf[3] = msg.Pad
	binary.BigEndian.PutUint16(buf[4:], msg.Offmask)
	NativeEndian.PutUint16(buf[6:], msg.Off)
	NativeEndian.PutUint16(buf[8:], uint16(msg.Offoff))
	NativeEndian.PutUint16(buf[10:], uint16(msg.Hoff))
	binary.BigEndian.PutUint32(buf[12:], msg.Hmask)

	// Serialize each key in the Keys array
	next := SizeOfTcU32Sel
	for _, key := range msg.Keys {
		keyBuf, err := key.Serialize()
		if err != nil {
			return nil, err
		}
		copy(buf[next:], keyBuf)
		next += SizeOfTcU32Key
	}

	return buf, nil
}
