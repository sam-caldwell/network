package core

// Serialize - Serializes the TcHtbCopt structure into a byte slice using binary encoding.
// This method uses the binary package to safely convert each field into its byte representation.
func (msg *TcHtbCopt) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofTcHtbCopt)

	// Serialize the Rate field
	data, err := msg.Rate.Serialize()
	if err != nil {
		return nil, err
	}
	copy(buf[0:], data)

	// Serialize the Ceil field, which starts after Rate
	data, err = msg.Rate.Serialize()
	if err != nil {
		return nil, err
	}
	copy(buf[SizeofTcRateSpec:], data)

	// Serialize the remaining fields (Buffer, Cbuffer, Quantum, Level, Prio)
	NativeEndian.PutUint32(buf[2*SizeofTcRateSpec:], msg.Buffer)
	NativeEndian.PutUint32(buf[2*SizeofTcRateSpec+4:], msg.Cbuffer)
	NativeEndian.PutUint32(buf[2*SizeofTcRateSpec+8:], msg.Quantum)
	NativeEndian.PutUint32(buf[2*SizeofTcRateSpec+12:], msg.Level)
	NativeEndian.PutUint32(buf[2*SizeofTcRateSpec+16:], msg.Prio)

	return buf, nil
}
