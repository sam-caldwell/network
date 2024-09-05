package core

// Len - Returns the size of the TcPolice structure in bytes.
//
// This function is used to retrieve the size of the TcPolice struct, which is
// important when working with netlink messages or other low-level networking
// interactions that require the exact byte size of data structures.
//
// SizeofTcPolice is a constant defined elsewhere that represents the size of the TcPolice struct.
func (msg *TcPolice) Len() int {
	return SizeofTcPolice
}
