package core

// XfrmMark represents the mark and mask used for packet classification in IPsec.
// This structure corresponds to the xfrm_mark structure in the Linux kernel.
//
// The 'mark' is used to classify packets that match certain criteria, while the 'mask' determines
// which bits of the mark should be considered when matching packets. These fields are used
// in the context of policy-based routing or IPsec Security Associations (SAs) to match traffic
// based on specific conditions.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h

type XfrmMark struct {
	// Value is the mark that is applied to packets to classify them for IPsec policies or SAs.
	// This value is compared to the packet's mark using the mask.
	Value uint32

	// Mask is used to define which bits of the Value should be considered when matching
	// against packet marks. Only the bits set in the mask are used for matching.
	Mask uint32
}
