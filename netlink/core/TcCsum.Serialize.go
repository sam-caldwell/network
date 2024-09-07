package core

// Serialize - Converts the TcCsum struct into a byte slice.
// This method ensures safe serialization by manually converting each field to a byte slice using the binary encoding package.
//
// Returns the serialized byte slice.
func (msg *TcCsum) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfTcCsum)

	// Serialize fields
	NativeEndian.PutUint32(buf[0:], msg.TcGen.Index)
	NativeEndian.PutUint32(buf[4:], msg.TcGen.Capab)
	NativeEndian.PutUint32(buf[8:], uint32(msg.TcGen.Action))
	NativeEndian.PutUint32(buf[12:], uint32(msg.TcGen.Refcnt))
	NativeEndian.PutUint32(buf[16:], uint32(msg.TcGen.Bindcnt))
	NativeEndian.PutUint32(buf[20:], msg.UpdateFlags)

	return buf, nil
}
