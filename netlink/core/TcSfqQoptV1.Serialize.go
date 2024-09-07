package core

// Serialize safely converts the TcSfqQoptV1 structure into a byte slice.
// This method manually encodes each field into the byte slice, ensuring
// that unsafe pointers are not used.
func (lhs *TcSfqQoptV1) Serialize() ([]byte, error) {
	// Create a byte slice with the exact size of the structure
	buf := make([]byte, lhs.Len())

	// Serialize the base TcSfqQopt structure
	data, err := lhs.TcSfqQopt.Serialize()
	if err != nil {
		return nil, err
	}
	copy(buf[:SizeOfTcSfqQopt], data)

	// Serialize the additional fields in TcSfqQoptV1
	NativeEndian.PutUint32(buf[SizeOfTcSfqQopt:], lhs.Depth)
	NativeEndian.PutUint32(buf[SizeOfTcSfqQopt+4:], lhs.HeadDrop)
	NativeEndian.PutUint32(buf[SizeOfTcSfqQopt+8:], lhs.Limit)
	NativeEndian.PutUint32(buf[SizeOfTcSfqQopt+12:], lhs.QthMin)
	NativeEndian.PutUint32(buf[SizeOfTcSfqQopt+16:], lhs.QthMax)
	buf[SizeOfTcSfqQopt+20] = lhs.Wlog
	buf[SizeOfTcSfqQopt+21] = lhs.Plog
	buf[SizeOfTcSfqQopt+22] = lhs.ScellLog
	buf[SizeOfTcSfqQopt+23] = lhs.Flags
	NativeEndian.PutUint32(buf[SizeOfTcSfqQopt+24:], lhs.MaxP)

	// Serialize the embedded TcSfqRedStats structure
	data, err = lhs.TcSfqQopt.Serialize()
	if err != nil {
		return nil, err
	}
	copy(buf[SizeOfTcSfqQopt+28:], data)

	return buf, nil
}
