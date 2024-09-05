package core

// Serialize safely converts the TcSfqQoptV1 structure into a byte slice.
// This method manually encodes each field into the byte slice, ensuring
// that unsafe pointers are not used.
func (x *TcSfqQoptV1) Serialize() []byte {
	// Create a byte slice with the exact size of the structure
	buf := make([]byte, x.Len())

	// Serialize the base TcSfqQopt structure
	copy(buf[:SizeofTcSfqQopt], x.TcSfqQopt.Serialize())

	// Serialize the additional fields in TcSfqQoptV1
	NativeEndian.PutUint32(buf[SizeofTcSfqQopt:], x.Depth)
	NativeEndian.PutUint32(buf[SizeofTcSfqQopt+4:], x.HeadDrop)
	NativeEndian.PutUint32(buf[SizeofTcSfqQopt+8:], x.Limit)
	NativeEndian.PutUint32(buf[SizeofTcSfqQopt+12:], x.QthMin)
	NativeEndian.PutUint32(buf[SizeofTcSfqQopt+16:], x.QthMax)
	buf[SizeofTcSfqQopt+20] = x.Wlog
	buf[SizeofTcSfqQopt+21] = x.Plog
	buf[SizeofTcSfqQopt+22] = x.ScellLog
	buf[SizeofTcSfqQopt+23] = x.Flags
	NativeEndian.PutUint32(buf[SizeofTcSfqQopt+24:], x.MaxP)

	// Serialize the embedded TcSfqRedStats structure
	copy(buf[SizeofTcSfqQopt+28:], x.TcSfqRedStats.Serialize())

	return buf
}
