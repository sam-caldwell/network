package core

// Serialize - Converts the TcPolice structure into a byte slice safely.
// This method manually encodes each field in the TcPolice struct to a byte slice,
// ensuring portability and safety by avoiding the use of unsafe pointers.
func (msg *TcPolice) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofTcPolice)

	// Manually serialize each field using the appropriate endianness
	NativeEndian.PutUint32(buf[0:4], msg.Index)
	NativeEndian.PutUint32(buf[4:8], uint32(msg.Action))
	NativeEndian.PutUint32(buf[8:12], msg.Limit)
	NativeEndian.PutUint32(buf[12:16], msg.Burst)
	NativeEndian.PutUint32(buf[16:20], msg.Mtu)
	r1, err := msg.Rate.Serialize()
	if err != nil {
		return nil, err
	}
	copy(buf[20:20+len(r1)], r1)
	r2, err := msg.PeakRate.Serialize()
	if err != nil {
		return nil, err
	}
	copy(buf[20+len(r2):20+2*len(r2)], r2)
	NativeEndian.PutUint32(buf[36:40], uint32(msg.Refcnt))
	NativeEndian.PutUint32(buf[40:44], uint32(msg.Bindcnt))
	NativeEndian.PutUint32(buf[44:48], msg.Capab)

	return buf, nil
}
