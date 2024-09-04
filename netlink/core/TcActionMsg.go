package core

// TcActionMsg - This struct represents the `tcamsg` structure used in the Traffic Control (TC) subsystem in Linux.
//
// This structure is used to carry traffic control action messages (like `qdisc`, `filter`, etc.).
// The `Family` field typically represents the address family (e.g., `AF_UNSPEC`, `AF_INET`, etc.),
// while the `Pad` field is used for alignment padding.
//
// Reference:
// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
// - Documentation: https://man7.org/linux/man-pages/man7/rtnetlink.7.html
type TcActionMsg struct {
	Family uint8   // Address family (e.g., AF_UNSPEC, AF_INET)
	Pad    [3]byte // Padding for alignment (unused)
}
