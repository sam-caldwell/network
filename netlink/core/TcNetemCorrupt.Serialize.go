package core

// Serialize - Converts the TcNetemCorrupt object into a byte slice using binary encoding.
func (msg *TcNetemCorrupt) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofTcNetemCorrupt)
	NativeEndian.PutUint32(buf[0:], msg.Probability)
	NativeEndian.PutUint32(buf[4:], msg.Correlation)
	return buf, nil
}
