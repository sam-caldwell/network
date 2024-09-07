package core

import (
	"errors"
)

// Deserialize - Deserialize the CnMsgOp struct
func (msg *CnMsgOp) Deserialize(data []byte) error {
	if len(data) < SizeofCnMsgOp {
		return errors.New("data too short to deserialize CnMsgOp")
	}

	// Deserialize CbID (Idx and Val)
	msg.ID.Idx = NativeEndian.Uint32(data[0:])
	msg.ID.Val = NativeEndian.Uint32(data[4:])

	// Deserialize Seq, Ack, Length, Flags, and Op
	msg.Seq = NativeEndian.Uint32(data[8:])
	msg.Ack = NativeEndian.Uint32(data[12:])
	msg.Length = NativeEndian.Uint16(data[16:])
	msg.Flags = NativeEndian.Uint16(data[18:])
	msg.Op = NativeEndian.Uint32(data[20:])

	return nil
}
