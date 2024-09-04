package core

// TCPDiagNoCookie defines the absence of a valid cookie.
// Reference: See the Linux kernel source code (tcp_diag) to understand how this is used.
const TCPDiagNoCookie = ^uint32(0) // ~0U in C
