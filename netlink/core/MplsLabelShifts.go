package core

// MplsLabelShifts - Enumerated values for MPLS Label Stack shifts.
//
// These values are used to shift bits in MPLS label stack entries for manipulation
// of the MPLS label, Traffic Class (TC), bottom-of-stack (S) flag, and Time-to-Live (TTL).
//
// Reference: RFC 5462, RFC 3032
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
const (
	// MplsLsLabelShift - MPLS_LS_LABEL_SHIFT - Shift to position the MPLS label (12 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsLabelShift uint32 = 12

	// MplsLsTcShift - MPLS_LS_TC_SHIFT - Shift to position the Traffic Class (9 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsTcShift uint32 = 9

	// MplsLsSShift - MPLS_LS_S_SHIFT - Shift to position the bottom-of-stack (8 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsSShift uint32 = 8

	// MplsLsTtlShift - MPLS_LS_TTL_SHIFT - Shift to position the TTL (0 bits).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/mpls.h
	MplsLsTtlShift uint32 = 0
)
