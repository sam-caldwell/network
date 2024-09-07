package core

// Len - return the length of XfrmAlgoAEAD struct in bytes
func (msg *XfrmAlgoAEAD) Len() int {
	return SizeofXfrmAlgoAEAD + int(msg.AlgKeyLen/8)
}
