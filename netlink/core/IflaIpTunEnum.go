package core

// IflaIpTunEnum -  - Enumeration for IP Tunnel Attributes
//
// This enumeration defines various attributes used for configuring IP tunnels in Linux.
// These attributes are associated with the IP tunnel link and control various aspects of tunnel behavior
// and configuration.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
type IflaIpTunEnum uint8

const (
	// IflaIptunUnspec - IFLA_IPTUN_UNSPEC - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute in the IP tunnel configuration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunUnspec IflaIpTunEnum = 0

	// IflaIptunLink - IFLA_IPTUN_LINK - Network link associated with the tunnel.
	// This attribute specifies the network link (interface) associated with the IP tunnel.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunLink IflaIpTunEnum = 1

	// IflaIptunLocal - IFLA_IPTUN_LOCAL - Local IP address for the tunnel.
	// This attribute specifies the local IP address that is used as the source address for packets sent through
	// the tunnel.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunLocal IflaIpTunEnum = 2

	// IflaIptunRemote - IFLA_IPTUN_REMOTE - Remote IP address for the tunnel.
	// This attribute specifies the remote IP address that the tunnel connects to, acting as the destination address.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunRemote IflaIpTunEnum = 3

	// IflaIptunTtl - IFLA_IPTUN_TTL - Time to Live (TTL) for tunnel packets.
	// This attribute specifies the TTL value for packets sent through the tunnel, controlling how many hops they
	// can pass through.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunTtl IflaIpTunEnum = 4

	// IflaIptunTos - IFLA_IPTUN_TOS - Type of Service (TOS) for tunnel packets.
	// This attribute specifies the TOS field value for packets sent through the tunnel, influencing their priority
	// and handling in the network.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunTos IflaIpTunEnum = 5

	// IflaIptunEncapLimit - IFLA_IPTUN_ENCAP_LIMIT - Encapsulation limit for the tunnel.
	// This attribute specifies the maximum number of nested encapsulations allowed for packets passing through
	// the tunnel.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunEncapLimit IflaIpTunEnum = 6

	// IflaIptunFlowinfo - IFLA_IPTUN_FLOWINFO - Flow information for the tunnel.
	// This attribute specifies additional flow information for packets passing through the tunnel, typically used
	// for QoS or traffic classification.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunFlowinfo IflaIpTunEnum = 7

	// IflaIptunFlags - IFLA_IPTUN_FLAGS - Tunnel flags.
	// This attribute specifies various flags that control the behavior of the IP tunnel, such as whether it should
	// operate in point-to-point mode.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunFlags IflaIpTunEnum = 8

	// IflaIptunProto - IFLA_IPTUN_PROTO - Tunnel protocol.
	// This attribute specifies the protocol used by the IP tunnel, such as GRE or IP-in-IP.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunProto IflaIpTunEnum = 9

	// IflaIptunPmtudisc - IFLA_IPTUN_PMTUDISC - Path MTU Discovery setting.
	// This attribute specifies whether Path MTU Discovery (PMTUD) is enabled for the tunnel, which helps determine
	// the maximum packet size that can be sent.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunPmtudisc IflaIpTunEnum = 10

	// IflaIptun6rdPrefix - IFLA_IPTUN_6RD_PREFIX - 6RD prefix for IPv6 tunnels.
	// This attribute specifies the IPv6 prefix used in 6RD (IPv6 Rapid Deployment) tunnels, which encapsulate IPv6
	// traffic over an IPv4 infrastructure.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptun6rdPrefix IflaIpTunEnum = 11

	// IflaIptun6rdRelayPrefix - IFLA_IPTUN_6RD_RELAY_PREFIX - 6RD relay prefix for IPv6 tunnels.
	// This attribute specifies the relay prefix used in 6RD tunnels, indicating the IPv4 address space that serves
	// as the relay point for IPv6 traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptun6rdRelayPrefix IflaIpTunEnum = 12

	// IflaIptun6rdPrefixlen - IFLA_IPTUN_6RD_PREFIXLEN - 6RD prefix length.
	// This attribute specifies the length of the IPv6 prefix used in 6RD tunnels.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptun6rdPrefixlen IflaIpTunEnum = 13

	// IflaIptun6rdRelayPrefixlen - IFLA_IPTUN_6RD_RELAY_PREFIXLEN - 6RD relay prefix length.
	// This attribute specifies the length of the IPv4 relay prefix used in 6RD tunnels.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptun6rdRelayPrefixlen IflaIpTunEnum = 14

	// IflaIptunEncapType - IFLA_IPTUN_ENCAP_TYPE - Encapsulation type.
	// This attribute specifies the type of encapsulation used by the tunnel, such as direct or GRE encapsulation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunEncapType IflaIpTunEnum = 15

	// IflaIptunEncapFlags - IFLA_IPTUN_ENCAP_FLAGS - Encapsulation flags.
	// This attribute specifies flags related to the encapsulation type, controlling specific behaviors like
	// check-summing.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunEncapFlags IflaIpTunEnum = 16

	// IflaIptunEncapSport - IFLA_IPTUN_ENCAP_SPORT - Encapsulation source port.
	// This attribute specifies the source port used in the encapsulation header of the tunnel packets.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunEncapSport IflaIpTunEnum = 17

	// IflaIptunEncapDport - IFLA_IPTUN_ENCAP_DPORT - Encapsulation destination port.
	// This attribute specifies the destination port used in the encapsulation header of the tunnel packets.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunEncapDport IflaIpTunEnum = 18

	// IflaIptunCollectMetadata - IFLA_IPTUN_COLLECT_METADATA - Collect metadata flag.
	// This attribute specifies whether the tunnel should collect metadata for use in routing decisions or traffic
	// management.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunCollectMetadata IflaIpTunEnum = 19

	// IflaIptunMax - IFLA_IPTUN_MAX - Maximum attribute value.
	// This constant represents the maximum valid value for IP tunnel attributes, used for validation and
	// boundary checks.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaIptunMax = iota - 1
)
