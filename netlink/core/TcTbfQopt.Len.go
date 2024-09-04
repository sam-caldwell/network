package core

// Len - Returns the size of the TcTbfQopt structure in bytes.
// This method is used to determine the memory size required for the struct when working with low-level system programming.
// The constant SizeofTcTbfQopt represents the size of this structure, which is defined elsewhere in the code.
func (msg *TcTbfQopt) Len() int {
	return SizeofTcTbfQopt
}
