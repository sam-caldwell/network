package core

// Len - Return the size of XfrmUserSaInfo struct (in bytes)
func (msg *XfrmUserSaInfo) Len() int {
	return SizeOfXfrmUserSaInfo
}
