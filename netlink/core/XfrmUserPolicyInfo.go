package core

// XfrmUserPolicyInfo represents the xfrm_userpolicy_info structure used in the Linux kernel's XFRM subsystem.
// It defines the configuration for an IPsec policy, including traffic selectors, lifetimes, priority, and other attributes.
//
// This structure corresponds to the `xfrm_userpolicy_info` structure in the Linux kernel.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmUserPolicyInfo struct {
	// Sel represents the traffic selector for the policy. It defines which traffic (source/destination addresses, ports, etc.)
	// the policy applies to. This field corresponds to the `xfrm_selector` structure.
	Sel XfrmSelector

	// Lft represents the lifetime configuration for the policy. It defines the soft and hard limits for bytes, packets, and time.
	// This field corresponds to the `xfrm_lifetime_cfg` structure.
	Lft XfrmLifetimeCfg

	// Curlft represents the current lifetime usage for the policy, including the number of bytes and packets processed.
	// This field corresponds to the `xfrm_lifetime_cur` structure.
	Curlft XfrmLifetimeCur

	// Priority defines the priority of the policy. Lower values represent higher priority.
	Priority uint32

	// Index is the unique index of the policy, used to identify it in the kernel's XFRM subsystem.
	Index uint32

	// Dir specifies the direction of traffic that the policy applies to (e.g., inbound or outbound).
	// This corresponds to the XFRM_DIR_* constants in the Linux kernel (e.g., XFRM_DIR_IN, XFRM_DIR_OUT).
	Dir uint8

	// Action specifies the action to be taken on matching traffic (e.g., allow or block).
	// This corresponds to the XFRM_POLICY_* constants in the Linux kernel (e.g., XFRM_POLICY_ALLOW, XFRM_POLICY_BLOCK).
	Action uint8

	// Flags contains additional flags for the policy configuration, such as whether the policy is disabled.
	Flags uint8

	// Share indicates the sharing mode for the policy, controlling how the policy can be used by multiple processes.
	Share uint8

	// Pad is a 4-byte padding field for alignment purposes.
	Pad [4]byte
}
