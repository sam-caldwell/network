package core

// Serialize - Converts the TcSfqRedStats structure into a byte slice safely.
// This method manually encodes each field in the TcSfqRedStats struct to a byte slice,
// ensuring portability and safety by avoiding the use of unsafe pointers.
func (msg *TcSfqRedStats) Serialize() []byte {
	buf := make([]byte, SizeofTcSfqRedStats)

	// Manually serialize each field using the appropriate endianness
	NativeEndian.PutUint32(buf[0:4], msg.ProbDrop)
	NativeEndian.PutUint32(buf[4:8], msg.ForcedDrop)
	NativeEndian.PutUint32(buf[8:12], msg.ProbMark)
	NativeEndian.PutUint32(buf[12:16], msg.ForcedMark)
	NativeEndian.PutUint32(buf[16:20], msg.ProbMarkHead)
	NativeEndian.PutUint32(buf[20:24], msg.ForcedMarkHead)

	return buf
}
