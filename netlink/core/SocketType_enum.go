package core

// SocketType - The socket type - sock_type
//
// Note: we define this as int8 because Linux C sources define this as 8-bits (SOCK_TYPE_MASK is 0xf
//
//	/** sock_type - Socket types
//	 *
//	 * Please notice that for binary compat reasons MIPS has to
//	 * override the enum sock_type in include/linux/net.h, so
//	 * we define ARCH_HAS_SOCKET_TYPES here.
//	 *
//	 * @SOCK_DGRAM - datagram (conn.less) socket
//	 * @SOCK_STREAM - stream (connection) socket
//	 * @SOCK_RAW - raw socket
//	 * @SOCK_RDM - reliably-delivered message
//	 * @SOCK_SEQPACKET - sequential packet socket
//	 * @SOCK_PACKET - linux specific way of getting packets at the dev level.
//	 *		  For writing rarp and other similar things on the user level.
//	 */
//
//	enum sock_type {
//	 	SOCK_DGRAM	= 1,
//	 	SOCK_STREAM	= 2,
//	 	SOCK_RAW	= 3,
//	 	SOCK_RDM	= 4,
//	 	SOCK_SEQPACKET	= 5,
//	 	SOCK_DCCP	= 6,
//	 	SOCK_PACKET	= 10,
//	};
//
// #define SOCK_MAX (SOCK_PACKET + 1)
// /* Mask which covers at least up to SOCK_MASK-1.  The
//   - * remaining bits are used as flags.
//     */
//
// #define SOCK_TYPE_MASK 0xf
type SocketType uint8

const (

	// SockDgram - SOCK_DGRAM - datagram (conn.less) socket
	//
	// https://github.com/torvalds/linux/blob/master/arch/mips/include/asm/socket.h
	SockDgram = 0x1

	// SockStream - SOCK_STREAM - stream (connection) socket
	//
	// https://github.com/torvalds/linux/blob/master/arch/mips/include/asm/socket.h
	SockStream = 0x2

	// SockRaw - SOCK_RAW - sock_raw- raw socket
	// Provides raw network protocol access.
	// Allows direct access to lower-level protocols and the ability to send and receive raw packets.
	//
	// https://en.wikipedia.org/wiki/Raw_socket
	SockRaw SocketType = 0x3

	// SockRdm - SOCK_RDM - sock_rdm - reliably-delivered message
	// Reliable Datagram socket type.
	// Provides a reliable datagram service, where delivery is guaranteed but not in a specific order.
	//
	// https://en.wikipedia.org/wiki/Reliable_User_Datagram_Protocol
	SockRdm SocketType = 0x4

	// SockSeqpacket - SOCK_SEQPACKET - sock_seqpacket- sequential packet socket
	// Sequenced, reliable, two-way connection-based data transmission.
	// Guarantees delivery of data in the order it was sent, and is connection-oriented.
	//
	// https://en.wikipedia.org/wiki/Sequenced_Packet_Exchange
	SockSeqpacket SocketType = 0x5

	// SockDccp - SOCK_DCCP - Datagram Congestion Protocol
	//
	// DCCP is a transport level protocol (like TCP and UDP) which aims to solve many different congestion
	// issues. This is useful for applications that don't need the data reliability/re-transmission of TCP,
	// but want a session and want congestion control unlike UDP.
	//
	// DCCP is currently at proposed standard RFC status (4340-4342).
	//
	// https://en.wikipedia.org/wiki/Sequenced_Packet_Exchange
	// https://wiki.linuxfoundation.org/networking/dccp
	SockDccp = 0x06

	// SockPacket - SOCK_PACKET - sock_packet - linux specific way of getting packets at the dev level.
	// For writing rarp and other similar things on the user level.
	//
	// Raw packet interface that bypasses the network stack.
	// Allows direct sending and receiving of packets at the device driver level.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_packet.h
	SockPacket SocketType = 0xa
)
