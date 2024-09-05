package core

// Len - returns the size of the TcGen structure in bytes.
// The function calculates the total memory footprint of the TcGen struct,
// which is useful for serialization and low-level operations.
//
// SizeofTcGen is assumed to be a predefined constant representing the size
// of the TcGen structure. This size corresponds to the sum of the sizes
// of all fields: Index (uint32), Capab (uint32), Action (int32), Refcnt (int32), and Bindcnt (int32).
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_common.h
func (msg *TcGen) Len() int {
	return SizeofTcGen
}
