package core

// Len - return the length of the RtGenMsg struct in bytes
func (msg *RtGenMsg) Len() int {

	return rtaAlignOf(SizeOfRtGenMsg)

}
