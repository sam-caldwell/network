package core

// Len - Return the size of the interface address message
func (msg *IfAddressMessage) Len() int {
	return SizeofIfAddrmsg
}
