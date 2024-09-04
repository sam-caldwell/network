package core

// TcMsg - represents the tcmsg struct from rtnetlink used for traffic control (TC) operations.
//
// This structure is used in netlink communication to manipulate traffic control subsystems such as queuing disciplines,
// filters, and actions in the Linux kernel.
//
// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
type TcMsg struct {
	Family  uint8   // Address family (AF_UNSPEC, AF_INET, AF_INET6, etc.)
	Pad     [3]byte // Padding to align the struct properly.
	Ifindex int32   // Interface index the message refers to.
	Handle  uint32  // TC handle identifier (combination of major/minor numbers for qdiscs, classes, or filters).
	Parent  uint32  // Parent identifier for hierarchical TC structures.
	Info    uint32  // Additional information (e.g., scheduling or filter-specific data).
}
