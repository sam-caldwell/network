package core

import (
	"encoding/binary"
	"errors"
)

// Deserialize - Deserialize the CnMsgOp struct
func (msg *CnMsgOp) Deserialize(data []byte) error {
	if len(data) < SizeofCnMsgOp {
		return errors.New("data too short to deserialize CnMsgOp")
	}

	// Deserialize CbID (Idx and Val)
	msg.ID.Idx = binary.LittleEndian.Uint32(data[0:])
	msg.ID.Val = binary.LittleEndian.Uint32(data[4:])

	// Deserialize Seq, Ack, Length, Flags, and Op
	msg.Seq = binary.LittleEndian.Uint32(data[8:])
	msg.Ack = binary.LittleEndian.Uint32(data[12:])
	msg.Length = binary.LittleEndian.Uint16(data[16:])
	msg.Flags = binary.LittleEndian.Uint16(data[18:])
	msg.Op = binary.LittleEndian.Uint32(data[20:])

	return nil
}
