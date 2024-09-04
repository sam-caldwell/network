package core

// MplsLabelStackShiftEnum - Enumerated values for MPLS Label Stack shifts.
//
// These values are used to shift bits in MPLS label stack entries for manipulation
// of the MPLS label, Traffic Class (TC), bottom-of-stack (S) flag, and Time-to-Live (TTL).
//
// Reference: RFC 5462, RFC 3032
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
type MplsLabelStackShiftEnum uint8

// MplsLabelStackShiftMask - Defines masks for MPLS Label Stack fields.
//
// These masks are used to isolate fields in MPLS label stack entries for
// operations such as shifting, setting, or clearing bits.
//
// Reference: RFC 5462, RFC 3032
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
type MplsLabelStackShiftMask uint32

const (

	// MplsLsLabelMask - MPLS_LS_LABEL_MASK - Mask to extract the MPLS label field (20 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsLabelMask MplsLabelStackShiftMask = 0xFFFFF000

	// MplsLsLabelShift - MPLS_LS_LABEL_SHIFT - Shift to position the MPLS label (12 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsLabelShift MplsLabelStackShiftEnum = 12

	// MplsLsTcMask - MPLS_LS_TC_MASK - Mask to extract the Traffic Class (TC) field (3 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsTcMask MplsLabelStackShiftMask = 0x00000E00

	// MplsLsTcShift - MPLS_LS_TC_SHIFT - Shift to position the Traffic Class (9 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsTcShift MplsLabelStackShiftEnum = 9

	// MplsLsSMask - MPLS_LS_S_MASK - Mask to extract the bottom-of-stack (S) bit.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsSMask MplsLabelStackShiftMask = 0x00000100

	// MplsLsSShift - MPLS_LS_S_SHIFT - Shift to position the bottom-of-stack (8 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsSShift MplsLabelStackShiftEnum = 8

	// MplsLsTtlMask - MPLS_LS_TTL_MASK - Mask to extract the Time-to-Live (TTL) field (8 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsTtlMask MplsLabelStackShiftMask = 0x000000FF

	// MplsLsTtlShift - MPLS_LS_TTL_SHIFT - Shift to position the TTL (0 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsTtlShift MplsLabelStackShiftEnum = 0
)
