package core

// Len - return the length of the XfrmUserSpiInfo struct in bytes
func (msg *XfrmUserSpiInfo) Len() int {
	return SizeOfXfrmUserSpiInfo
}
