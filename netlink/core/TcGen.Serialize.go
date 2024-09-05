package core

// Serialize - converts the TcGen structure into a byte slice.
func (msg *TcGen) Serialize() []byte {
	buf := make([]byte, SizeofTcGen)

	// Serialize fields in order using binary encoding
	NativeEndian.PutUint32(buf[0:], msg.Index)
	NativeEndian.PutUint32(buf[4:], msg.Capab)
	NativeEndian.PutUint32(buf[8:], uint32(msg.Action))
	NativeEndian.PutUint32(buf[12:], uint32(msg.Refcnt))
	NativeEndian.PutUint32(buf[16:], uint32(msg.Bindcnt))

	return buf
}
