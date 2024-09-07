package core

// XfrmSaXflag - represents various flags used in the management of IPsec Security Associations (SAs).
// These flags control specific behaviors of the SAs.
type XfrmSaXflag uint8

const (
	// XfrmSaXflagDontEncapDscp - XFRM_SA_XFLAG_DONT_ENCAP_DSCP - disables the encapsulation of DSCP (Differentiated Services Code Point) for the IPsec SA.
	// This flag ensures that the DSCP values in the original packet are not copied into the encapsulated packet.
	XfrmSaXflagDontEncapDscp XfrmSaXflag = 1

	// XfrmSaXflagOseqMayWrap - XFRM_SA_XFLAG_OSEQ_MAY_WRAP - indicates that the outbound sequence numbers for the IPsec SA may wrap around.
	// This flag is used when the sequence number space is allowed to wrap, usually when Extended Sequence Numbers (ESN) are enabled.
	XfrmSaXflagOseqMayWrap XfrmSaXflag = 2
)
