package core

// IflaGre -  - Enumeration for GRE (Generic Routing Encapsulation) Tunnel Attributes
//
// This enumeration defines various attributes used for configuring GRE tunnels in Linux.
// GRE tunnels are used to encapsulate a wide variety of network layer protocols inside point-to-point connections.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
type IflaGre uint8

const (
	// IflaGreUnspec - IFLA_GRE_UNSPEC  - Unspecified attribute.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreUnspec IflaGre = 0

	// IflaGreLink - IFLA_GRE_LINK - Network link associated with the GRE tunnel.
	// This attribute specifies the network link (interface) associated with the GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreLink IflaGre = 1

	// IflaGreIflags - IFLA_GRE_IFLAGS - Input flags for the GRE tunnel.
	// This attribute specifies flags related to the input side of the GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreIflags IflaGre = 2

	// IflaGreOflags - IFLA_GRE_OFLAGS - Output flags for the GRE tunnel.
	// This attribute specifies flags related to the output side of the GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreOflags IflaGre = 3

	// IflaGreIkey - IFLA_GRE_IKEY - Input key for the GRE tunnel.
	// This attribute specifies the key used for identifying incoming packets in the GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreIkey IflaGre = 4

	// IflaGreOkey - IFLA_GRE_OKEY - Output key for the GRE tunnel.
	// This attribute specifies the key used for identifying outgoing packets in the GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreOkey IflaGre = 5

	// IflaGreLocal - IFLA_GRE_LOCAL - Local IP address for the GRE tunnel.
	// This attribute specifies the local IP address that is used as the source address for packets sent
	// through the GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreLocal IflaGre = 6

	// IflaGreRemote - IFLA_GRE_REMOTE - Remote IP address for the GRE tunnel.
	// This attribute specifies the remote IP address that the GRE tunnel connects to, acting as the destination
	// address.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreRemote IflaGre = 7

	// IflaGreTtl - IFLA_GRE_TTL - Time to Live (TTL) for GRE tunnel packets.
	// This attribute specifies the TTL value for packets sent through the GRE tunnel, controlling how many hops they
	// can pass through.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreTtl IflaGre = 8

	// IflaGreTos - IFLA_GRE_TOS - Type of Service (TOS) for GRE tunnel packets.
	// This attribute specifies the TOS field value for packets sent through the GRE tunnel, influencing their
	// priority and handling in the network.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreTos IflaGre = 9

	// IflaGrePmtudisc - IFLA_GRE_PMTUDISC - Path MTU Discovery setting.
	// This attribute specifies whether Path MTU Discovery (PMTUD) is enabled for the GRE tunnel, helping determine
	// the maximum packet size that can be sent.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGrePmtudisc IflaGre = 10

	// IflaGreEncapLimit - IFLA_GRE_ENCAP_LIMIT - Encapsulation limit for the GRE tunnel.
	// This attribute specifies the maximum number of nested encapsulations allowed for packets passing through the
	// GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreEncapLimit IflaGre = 11

	// IflaGreFlowinfo - IFLA_GRE_FLOWINFO - Flow information for the GRE tunnel.
	// This attribute specifies additional flow information for packets passing through the GRE tunnel, typically
	// used for QoS or traffic classification.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreFlowinfo IflaGre = 12

	// IflaGreFlags - IFLA_GRE_FLAGS - GRE tunnel flags.
	// This attribute specifies various flags that control the behavior of the GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreFlags IflaGre = 13

	// IflaGreEncapType - IFLA_GRE_ENCAP_TYPE - Encapsulation type.
	// This attribute specifies the type of encapsulation used by the GRE tunnel.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreEncapType IflaGre = 14

	// IflaGreEncapFlags - IFLA_GRE_ENCAP_FLAGS - Encapsulation flags.
	// This attribute specifies flags related to the encapsulation type, controlling specific behaviors like
	// check-summing.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreEncapFlags IflaGre = 15

	// IflaGreEncapSport - IFLA_GRE_ENCAP_SPORT - Encapsulation source port.
	// This attribute specifies the source port used in the encapsulation header of the GRE tunnel packets.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreEncapSport IflaGre = 16

	// IflaGreEncapDport - IFLA_GRE_ENCAP_DPORT - Encapsulation destination port.
	// This attribute specifies the destination port used in the encapsulation header of the GRE tunnel packets.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreEncapDport IflaGre = 17

	// IflaGreCollectMetadata - IFLA_GRE_COLLECT_METADATA - Collect metadata flag.
	// This attribute specifies whether the GRE tunnel should collect metadata for use in routing decisions or
	// traffic management.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreCollectMetadata IflaGre = 18

	// IflaGreMax - IFLA_GRE_MAX - Maximum attribute value.
	// This constant represents the maximum valid value for GRE tunnel attributes, used for validation and
	// boundary checks.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGreMax = iota - 1
)
