package core

// Serialize - Serialize the CnMsgOp struct
func (msg *CnMsgOp) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofCnMsgOp)

	// Serialize CbID (Idx and Val)
	NativeEndian.PutUint32(buf[0:], msg.ID.Idx)
	NativeEndian.PutUint32(buf[4:], msg.ID.Val)

	// Serialize Seq, Ack, Length, Flags, and Op
	NativeEndian.PutUint32(buf[8:], msg.Seq)
	NativeEndian.PutUint32(buf[12:], msg.Ack)
	NativeEndian.PutUint16(buf[16:], msg.Length)
	NativeEndian.PutUint16(buf[18:], msg.Flags)
	NativeEndian.PutUint32(buf[20:], msg.Op)

	return buf, nil
}
