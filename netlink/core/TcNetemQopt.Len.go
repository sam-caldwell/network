package core

// Len returns the size of the TcNetemQopt structure in bytes.
func (msg *TcNetemQopt) Len() int {
	return SizeofTcNetemQopt
}
