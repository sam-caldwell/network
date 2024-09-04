package core

// Len - Return the size of the TcActionMsg object (in bytes).
//
// This method calculates the size of the `TcActionMsg` struct.
// The `SizeofTcActionMsg` constant should represent the size of the structure in bytes,
// ensuring that memory alignment and padding are considered.
//
// Reference:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
func (msg *TcActionMsg) Len() int {
	return SizeofTcActionMsg
}
