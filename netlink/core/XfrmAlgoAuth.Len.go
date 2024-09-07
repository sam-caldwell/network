package core

// Len - return the size of the XfrmAlgoAuth struct in bytes
func (msg *XfrmAlgoAuth) Len() int {
	return SizeOfXfrmAlgoAuth + int(msg.AlgKeyLen/8)
}
