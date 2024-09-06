package core

// TcPeditKeyEx represents the extended key structure used in the pedit action within the Linux kernel's traffic control (TC).
// This structure allows specifying additional commands and header types for more granular packet modifications.
//
// The `HeaderType` field defines the type of header being modified (e.g., Ethernet, IPv4, IPv6, TCP, UDP),
// and the `Cmd` field specifies the command to be executed (e.g., set or add to the value).
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type TcPeditKeyEx struct {
	// HeaderType specifies the type of header being modified (e.g., Ethernet, IPv4, IPv6, TCP, UDP).
	HeaderType PeditHeaderType

	// Cmd specifies the command to be executed, such as setting or adding to the value.
	Cmd PeditCmd
}
