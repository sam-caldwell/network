package core

// Len - returns the size of XfrmMark struct in bytes
func (msg *XfrmMark) Len() int {
	return SizeOfXfrmMark
}
