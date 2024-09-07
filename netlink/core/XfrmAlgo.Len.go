package core

// Len - Return the length in bytes for the XfrmAlgo struct
func (msg *XfrmAlgo) Len() int {
	return SizeofXfrmAlgo + int(msg.AlgKeyLen/8)
}
