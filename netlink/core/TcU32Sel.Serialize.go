package core

import (
	"encoding/binary"
)

// Serialize - converts the TcU32Sel object into a byte slice safely.
func (msg *TcU32Sel) Serialize() ([]byte, error) {
	// Calculate the total size needed for the buffer
	totalSize := SizeOfTcU32Sel + len(msg.Keys)*SizeOfTcU32Key
	buf := make([]byte, totalSize)

	// Serialize the fixed part of the struct
	buf[0] = msg.Flags
	buf[1] = msg.Offshift
	buf[2] = msg.Nkeys
	buf[3] = msg.Pad
	binary.BigEndian.PutUint16(buf[4:], msg.Offmask)
	NativeEndian.PutUint16(buf[6:], msg.Off)
	NativeEndian.PutUint16(buf[8:], uint16(msg.Offoff))
	NativeEndian.PutUint16(buf[10:], uint16(msg.Hoff))
	binary.BigEndian.PutUint32(buf[12:], msg.Hmask)

	// Serialize each key and append it to the buffer
	offset := SizeOfTcU32Sel // Start after the fixed-size portion
	for _, key := range msg.Keys {
		keyBuf, err := key.Serialize()
		if err != nil {
			return nil, err
		}
		// Copy the serialized key data into the buffer at the correct offset
		copy(buf[offset:offset+SizeOfTcU32Key], keyBuf)
		offset += SizeOfTcU32Key
	}

	return buf, nil
}
