package core

// Serialize - Converts the TcNetemRate structure into a byte slice using binary encoding.
func (msg *TcNetemRate) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfTcRateSpec)
	NativeEndian.PutUint32(buf[0:], msg.Rate)
	NativeEndian.PutUint32(buf[4:], uint32(msg.PacketOverhead))
	NativeEndian.PutUint32(buf[8:], msg.CellSize)
	NativeEndian.PutUint32(buf[12:], uint32(msg.CellOverhead))
	return buf, nil
}
