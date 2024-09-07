package core

// Serialize - Safely serializes TcNetemReorder into a byte slice using binary encoding.
func (msg *TcNetemReorder) Serialize() ([]byte, error) {

	buf := make([]byte, SizeOfTcNetemReorder)

	// Serialize each field
	NativeEndian.PutUint32(buf[0:], msg.Probability)
	NativeEndian.PutUint32(buf[4:], msg.Correlation)

	return buf, nil

}
