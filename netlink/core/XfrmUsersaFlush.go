package core

// XfrmUsersaFlush - represents the `xfrm_usersa_flush` structure in the Linux kernel's IPsec subsystem.
// This structure is used when flushing (removing) all Security Associations (SAs) for a specific protocol.
// It indicates which protocol (such as ESP, AH, etc.) to apply the flush operation to.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmUsersaFlush struct {
	// Proto specifies the protocol associated with the Security Associations to be flushed.
	// This could be one of the IPsec protocols, such as ESP (IPPROTO_ESP) or AH (IPPROTO_AH).
	// Example values:
	// - 50: ESP (Encapsulating Security Payload)
	// - 51: AH (Authentication Header)
	Proto uint8
}
