package core

import (
	"encoding/binary"
)

// Serialize - Serialize the CnMsgOp struct
func (msg *CnMsgOp) Serialize() []byte {
	buf := make([]byte, SizeofCnMsgOp)

	// Serialize CbID (Idx and Val)
	binary.LittleEndian.PutUint32(buf[0:], msg.ID.Idx)
	binary.LittleEndian.PutUint32(buf[4:], msg.ID.Val)

	// Serialize Seq, Ack, Length, Flags, and Op
	binary.LittleEndian.PutUint32(buf[8:], msg.Seq)
	binary.LittleEndian.PutUint32(buf[12:], msg.Ack)
	binary.LittleEndian.PutUint16(buf[16:], msg.Length)
	binary.LittleEndian.PutUint16(buf[18:], msg.Flags)
	binary.LittleEndian.PutUint32(buf[20:], msg.Op)

	return buf
}
