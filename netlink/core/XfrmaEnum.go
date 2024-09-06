package core

// XfrmaEnum - Enumeration for XFRM attributes (xfrm_attrs)
// These attributes are used in Netlink messages related to IPsec and XFRM (Transformations) in the Linux kernel.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmaEnum uint8

const (
	// XfrmaUnspec - XFRMA_UNSPEC - Unspecified attribute, used as a placeholder.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmaUnspec XfrmaEnum = iota

	// XfrmaAlgAuth - XFRMA_ALG_AUTH - Authentication algorithm attribute (struct xfrm_algo).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmaAlgAuth

	// XfrmaAlgCrypt - XFRMA_ALG_CRYPT - Encryption algorithm attribute (struct xfrm_algo).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmaAlgCrypt

	// XfrmaAlgComp - XFRMA_ALG_COMP - Compression algorithm attribute (struct xfrm_algo).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmaAlgComp

	// XfrmaEncap - XFRMA_ENCAP - Encapsulation attribute (struct xfrm_algo + struct xfrm_encap_tmpl).
	XfrmaEncap

	// XfrmaTmpl - XFRMA_TMPL - Template attribute (1 or more struct xfrm_user_tmpl).
	XfrmaTmpl

	// XfrmaSa - XFRMA_SA - Security Association (SA) attribute (struct xfrm_usersa_info).
	XfrmaSa

	// XfrmaPolicy - XFRMA_POLICY - Policy attribute (struct xfrm_userpolicy_info).
	XfrmaPolicy

	// XfrmaSecCtx - XFRMA_SEC_CTX - Security context attribute (struct xfrm_sec_ctx).
	XfrmaSecCtx

	// XfrmaLtimeVal - XFRMA_LTIME_VAL - Lifetime value attribute.
	XfrmaLtimeVal

	// XfrmaReplayVal - XFRMA_REPLAY_VAL - Replay protection value attribute.
	XfrmaReplayVal

	// XfrmaReplayThresh - XFRMA_REPLAY_THRESH - Replay protection threshold attribute.
	XfrmaReplayThresh

	// XfrmaEtimerThresh - XFRMA_ETIMER_THRESH - Event timer threshold attribute.
	XfrmaEtimerThresh

	// XfrmaSrcAddr - XFRMA_SRCADDR - Source address attribute (xfrm_address_t).
	XfrmaSrcAddr

	// XfrmaCoAddr - XFRMA_COADDR - Co-address attribute (xfrm_address_t).
	XfrmaCoAddr

	// XfrmaLastUsed - XFRMA_LASTUSED - Last used time attribute (unsigned long).
	XfrmaLastUsed

	// XfrmaPolicyType - XFRMA_POLICY_TYPE - Policy type attribute (struct xfrm_userpolicy_type).
	XfrmaPolicyType

	// XfrmaMigrate - XFRMA_MIGRATE - Migration attribute.
	XfrmaMigrate

	// XfrmaAlgAead - XFRMA_ALG_AEAD - AEAD algorithm attribute (struct xfrm_algo_aead).
	XfrmaAlgAead

	// XfrmaKmAddress - XFRMA_KMADDRESS - Key management address attribute (struct xfrm_user_kmaddress).
	XfrmaKmAddress

	// XfrmaAlgAuthTrunc - XFRMA_ALG_AUTH_TRUNC - Truncated authentication algorithm attribute (struct xfrm_algo_auth).
	XfrmaAlgAuthTrunc

	// XfrmaMark - XFRMA_MARK - Mark attribute (struct xfrm_mark).
	XfrmaMark

	// XfrmaTfcpad - XFRMA_TFCPAD - Padding attribute (__u32).
	XfrmaTfcpad

	// XfrmaReplayEsnVal - XFRMA_REPLAY_ESN_VAL - Extended Sequence Numbers (ESN) replay value attribute (struct xfrm_replay_esn).
	XfrmaReplayEsnVal

	// XfrmaSaExtraFlags - XFRMA_SA_EXTRA_FLAGS - Extra flags attribute for SA (__u32).
	XfrmaSaExtraFlags

	// XfrmaProto - XFRMA_PROTO - Protocol attribute (__u8).
	XfrmaProto

	// XfrmaAddressFilter - XFRMA_ADDRESS_FILTER - Address filter attribute (struct xfrm_address_filter).
	XfrmaAddressFilter

	// XfrmaPad - XFRMA_PAD - Padding attribute.
	XfrmaPad

	// XfrmaOffloadDev - XFRMA_OFFLOAD_DEV - Offload device attribute (struct xfrm_state_offload).
	XfrmaOffloadDev

	// XfrmaSetMark - XFRMA_SET_MARK - Set mark attribute (__u32).
	XfrmaSetMark

	// XfrmaSetMarkMask - XFRMA_SET_MARK_MASK - Set mark mask attribute (__u32).
	XfrmaSetMarkMask

	// XfrmaIfId - XFRMA_IF_ID - Interface ID attribute (__u32).
	XfrmaIfId

	// XfrmaMax - XFRMA_MAX - Maximum valid attribute type.
	XfrmaMax = iota - 1
)
