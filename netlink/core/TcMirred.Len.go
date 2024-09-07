package core

// Len - Returns the size of the TcMirred structure in bytes.
// This function is used when serializing or deserializing the TcMirred structure.
// It ensures the proper size is allocated when working with netlink messages.
//
// The TcMirred structure is part of the Linux Traffic Control system, specifically used for mirroring and redirection actions.
//
// See:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
func (msg *TcMirred) Len() int {
	return SizeOfTcMirred
}
