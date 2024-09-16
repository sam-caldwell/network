package core

// SocketCommand - commands used to manage sockets
type SocketCommand int

const (

	// SockDiagByFamily - SOCK_DIAG_BY_FAMILY - sock_diag_by_family
	// Socket diagnostics by family.
	// Provides a mechanism to get detailed information about sockets for a specific protocol family.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/sock_diag.h
	SockDiagByFamily SocketType = 0x14

	// SockDestroy - SOCK_DESTROY - sock_destroy
	// Constant used internally to destroy a Socket.
	// This is typically used for Socket cleanup operations in the kernel.
	// https://github.com/torvalds/linux/blob/master/net/core/Sock.c
	SockDestroy SocketCommand = 0x15

	// SockIocType - SOCK_IOC_TYPE - sock_ioc_type
	// IOCTL command type for Sockets.
	// Defines the type of IOCTL (Input/Output Control) command that can be issued on a Socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/Sockios.h
	SockIocType SocketType = 0x89

	// SockSndbufLock - SOCK_SNDBUF_LOCK - sock_sndbuf_lock
	// Locks the send buffer to prevent changes.
	// This is used to lock the size of the send buffer, preventing dynamic adjustments.
	// https://github.com/torvalds/linux/blob/master/net/core/Sock.c
	SockSndbufLock SocketType = 0x1

	// SockTxrehashDefault - SOCK_TXREHASH_DEFAULT - sock_txrehash_default
	// Default setting for Socket transmission rehashing.
	// Determines the default behavior for rehashing the transmission queue.
	// https://github.com/torvalds/linux/blob/master/net/core/Sock.c
	SockTxrehashDefault SocketType = 0xff

	// SockTxrehashDisabled - SOCK_TXREHASH_DISABLED - sock_txrehash_disabled
	// Disable Socket transmission rehashing.
	// Disables the rehashing of the transmission queue to improve performance in certain cases.
	// https://github.com/torvalds/linux/blob/master/net/core/Sock.c
	SockTxrehashDisabled SocketType = 0x0

	// SockTxrehashEnabled - SOCK_TXREHASH_ENABLED - sock_txrehash_enabled
	// Enable Socket transmission rehashing.
	// Allows rehashing of the transmission queue for better load distribution.
	// https://github.com/torvalds/linux/blob/master/net/core/Sock.c
	SockTxrehashEnabled SocketType = 0x1

	// SockRcvbufLock - SOCK_RCVBUF_LOCK - sock_rcvbuf_lock
	// Locks the receive buffer to prevent changes.
	// This is used to lock the size of the receive buffer, preventing dynamic adjustments.
	// https://github.com/torvalds/linux/blob/master/net/core/Sock.c
	SockRcvbufLock SocketType = 0x2
)
