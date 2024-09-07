package core

// XfrmUserPolicyId represents the xfrm_userpolicy_id structure used in the Linux kernel.
// It is used to define the policy identifiers for IPsec policies in the XFRM subsystem.
// The `Sel` field contains the traffic selector (XfrmSelector) to define which traffic the policy applies to.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmUserPolicyId struct {
	// Sel represents the traffic selector for the policy. It defines the source and destination
	// addresses, ports, protocol, and other matching parameters for the IPsec policy.
	// This field corresponds to the xfrm_selector structure in the Linux kernel.
	Sel XfrmSelector

	// Index is the unique index number for the policy, used to identify the specific policy
	// in the kernel's XFRM subsystem. It is a 32-bit unsigned integer.
	Index uint32

	// Dir specifies the direction of traffic the policy applies to.
	// It is a 8-bit unsigned integer that indicates whether the policy is for inbound or outbound traffic.
	// Corresponds to the XFRM_DIR_* constants (e.g., XFRM_DIR_IN, XFRM_DIR_OUT) in the Linux kernel.
	Dir uint8

	// Pad is a 3-byte padding field to ensure proper alignment of the structure in memory.
	// This is commonly required in kernel structures to maintain 4-byte or 8-byte alignment.
	Pad [3]byte
}
