package core

// SEG6LocalActionString - Converts the provided Seg6LocalActionEnum value into its corresponding string representation.
//
// The function is useful for converting SRv6 local actions into a human-readable format for debugging, logging, or configuration purposes.
// If the action doesn't match any known enum, the function returns "unknown".
//
// Seg6LocalActionEnum values refer to different Segment Routing IPv6 (SRv6) actions.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
func SEG6LocalActionString(action Seg6LocalActionEnum) string {
	switch action {
	case Seg6LocalActionEnd:
		return "End" // End action (node segment).
	case Seg6LocalActionEndX:
		return "End.X" // Adjacency segment (IPv6 cross-connect).
	case Seg6LocalActionEndT:
		return "End.T" // Lookup next segment's NH in the table.
	case Seg6LocalActionEndDx2:
		return "End.DX2" // Decapsulate and L2 cross-connect.
	case Seg6LocalActionEndDx4:
		return "End.DX4" // Decapsulate and IPv4 cross-connect.
	case Seg6LocalActionEndDx6:
		return "End.DX6" // Decapsulate and IPv6 cross-connect.
	case Seg6LocalActionEndDt4:
		return "End.DT4" // Decapsulate and lookup destination address in IPv4 table.
	case Seg6LocalActionEndDt6:
		return "End.DT6" // Decapsulate and lookup destination address in IPv6 table.
	case Seg6LocalActionEndB6:
		return "End.B6" // Binding segment with insertion.
	case Seg6LocalActionEndB6Encap:
		return "End.B6.Encaps" // Binding segment with encapsulation.
	case Seg6LocalActionEndBm:
		return "End.BM" // Binding segment with MPLS encapsulation.
	case Seg6LocalActionEndS:
		return "End.S" // Lookup last segment in table.
	case Seg6LocalActionEndAs:
		return "End.AS" // Forward to SR-unaware VNF with static proxy.
	case Seg6LocalActionEndAm:
		return "End.AM" // Forward to SR-unaware VNF with masquerading.
	case Seg6LocalActionEndBpf:
		return "End.BPF" // Custom BPF action.
	default:
		return "unknown"
	}
}
