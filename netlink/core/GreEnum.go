package core

// GreEnum - Enumeration for GRE (Generic Routing Encapsulation) flags and fields
//
// This enumeration defines various flags and fields related to GRE (Generic Routing Encapsulation).
// GRE is a tunneling protocol used to encapsulate a wide variety of network layer protocols inside virtual
// point-to-point connections.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
type GreEnum uint16

const (
	// GreCsum - GRE_CSUM -
	// This flag indicates that a checksum is present in the GRE header.
	// The checksum is used to verify the integrity of the encapsulated data.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreCsum GreEnum = 0x8000

	// GreRouting - GRE_ROUTING -
	// This flag indicates the presence of routing information in the GRE packet.
	// Routing information is used for source routing, where the packet carries information about the path it
	// should take.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreRouting GreEnum = 0x4000

	// GreKey - GRE_KEY -
	// This flag indicates that a key is present in the GRE header.
	// The key is typically used to identify a particular flow or connection in a GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreKey GreEnum = 0x2000

	// GreSeq - GRE_SEQ -
	// This flag indicates that a sequence number is present in the GRE header.
	// The sequence number is used to ensure packet order and detect duplicates.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreSeq GreEnum = 0x1000

	// GreStrict - GRE_STRICT -
	// This flag indicates the use of strict source routing.
	// In strict source routing, the packet must follow a precise path as specified in the routing information.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreStrict GreEnum = 0x0800

	// GreRec - GRE_REC -
	// This field specifies the recursion control in the GRE header.
	// It is used to control the number of times a GRE packet can be encapsulated.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreRec GreEnum = 0x0700

	// GreAck - GRE_ACK -
	// This field contains reserved flags in the GRE header.
	// These flags are reserved for future use or for specific protocol extensions.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreAck GreEnum = 0x0080

	// GreFlags - GRE_FLAGS -
	// This field contains reserved flags in the GRE header.
	// These flags are reserved for future use or for specific protocol extensions.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreFlags GreEnum = 0x00F8

	// GreVersion - GRE_VERSION -
	// This field specifies the GRE protocol version.
	// It is used to identify the version of the GRE protocol being used in the packet.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	GreVersion GreEnum = 0x0007
)
