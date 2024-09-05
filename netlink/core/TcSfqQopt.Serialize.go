package core

// Serialize - Safely converts the TcSfqQopt structure into a byte slice.
// This function manually encodes each field into the byte slice.
func (x *TcSfqQopt) Serialize() []byte {
	buf := make([]byte, SizeofTcSfqQopt)

	// Encode each field into the byte slice using the appropriate byte size and endianness
	buf[0] = x.Quantum
	NativeEndian.PutUint32(buf[1:5], uint32(x.Perturb))
	NativeEndian.PutUint32(buf[5:9], x.Limit)
	buf[9] = x.Divisor
	buf[10] = x.Flows

	return buf
}
