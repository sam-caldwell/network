package core

// Len - return size of th TcHtbGlob (Hierarchical Token Bucket) struct in bytes
func (msg *TcHtbGlob) Len() int {
	return SizeofTcHtbGlob
}
