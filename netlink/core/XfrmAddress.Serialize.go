package core

// Serialize safely converts the XfrmAddress structure into a byte slice.
// This method manually encodes each byte of the array into a slice to avoid using unsafe pointers.
func (x *XfrmAddress) Serialize() ([]byte, error) {
	// Create a byte slice with the exact size of the XfrmAddress
	buf := make([]byte, SizeofXfrmAddress)

	// Copy the content of the XfrmAddress into the byte slice
	copy(buf, x[:])

	return buf, nil
}
