package core

// Serialize serializes TcNetemCorr into a byte slice using binary encoding.
func (msg *TcNetemCorr) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofTcNetemCorr)

	// Binary encoding of each field
	NativeEndian.PutUint32(buf[0:], msg.DelayCorr)
	NativeEndian.PutUint32(buf[4:], msg.LossCorr)
	NativeEndian.PutUint32(buf[8:], msg.DupCorr)

	return buf, nil
}
