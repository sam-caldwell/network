package core

// Serialize - Serialize the TcRateSpec struct into a byte slice.
// This function ensures safe serialization by manually converting each field
// to a byte slice using the binary encoding package.
func (msg *TcRateSpec) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofTcRateSpec)
	buf[0] = msg.CellLog   // Directly assign for uint8
	buf[1] = msg.Linklayer // Directly assign for uint8
	NativeEndian.PutUint16(buf[2:], msg.Overhead)
	NativeEndian.PutUint16(buf[4:], uint16(msg.CellAlign))
	NativeEndian.PutUint16(buf[6:], msg.Mpu)
	NativeEndian.PutUint32(buf[8:], msg.Rate)
	return buf, nil
}
