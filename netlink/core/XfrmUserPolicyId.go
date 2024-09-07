package core

// XfrmUserPolicyId - xfrm_userpolicy_id
type XfrmUserPolicyId struct {
	Sel   XfrmSelector
	Index uint32
	Dir   uint8
	Pad   [3]byte
}
