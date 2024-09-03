package core

// IflaVfStats - Enumeration for Virtual Function (VF) statistics attributes
//
// This enumeration defines various attributes related to network statistics for Virtual Functions (VFs).
// These statistics are used to monitor and manage the performance and activity of VFs in a network interface.

type IflaVfStatsEnum uint8

const (
	// IflaVfStatsRxPackets - IFLA_VF_STATS_RX_PACKETS -
	// This attribute represents the number of packets received by the Virtual Function (VF).
	IflaVfStatsRxPackets IflaVfStatsEnum = 0

	// IflaVfStatsTxPackets - IFLA_VF_STATS_TX_PACKETS -
	// This attribute represents the number of packets transmitted by the Virtual Function (VF).
	IflaVfStatsTxPackets IflaVfStatsEnum = 1

	// IflaVfStatsRxBytes - IFLA_VF_STATS_RX_BYTES -
	// This attribute represents the number of bytes received by the Virtual Function (VF).
	IflaVfStatsRxBytes IflaVfStatsEnum = 2

	// IflaVfStatsTxBytes - IFLA_VF_STATS_TX_BYTES -
	// This attribute represents the number of bytes transmitted by the Virtual Function (VF).
	IflaVfStatsTxBytes IflaVfStatsEnum = 3

	// IflaVfStatsBroadcast - IFLA_VF_STATS_BROADCAST -
	// This attribute represents the number of broadcast packets received by the Virtual Function (VF).
	IflaVfStatsBroadcast IflaVfStatsEnum = 4

	// IflaVfStatsMulticast - IFLA_VF_STATS_MULTICAST -
	// This attribute represents the number of multicast packets received by the Virtual Function (VF).
	IflaVfStatsMulticast IflaVfStatsEnum = 5

	// IflaVfStatsRxDropped - IFLA_VF_STATS_RX_DROPPED -
	// This attribute represents the number of packets dropped by the Virtual Function (VF) upon receipt.
	IflaVfStatsRxDropped IflaVfStatsEnum = 6

	// IflaVfStatsTxDropped - IFLA_VF_STATS_TX_DROPPED -
	// This attribute represents the number of packets dropped by the Virtual Function (VF) upon transmission.
	IflaVfStatsTxDropped IflaVfStatsEnum = 7

	// IflaVfStatsMax - IFLA_VF_STATS_MAX -
	// This constant represents the maximum valid value for VF statistics attributes.
	// It is used as a boundary marker to validate or limit the range of VF statistics attributes.
	IflaVfStatsMax IflaVfStatsEnum = iota - 1
)
