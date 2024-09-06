package core

// XfrmAddress represents an IP address used in XFRM (IPsec) related structures.
//
// In the Linux kernel, the `xfrm_address_t` is defined as a union of a 32-bit
// IPv4 address (`a4`) and a 128-bit IPv6 address (`a6`), allowing flexibility
// to handle both IPv4 and IPv6 addresses in IPsec security policies and states.
//
// This Go struct is defined as a fixed-size byte array that can accommodate either
// an IPv4 or IPv6 address, corresponding to the original union definition.
//
//	typedef union {
//	  __be32    a4;
//	  __be32    a6[4];
//	} xfrm_address_t;
//
// References:
// - Linux Kernel: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmAddress [SizeofXfrmAddress]byte
