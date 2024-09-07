package core

// Serialize - Safely converts the TcSfqQopt structure into a byte slice.
// This function manually encodes each field into the byte slice.
func (msg *TcSfqQopt) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfTcSfqQopt)

	// Encode each field into the byte slice using the appropriate byte size and endianness
	buf[0] = msg.Quantum
	NativeEndian.PutUint32(buf[1:5], uint32(msg.Perturb))
	NativeEndian.PutUint32(buf[5:9], msg.Limit)
	buf[9] = msg.Divisor
	buf[10] = msg.Flows

	return buf, nil
}
