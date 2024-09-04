package core

import "encoding/binary"

// Serialize serializes TcNetemCorr into a byte slice using binary encoding.
func (msg *TcNetemCorr) Serialize() []byte {
	buf := make([]byte, SizeofTcNetemCorr)

	// Binary encoding of each field
	binary.LittleEndian.PutUint32(buf[0:], msg.DelayCorr)
	binary.LittleEndian.PutUint32(buf[4:], msg.LossCorr)
	binary.LittleEndian.PutUint32(buf[8:], msg.DupCorr)

	return buf
}
