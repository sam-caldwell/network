package core

// SocketCommand - commands used to manage sockets
type SocketCommand int

const (

	// sockDiagByFamily - SOCK_DIAG_BY_FAMILY - sock_diag_by_family
	// Socket diagnostics by family.
	// Provides a mechanism to get detailed information about sockets for a specific protocol family.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/sock_diag.h
	sockDiagByFamily SocketType = 0x14

	// sockDestroy - SOCK_DESTROY - sock_destroy
	// Constant used internally to destroy a socket.
	// This is typically used for socket cleanup operations in the kernel.
	// https://github.com/torvalds/linux/blob/master/net/core/sock.c
	sockDestroy SocketCommand = 0x15

	// sockIocType - SOCK_IOC_TYPE - sock_ioc_type
	// IOCTL command type for sockets.
	// Defines the type of IOCTL (Input/Output Control) command that can be issued on a socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/sockios.h
	sockIocType SocketType = 0x89

	// sockSndbufLock - SOCK_SNDBUF_LOCK - sock_sndbuf_lock
	// Locks the send buffer to prevent changes.
	// This is used to lock the size of the send buffer, preventing dynamic adjustments.
	// https://github.com/torvalds/linux/blob/master/net/core/sock.c
	sockSndbufLock SocketType = 0x1

	// sockTxrehashDefault - SOCK_TXREHASH_DEFAULT - sock_txrehash_default
	// Default setting for socket transmission rehashing.
	// Determines the default behavior for rehashing the transmission queue.
	// https://github.com/torvalds/linux/blob/master/net/core/sock.c
	sockTxrehashDefault SocketType = 0xff

	// sockTxrehashDisabled - SOCK_TXREHASH_DISABLED - sock_txrehash_disabled
	// Disable socket transmission rehashing.
	// Disables the rehashing of the transmission queue to improve performance in certain cases.
	// https://github.com/torvalds/linux/blob/master/net/core/sock.c
	sockTxrehashDisabled SocketType = 0x0

	// sockTxrehashEnabled - SOCK_TXREHASH_ENABLED - sock_txrehash_enabled
	// Enable socket transmission rehashing.
	// Allows rehashing of the transmission queue for better load distribution.
	// https://github.com/torvalds/linux/blob/master/net/core/sock.c
	sockTxrehashEnabled SocketType = 0x1

	// sockRcvbufLock - SOCK_RCVBUF_LOCK - sock_rcvbuf_lock
	// Locks the receive buffer to prevent changes.
	// This is used to lock the size of the receive buffer, preventing dynamic adjustments.
	// https://github.com/torvalds/linux/blob/master/net/core/sock.c
	sockRcvbufLock SocketType = 0x2
)
