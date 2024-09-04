package core

// TCPDiagSockID represents the socket identity in the TCP diagnostic subsystem.
// This structure is useful for identifying the state and configuration of TCP connections in diagnostic tools.
//
// In the Linux kernel, this struct is defined to help monitor the state of TCP connections via the netlink socket interface.
// Reference: See the Linux kernel source code for TCP diagnostic support (tcp_diag).
//
// Struct fields mirror their C counterparts:
//
// - Sport: Source port of the TCP connection (tcpdiag_sport).
// - Dport: Destination port of the TCP connection (tcpdiag_dport).
// - Src: Source IP address, represented as an array of 4 x 32-bit fields to support IPv4/IPv6 (tcpdiag_src).
// - Dst: Destination IP address, also represented as an array for IPv4/IPv6 (tcpdiag_dst).
// - If: Interface index (tcpdiag_if).
// - Cookie: Unique identifier for the connection (tcpdiag_cookie), used for correlating connection tracking entries across namespaces.
//
// TCPDiagNoCookie is a constant that defines an absence of a valid cookie, corresponding to `~0U` in the C code.

type TCPDiagSockID struct {
	Sport  uint16    // tcpdiag_sport: Source port
	Dport  uint16    // tcpdiag_dport: Destination port
	Src    [4]uint32 // tcpdiag_src: Source IPv4/IPv6 address (4 x 32-bit fields)
	Dst    [4]uint32 // tcpdiag_dst: Destination IPv4/IPv6 address (4 x 32-bit fields)
	If     uint32    // tcpdiag_if: Interface index
	Cookie [2]uint32 // tcpdiag_cookie: Cookie for the connection
}
