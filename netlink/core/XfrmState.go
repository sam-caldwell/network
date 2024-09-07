package core

// XfrmState - represents various flags used in the XFRM (IPsec) state management.
// These flags control various behaviors of IPsec Security Associations (SAs).
type XfrmState uint8

const (
	// XfrmStateNoEcn - XFRM_STATE_NOECN - disables ECN (Explicit Congestion Notification) for the IPsec state.
	// This is used when ECN is not supported or desired for the Security Association.
	XfrmStateNoEcn XfrmState = 1

	// XfrmStateDecapDscp - XFRM_STATE_DECAP_DSCP - allows the IPsec state to decapsulate DSCP
	// (Differentiated Services Code Point). This flag ensures that DSCP markings are preserved when packets
	// are decapsulated.
	XfrmStateDecapDscp XfrmState = 2

	// XfrmStateNoPmtuDisc - XFRM_STATE_NOPMTUDISC - disables Path MTU Discovery for the IPsec state.
	// When this flag is set, the system will not perform PMTU discovery, which may lead to fragmentation.
	XfrmStateNoPmtuDisc XfrmState = 4

	// XfrmStateWildRecv - XFRM_STATE_WILDRECV - allows the IPsec state to accept packets from any source IP address.
	// This flag is useful when the source address is not known in advance (wildcard reception).
	XfrmStateWildRecv XfrmState = 8

	// XfrmStateIcmp - XFRM_STATE_ICMP - enables handling of ICMP (Internet Control Message Protocol) traffic for
	// the IPsec state. This flag is used when ICMP traffic needs to be managed specifically within the IPsec
	// Security Association.
	XfrmStateIcmp XfrmState = 16

	// XfrmStateAfUnspec - XFRM_STATE_AF_UNSPEC - allows the IPsec state to be unspecified in terms of address family.
	// This flag is used when the IPsec Security Association is not tied to a specific address family (IPv4/IPv6).
	XfrmStateAfUnspec XfrmState = 32

	// XfrmStateAlign4 - XFRM_STATE_ALIGN4 - ensures that the IPsec state aligns data on 4-byte boundaries.
	// This flag is used to optimize data alignment for certain systems or protocols.
	XfrmStateAlign4 XfrmState = 64

	// XfrmStateEsn - XFRM_STATE_ESN - enables Extended Sequence Numbers (ESN) for the IPsec state.
	// ESN provides a larger sequence number space, improving security and avoiding sequence number wraparound.
	XfrmStateEsn XfrmState = 128
)
