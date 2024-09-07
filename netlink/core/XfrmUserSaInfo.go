package core

// XfrmUsersaInfo represents the `xfrm_usersa_info` structure in the Linux kernel's XFRM subsystem.
// This structure defines an IPsec Security Association (SA) and includes details about the selector, identifiers,
// lifetime configuration, current lifetime, statistics, and other parameters.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmUsersaInfo struct {
	// Sel represents the XfrmSelector, which is used to define traffic matching rules for the IPsec SA.
	// It specifies the source and destination addresses, ports, protocol, and other traffic-related fields.
	Sel XfrmSelector

	// Id - represents the XfrmId, which contains the Security Parameter Index (SPI), destination address, and protocol.
	// This uniquely identifies the IPsec SA.
	Id XfrmId

	// Saddr - represents the source address for the Security Association (SA).
	// It can be an IPv4 or IPv6 address, depending on the address family.
	Saddr XfrmAddress

	// Lft - represents the lifetime configuration of the SA, defining when the SA should expire based on byte count or time.
	// This field corresponds to the xfrm_lifetime_cfg structure in the Linux kernel.
	Lft XfrmLifetimeCfg

	// Curlft - represents the current lifetime usage of the SA, tracking how many bytes or packets have been sent and how long
	// the SA has been in use.
	Curlft XfrmLifetimeCur

	// Stats - represents various statistics related to the Security Association (SA), such as replay attacks, integrity check failures,
	// and the size of the replay window.
	Stats XfrmStats

	// Seq - represents a sequence number used to identify the SA in the context of netlink messages.
	// It helps track the state and synchronization of the SA.
	Seq uint32

	// Reqid - is a unique identifier used to match policies to Security Associations (SAs).
	// It is particularly useful in scenarios where multiple policies share a single SA.
	Reqid uint32

	// Family - specifies the address family (AF_INET for IPv4, AF_INET6 for IPv6) of the SA.
	// It indicates whether the SA operates on IPv4 or IPv6 traffic.
	Family uint16

	// Mode specifies the mode of the Security Association (SA), either transport mode or tunnel mode.
	// Transport mode secures only the payload, while tunnel mode secures the entire packet.
	Mode uint8

	// ReplayWindow specifies the size of the replay window, which helps prevent replay attacks by keeping track of recently received packets.
	// This field is essential for ensuring that packets are not replayed or duplicated.
	ReplayWindow uint8

	// Flags represents additional flags used to control the behavior of the SA, such as ECN, DSCP, and PMTU settings.
	Flags uint8

	// Pad is a 7-byte padding field used for memory alignment purposes.
	Pad [7]byte
}
