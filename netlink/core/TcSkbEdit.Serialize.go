package core

// Serialize converts the TcSkbEdit structure into a byte slice in a safe way.
func (msg *TcSkbEdit) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofTcSkbEdit)

	// Write each field to the buffer
	NativeEndian.PutUint32(buf[0:4], msg.Index)
	NativeEndian.PutUint32(buf[4:8], msg.Capab)
	NativeEndian.PutUint32(buf[8:12], uint32(msg.Action))
	NativeEndian.PutUint32(buf[12:16], uint32(msg.Refcnt))
	NativeEndian.PutUint32(buf[16:20], uint32(msg.Bindcnt))

	return buf, nil
}
