package core

// Len returns the size of the TcSkbEdit structure in bytes.
// This method is used to calculate the total size of the TcSkbEdit structure
// in memory, which helps during the process of serializing or deserializing
// the structure.
//
// The size of TcSkbEdit is determined based on the size of its fields,
// ensuring it aligns correctly with the underlying system architecture.
//
// See the related Linux kernel code:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
func (msg *TcSkbEdit) Len() int {
	return SizeofTcSkbEdit
}
