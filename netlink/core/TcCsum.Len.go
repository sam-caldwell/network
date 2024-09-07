package core

// Len - Returns the size of the TcCsum structure in bytes.
//
// The TcCsum structure size is determined by the SizeOfTcCsum constant, which reflects the size of this struct
// in the Linux kernel. This function is essential when serializing or deserializing the structure in
// networking contexts, such as when interacting with netlink messages or other low-level kernel operations.
//
// For more details, refer to:
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
func (msg *TcCsum) Len() int {
	return SizeOfTcCsum
}
