package core

// Len returns the total size of the TcU32Sel object in bytes.
//
// The size is calculated as the base size of TcU32Sel plus the size of all keys in the Keys array.
// Each key is represented by a TcU32Key, and the total number of keys is stored in the Nkeys field.
func (msg *TcU32Sel) Len() int {
	return SizeofTcU32Sel + int(msg.Nkeys)*SizeofTcU32Key
}
