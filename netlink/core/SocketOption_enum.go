//go:build linux

package core

// SocketOption - Used for getting and setting options at the socket level.
//
// References:
// - https://github.com/torvalds/linux/blob/master/include/uapi/asm-generic/socket.h
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h

type SocketOption int

const (
	// SolSocket - SOL_SOCKET - Socket level option.
	// Used for getting and setting options at the socket level.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SolSocket SocketOption = 0x1

	// SoAcceptconn - SO_ACCEPTCONN - Check if the socket is in listening mode.
	// It returns whether a socket has been marked to accept incoming connections.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoAcceptconn SocketOption = 0x1e

	// SoAttachBpf - SO_ATTACH_BPF - Attach an extended BPF (eBPF) program to the socket.
	// This is used for attaching a BPF program to a socket for packet filtering.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoAttachBpf SocketOption = 0x32

	// SoAttachReuseportCbpf - SO_ATTACH_REUSEPORT_CBPF - Attach a classic BPF program for reuse port selection.
	// Allows attaching a BPF program to control which socket in a reuseport group should handle the packet.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoAttachReuseportCbpf SocketOption = 0x33

	// SoAttachReuseportEbpf - SO_ATTACH_REUSEPORT_EBPF - Attach an eBPF program for reuse port selection.
	// Similar to SO_ATTACH_REUSEPORT_CBPF but uses the extended BPF.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoAttachReuseportEbpf SocketOption = 0x34

	// SoBindtodevice - SO_BINDTODEVICE - Bind the socket to a specific network device.
	// This forces the socket to send and receive packets only through a specified network interface.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoBindtodevice SocketOption = 0x19

	// SoBindtoifindex - SO_BINDTOIFINDEX - Bind the socket to a specific network interface by index.
	// Similar to SO_BINDTODEVICE, but uses the interface index instead of name.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoBindtoifindex SocketOption = 0x3e

	// SoBpfExtensions - SO_BPF_EXTENSIONS - Check for BPF extensions supported by the kernel.
	// It provides information about additional features supported by the BPF infrastructure.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoBpfExtensions SocketOption = 0x30

	// SoBroadcast - SO_BROADCAST - Allow transmission of broadcast messages.
	// Enables the ability to send broadcast messages on the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoBroadcast SocketOption = 0x6

	// SoBsdcompat - SO_BSDCOMPAT - Enable BSD-style socket compatibility.
	// This option is used to maintain compatibility with older BSD socket behavior.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoBsdcompat SocketOption = 0xe

	// SoBufLock - SO_BUF_LOCK - Lock the socket buffer to prevent changes.
	// This option is used to lock the socket buffer, preventing it from being modified.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoBufLock SocketOption = 0x48

	// SoBusyPoll - SO_BUSY_POLL - Enable busy polling on the socket.
	// This option enables the socket to use busy polling for incoming packets, reducing latency.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoBusyPoll SocketOption = 0x2e

	// SoBusyPollBudget - SO_BUSY_POLL_BUDGET - Set the budget for busy polling.
	// This defines the maximum number of packets the socket will process during busy polling.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoBusyPollBudget SocketOption = 0x46

	// SoCnxAdvice - SO_CNX_ADVICE - Provide connection advice to the socket.
	// Used to give hints to the kernel about the expected behavior of the connection.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoCnxAdvice SocketOption = 0x35

	// SoCookie - SO_COOKIE - Get the socket's cookie value.
	// Provides a unique identifier for the socket, often used for debugging purposes.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoCookie SocketOption = 0x39

	// SoDetachReuseportBpf - SO_DETACH_REUSEPORT_BPF - Detach a BPF program from a reuse port socket.
	// This option allows removing an attached BPF program from a reuse port socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoDetachReuseportBpf SocketOption = 0x44

	// SoDomain - SO_DOMAIN - Get the socket domain (address family).
	// Returns the domain of the socket, such as AF_INET or AF_INET6.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoDomain SocketOption = 0x27

	// SoDontroute - SO_DONTROUTE - Bypass the routing table for outgoing packets.
	// Sends packets directly to the interface without routing them, often used for diagnostics.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoDontroute SocketOption = 0x5

	// SoError - SO_ERROR - Get the error status and clear it.
	// Used to retrieve and clear the pending error on the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoError SocketOption = 0x4

	// SoIncomingCpu - SO_INCOMING_CPU - Get the CPU number where the socket is receiving packets.
	// This option retrieves the CPU number handling the socket's incoming packets.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoIncomingCpu SocketOption = 0x31

	// SoIncomingNapiId - SO_INCOMING_NAPI_ID - Get the NAPI ID where the socket is receiving packets.
	// Retrieves the NAPI ID handling the socket's incoming packets, useful for network performance tuning.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoIncomingNapiId SocketOption = 0x38

	// SoKeepalive - SO_KEEPALIVE - Enable sending keepalive messages on the socket.
	// This option enables periodic transmission of keepalive messages to maintain the connection.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoKeepalive SocketOption = 0x9

	// SoLinger - SO_LINGER - Set the linger option on the socket.
	// Defines the behavior of the socket when closed with unsent data.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoLinger SocketOption = 0xd

	// SoLockFilter - SO_LOCK_FILTER - Lock the attached BPF filter on the socket.
	// Prevents the attached BPF filter from being replaced or removed.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoLockFilter SocketOption = 0x2c

	// SoMark - SO_MARK - Set a mark value on packets sent through the socket.
	// Marks outgoing packets with a value for routing or firewalling decisions.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoMark SocketOption = 0x24

	// SoMaxPacingRate - SO_MAX_PACING_RATE - Set the maximum pacing rate for the socket.
	// Limits the rate at which packets are sent to prevent network congestion.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoMaxPacingRate SocketOption = 0x2f

	// SoMeminfo - SO_MEMINFO - Get memory usage information for the socket.
	// Retrieves statistics about the socket's memory usage for receive and send buffers.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoMeminfo SocketOption = 0x37

	// SoNetnsCookie - SO_NETNS_COOKIE - Get the network namespace cookie for the socket.
	// Provides an identifier for the network namespace in which the socket operates.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoNetnsCookie SocketOption = 0x47

	// SoNofcs - SO_NOFCS - Disable generation of FCS (Frame Check Sequence).
	// Used to disable the generation of FCS for outgoing packets.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoNofcs SocketOption = 0x2b

	// SoOobinline - SO_OOBINLINE - Leave out-of-band data inline.
	// This option allows receiving out-of-band data in the normal data stream.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoOobinline SocketOption = 0xa

	// SoPasscred - SO_PASSCRED - Receive the credentials of the sending process.
	// Allows the socket to receive the credentials (UID, GID, PID) of the sending process.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPasscred SocketOption = 0x10

	// SoPasspidfd - SO_PASSPIDFD - Receive the PID file descriptor of the sending process.
	// Enables receiving the file descriptor associated with the sending process's PID.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPasspidfd SocketOption = 0x4c

	// SoPasssec - SO_PASSSEC - Receive the security context of the sending process.
	// Allows the socket to receive security context information of the sender.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPasssec SocketOption = 0x22

	// SoPeekOff - SO_PEEK_OFF - Set an offset for peeking into the receive queue.
	// This option allows peeking into the receive queue starting at a specified offset.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPeekOff SocketOption = 0x2a

	// SoPeercred - SO_PEERCRED - Get the credentials of the peer process.
	// Retrieves the credentials (UID, GID, PID) of the process on the other side of the connection.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPeercred SocketOption = 0x11

	// SoPeergroups - SO_PEERGROUPS - Get the peer's supplementary groups.
	// Returns the supplementary groups of the peer process.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPeergroups SocketOption = 0x3b

	// SoPeerpidfd - SO_PEERPIDFD - Get the peer's PID file descriptor.
	// Provides a file descriptor for the peer process's PID.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPeerpidfd SocketOption = 0x4d

	// SoPeersec - SO_PEERSEC - Get the security context of the peer process.
	// Retrieves the security context of the process on the other side of the connection.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPeersec SocketOption = 0x1f

	// SoPreferBusyPoll - SO_PREFER_BUSY_POLL - Indicate preference for busy polling.
	// Suggests to the kernel that busy polling is preferred for this socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoPreferBusyPoll SocketOption = 0x45

	// SoProtocol - SO_PROTOCOL - Get the protocol type of the socket.
	// Returns the protocol type of the socket (e.g., TCP, UDP).
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoProtocol SocketOption = 0x26

	// SoRcvbuf - SO_RCVBUF - Set or get the size of the receive buffer.
	// Defines the size of the socket's receive buffer to control flow.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoRcvbuf SocketOption = 0x8

	// SoRcvbufforce - SO_RCVBUFFORCE - Force the size of the receive buffer.
	// Allows setting the receive buffer size beyond the system limit.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoRcvbufforce SocketOption = 0x21

	// SoRcvlowat - SO_RCVLOWAT - Set or get the minimum number of bytes to process for input.
	// Defines the minimum amount of data in the receive buffer that triggers a read.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoRcvlowat SocketOption = 0x12

	// SoRcvmark - SO_RCVMARK - Get the mark value for the last received packet.
	// Retrieves the mark associated with the last received packet on the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoRcvmark SocketOption = 0x4b

	// SoRcvtimeo - SO_RCVTIMEO - Set or get the receive timeout.
	// Defines the timeout value for receiving data on the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoRcvtimeo SocketOption = 0x14

	// SoRcvtimeoNew - SO_RCVTIMEO_NEW - New receive timeout interface with nanosecond precision.
	// An updated version of SO_RCVTIMEO that provides nanosecond precision for the timeout.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoRcvtimeoNew SocketOption = 0x42

	// SoRcvtimeoOld - SO_RCVTIMEO_OLD - Old receive timeout interface.
	// The original SO_RCVTIMEO, supporting only microsecond precision.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoRcvtimeoOld SocketOption = 0x14

	// SoReserveMem - SO_RESERVE_MEM - Reserve memory for the socket.
	// Allows reserving memory for socket operations, typically for specific protocols.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoReserveMem SocketOption = 0x49

	// SoReuseaddr - SO_REUSEADDR - Allow reuse of local addresses.
	// Allows binding a socket to an address that is already in use.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoReuseaddr SocketOption = 0x2

	// SoReuseport - SO_REUSEPORT - Allow multiple sockets to bind to the same port.
	// Enables multiple sockets to bind to the same port and receive datagrams.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoReuseport SocketOption = 0xf

	// SoRxqOvfl - SO_RXQ_OVFL - Enable receive queue overflow notification.
	// Notifies the application when packets are dropped due to receive queue overflow.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoRxqOvfl SocketOption = 0x28

	// SoSecurityAuthentication - SO_SECURITY_AUTHENTICATION - Security authentication option.
	// Placeholder for security authentication settings on the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSecurityAuthentication SocketOption = 0x16

	// SoSecurityEncryptionNetwork - SO_SECURITY_ENCRYPTION_NETWORK - Security encryption for network layer.
	// Enables or configures network layer encryption for data sent through the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSecurityEncryptionNetwork SocketOption = 0x18

	// SoSecurityEncryptionTransport - SO_SECURITY_ENCRYPTION_TRANSPORT - Security encryption for transport layer.
	// Enables or configures transport layer encryption for data sent through the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSecurityEncryptionTransport SocketOption = 0x17

	// SoSelectErrQueue - SO_SELECT_ERR_QUEUE - Select error queue for the socket.
	// Used to select the error queue for monitoring and processing socket errors.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSelectErrQueue SocketOption = 0x2d

	// SoSndbuf - SO_SNDBUF - Set or get the size of the send buffer.
	// Defines the size of the socket's send buffer for controlling flow.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSndbuf SocketOption = 0x7

	// SoSndbufforce - SO_SNDBUFFORCE - Force the size of the send buffer.
	// Allows setting the send buffer size beyond the system limit.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSndbufforce SocketOption = 0x20

	// SoSndlowat - SO_SNDLOWAT - Set or get the minimum number of bytes to process for output.
	// Defines the minimum amount of data in the send buffer that triggers a write.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSndlowat SocketOption = 0x13

	// SoSndtimeo - SO_SNDTIMEO - Set or get the send timeout.
	// Defines the timeout value for sending data on the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSndtimeo SocketOption = 0x15

	// SoSndtimeoNew - SO_SNDTIMEO_NEW - New send timeout interface with nanosecond precision.
	// An updated version of SO_SNDTIMEO that provides nanosecond precision for the timeout.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSndtimeoNew SocketOption = 0x43

	// SoSndtimeoOld - SO_SNDTIMEO_OLD - Old send timeout interface.
	// The original SO_SNDTIMEO, supporting only microsecond precision.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoSndtimeoOld SocketOption = 0x15

	// SoTimestamping - SO_TIMESTAMPING - Enable hardware and software timestamping.
	// Allows the socket to collect timestamps for sent and received packets.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTimestamping SocketOption = 0x25

	// SoTimestampingNew - SO_TIMESTAMPING_NEW - New timestamping interface with enhanced features.
	// An updated version of SO_TIMESTAMPING that supports additional timestamping capabilities.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTimestampingNew SocketOption = 0x41

	// SoTimestampingOld - SO_TIMESTAMPING_OLD - Old timestamping interface.
	// The original SO_TIMESTAMPING, providing basic timestamping capabilities.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTimestampingOld SocketOption = 0x25

	// SoTimestampns - SO_TIMESTAMPNS - Enable nanosecond precision timestamping.
	// Allows the socket to collect timestamps with nanosecond precision.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTimestampns SocketOption = 0x23

	// SoTimestampnsNew - SO_TIMESTAMPNS_NEW - New nanosecond precision timestamping interface.
	// An updated version of SO_TIMESTAMPNS that provides enhanced nanosecond precision timestamping.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTimestampnsNew SocketOption = 0x40

	// SoTimestampnsOld - SO_TIMESTAMPNS_OLD - Old nanosecond precision timestamping interface.
	// The original SO_TIMESTAMPNS, providing basic nanosecond precision timestamping.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTimestampnsOld SocketOption = 0x23

	// SoTimestampNew - SO_TIMESTAMP_NEW - New timestamp interface.
	// An updated version of SO_TIMESTAMP that provides additional features.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTimestampNew SocketOption = 0x3f

	// SoTxrehash - SO_TXREHASH - Rehash the socket's transmission queue.
	// Used to rehash the transmission queue to improve performance.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTxrehash SocketOption = 0x4a

	// SoTxtime - SO_TXTIME - Set a specific transmission time for packets.
	// This option allows specifying a transmission time for outgoing packets.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoTxtime SocketOption = 0x3d

	// SoType - SO_TYPE - Get the socket type.
	// Returns the type of the socket, such as SOCK_STREAM or SOCK_DGRAM.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoType SocketOption = 0x3

	// SoWifiStatus - SO_WIFI_STATUS - Get WiFi status information.
	// Provides information about the WiFi status associated with the socket.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoWifiStatus SocketOption = 0x29

	// SoZerocopy - SO_ZEROCOPY - Enable zero-copy operation for the socket.
	// Allows the socket to perform zero-copy data transfers, improving performance.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/socket.h
	SoZerocopy SocketOption = 0x3c
)
