package core

import "encoding/binary"

// Serialize converts the TcU32Sel object into a byte slice safely.
// It manually converts the fields and the keys into a byte slice using the encoding/binary package.
func (msg *TcU32Sel) Serialize() []byte {
	// Allocate a buffer large enough to hold the entire structure, including keys.
	buf := make([]byte, msg.Len())

	// Serialize individual fields into the buffer
	buf[0] = msg.Flags
	buf[1] = msg.Offshift
	buf[2] = msg.Nkeys
	buf[3] = msg.Pad
	binary.BigEndian.PutUint16(buf[4:], msg.Offmask)
	binary.LittleEndian.PutUint16(buf[6:], msg.Off)
	binary.LittleEndian.PutUint16(buf[8:], uint16(msg.Offoff))
	binary.LittleEndian.PutUint16(buf[10:], uint16(msg.Hoff))
	binary.BigEndian.PutUint32(buf[12:], msg.Hmask)

	// Serialize each key in the Keys array
	next := SizeofTcU32Sel
	for _, key := range msg.Keys {
		keyBuf := key.Serialize()
		copy(buf[next:], keyBuf)
		next += SizeofTcU32Key
	}

	return buf
}
