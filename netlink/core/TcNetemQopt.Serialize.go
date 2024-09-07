package core

// Serialize - Serialize the TcNetemQopt struct into a byte slice.
// This function ensures safe serialization by manually converting each field
// to a byte slice using the binary encoding package.
//
// Returns the serialized byte slice.
func (msg *TcNetemQopt) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfTcNetemQopt)
	NativeEndian.PutUint32(buf[0:], msg.Latency)
	NativeEndian.PutUint32(buf[4:], msg.Limit)
	NativeEndian.PutUint32(buf[8:], msg.Loss)
	NativeEndian.PutUint32(buf[12:], msg.Gap)
	NativeEndian.PutUint32(buf[16:], msg.Duplicate)
	NativeEndian.PutUint32(buf[20:], msg.Jitter)
	return buf, nil
}
