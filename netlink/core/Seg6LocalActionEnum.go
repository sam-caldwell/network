package core

// Seg6LocalActionEnum - Segment Routing IPv6 (SRv6) Local Actions.
//
// This enumerates the available local actions for SRv6, defining how packets are processed at SRv6 endpoints.
// Each local action corresponds to a specific operation such as decapsulation, cross-connect, or table lookup.
//
// For more details, see the Linux kernel source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
type Seg6LocalActionEnum uint8

const (
	// Seg6LocalActionUnspec - SEG6_LOCAL_ACTION_UNSPEC - Unspecified action (0).
	// This is a placeholder for an unspecified action.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionUnspec Seg6LocalActionEnum = 0

	// Seg6LocalActionEnd - SEG6_LOCAL_ACTION_END - Node segment (1).
	// The SRv6 packet reaches its final destination at this node.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEnd Seg6LocalActionEnum = 1

	// Seg6LocalActionEndX - SEG6_LOCAL_ACTION_END_X - Adjacency segment (IPv6 cross-connect) (2).
	// The packet is forwarded to the next hop based on IPv6 cross-connect.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndX Seg6LocalActionEnum = 2

	// Seg6LocalActionEndT - SEG6_LOCAL_ACTION_END_T - Lookup next segment in routing table (3).
	// The next hop is determined by performing a lookup in the routing table.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndT Seg6LocalActionEnum = 3

	// Seg6LocalActionEndDx2 - SEG6_LOCAL_ACTION_END_DX2 - Decap and L2 cross-connect (4).
	// The packet is decapsulated and forwarded based on L2 cross-connect.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndDx2 Seg6LocalActionEnum = 4

	// Seg6LocalActionEndDx6 - SEG6_LOCAL_ACTION_END_DX6 - Decap and IPv6 cross-connect (5).
	// The packet is decapsulated and forwarded based on IPv6 cross-connect.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndDx6 Seg6LocalActionEnum = 5

	// Seg6LocalActionEndDx4 - SEG6_LOCAL_ACTION_END_DX4 - Decap and IPv4 cross-connect (6).
	// The packet is decapsulated and forwarded based on IPv4 cross-connect.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndDx4 Seg6LocalActionEnum = 6

	// Seg6LocalActionEndDt6 - SEG6_LOCAL_ACTION_END_DT6 - Decap and lookup of destination address in IPv6 table (7).
	// The packet is decapsulated, and the next hop is determined by looking up the destination address in the IPv6 routing table.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndDt6 Seg6LocalActionEnum = 7

	// Seg6LocalActionEndDt4 - SEG6_LOCAL_ACTION_END_DT4 - Decap and lookup of destination address in IPv4 table (8).
	// The packet is decapsulated, and the next hop is determined by looking up the destination address in the IPv4 routing table.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndDt4 Seg6LocalActionEnum = 8

	// Seg6LocalActionEndB6 - SEG6_LOCAL_ACTION_END_B6 - Binding segment with insertion (9).
	// The packet is processed as a binding segment with segment list insertion.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndB6 Seg6LocalActionEnum = 9

	// Seg6LocalActionEndB6Encap - SEG6_LOCAL_ACTION_END_B6_ENCAP - Binding segment with encapsulation (10).
	// The packet is processed as a binding segment with encapsulation.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndB6Encap Seg6LocalActionEnum = 10

	// Seg6LocalActionEndBm - SEG6_LOCAL_ACTION_END_BM - Binding segment with MPLS encapsulation (11).
	// The packet is processed as a binding segment with MPLS encapsulation.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndBm Seg6LocalActionEnum = 11

	// Seg6LocalActionEndS - SEG6_LOCAL_ACTION_END_S - Lookup last segment in the table (12).
	// The packet is processed based on a lookup of the last segment in the routing table.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndS Seg6LocalActionEnum = 12

	// Seg6LocalActionEndAs - SEG6_LOCAL_ACTION_END_AS - Forward to SR-unaware VNF with static proxy (13).
	// The packet is forwarded to a service function (VNF) that is unaware of SRv6 using a static proxy.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndAs Seg6LocalActionEnum = 13

	// Seg6LocalActionEndAm - SEG6_LOCAL_ACTION_END_AM - Forward to SR-unaware VNF with masquerading (14).
	// The packet is forwarded to a service function (VNF) that is unaware of SRv6 using masquerading.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndAm Seg6LocalActionEnum = 14

	// Seg6LocalActionEndBpf - SEG6_LOCAL_ACTION_END_BPF - Custom BPF action (15).
	// A custom action is applied to the packet using a BPF (Berkeley Packet Filter) program.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndBpf Seg6LocalActionEnum = 15

	// Seg6LocalActionEndDt46 - SEG6_LOCAL_ACTION_END_DT46 - Decap and lookup in either IPv4 or IPv6 table (16).
	// The packet is decapsulated, and the next hop is determined by looking up the destination address in either the IPv4 or IPv6 routing table.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6_local.h
	Seg6LocalActionEndDt46 Seg6LocalActionEnum = iota - 1
)
