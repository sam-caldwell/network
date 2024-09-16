package core

// SocketType - The socket type
type SocketType int

const (

	// sockPacket - SOCK_PACKET - sock_packet
	// Raw packet interface that bypasses the network stack.
	// Allows direct sending and receiving of packets at the device driver level.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_packet.h
	sockPacket SocketType = 0xa

	// sockRaw - SOCK_RAW - sock_raw
	// Provides raw network protocol access.
	// Allows direct access to lower-level protocols and the ability to send and receive raw packets.
	// https://en.wikipedia.org/wiki/Raw_socket
	sockRaw SocketType = 0x3

	// sockRdm - SOCK_RDM - sock_rdm
	// Reliable Datagram socket type.
	// Provides a reliable datagram service, where delivery is guaranteed but not in a specific order.
	// https://en.wikipedia.org/wiki/Reliable_User_Datagram_Protocol
	sockRdm SocketType = 0x4

	// sockSeqpacket - SOCK_SEQPACKET - sock_seqpacket
	// Sequenced, reliable, two-way connection-based data transmission.
	// Guarantees delivery of data in the order it was sent, and is connection-oriented.
	// https://en.wikipedia.org/wiki/Sequenced_Packet_Exchange
	sockSeqpacket SocketType = 0x5
)
