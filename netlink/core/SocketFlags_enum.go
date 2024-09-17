package core

type SocketFlags int

const (
	// SockCloExec - SOCK_CLOEXEC - sock_cloexec
	//
	// Purpose:
	//    Sets the close-on-exec (FD_CLOEXEC) flag on the new file descriptor created by the socket system call.
	//
	// Meaning:
	//    When this flag is set, the file descriptor will be automatically closed during the execution of
	//    an `exec` family function (e.g., `execve`).
	//
	// Use Cases:
	//    Useful for avoiding file descriptor leaks in programs that fork and execute other programs, as it ensures
	//    that file descriptors do not remain open unintentionally across exec calls.
	//
	// References:
	// - https://man7.org/linux/man-pages/man2/socket.2.html
	// - https://github.com/torvalds/linux/blob/master/include/linux/net.h
	SockCloExec SocketFlags = 0x80000

	// SockNonBlock - SOCK_NONBLOCK - sock_nonblock
	//
	// Purpose:
	//	  Sets the non-blocking mode for the socket.
	//
	// Meaning:
	//    When this flag is set, socket operations (e.g., read, write) will return immediately, even if they would
	//    normally block. Instead of blocking, they will return an error (EAGAIN or EWOULDBLOCK) if they cannot
	//    proceed immediately.
	//
	// Use Cases:
	//    Used in network applications where non-blocking I/O is required, such as in event-driven or asynchronous
	//    network servers, to avoid blocking the execution flow while waiting for I/O operations to complete.
	//
	// References:
	// - https://man7.org/linux/man-pages/man2/socket.2.html
	// - https://github.com/torvalds/linux/blob/master/include/linux/net.h
	SockNonBlock SocketFlags = 0x800
)
