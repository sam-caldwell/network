package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeVfStats - Parses a byte slice into a VfStats struct.
//
// This function deserializes a byte slice containing network Virtual Function (VF) statistics attributes.
// The byte slice is expected to follow the format used by Linux kernel for Virtual Function statistics
// (as defined in the Linux kernel networking subsystem). It extracts and maps individual statistics
// like received/transmitted packets, bytes, and dropped packets, using `ifla_vf_stats` types.
//
// References:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
//   (specifically the `ifla_vf_stats` enumeration and related attributes)
//
// Parameters:
// - b []byte: A byte slice containing serialized VF statistics.
//
// Returns:
// - VfStats: A structure containing parsed VF statistics.
// - error: An error if deserialization fails or the byte slice is too short.
func DeserializeVfStats(b []byte) (VfStats, error) {
	var (
		err      error
		vfstat   VfStats    // Initialize the VfStats structure to hold parsed statistics.
		valueVar uint64     // Temporary variable to hold individual statistic values.
		stats    []NetlinkRouteAttr // Holds parsed netlink route attributes.
	)

	// Ensure that the input byte slice is not nil and has a minimum size to hold VfStats.
	if b == nil || len(b) < SizeOfVfStats {
		return vfstat, errors.New("input too short")
	}

	// Parse the byte slice into an array of netlink route attributes using ParseRouteAttr function.
	// Each attribute corresponds to a specific VF statistic (e.g., RxPackets, TxPackets, etc.).
	if stats, err = ParseRouteAttr(b); err != nil {
		// Return if parsing fails, but return an empty VfStats struct.
		return vfstat, nil
	}

	// Loop through the parsed netlink route attributes and extract VF statistics.
	for _, stat := range stats {
		// Read the statistic value (uint64) from the attribute's value field.
		// Each statistic is stored in a uint64 format in network byte order.
		if err := binary.Read(bytes.NewBuffer(stat.Value), NativeEndian, &valueVar); err != nil {
			break
		}

		// Match the attribute type to the corresponding statistic and assign the value.
		switch IflaVfStatsEnum(stat.Attr.Type) {
		case IflaVfStatsRxPackets:
			// Corresponds to the number of received packets by the VF.
			vfstat.RxPackets = valueVar
		case IflaVfStatsTxPackets:
			// Corresponds to the number of transmitted packets by the VF.
			vfstat.TxPackets = valueVar
		case IflaVfStatsRxBytes:
			// Corresponds to the number of received bytes by the VF.
			vfstat.RxBytes = valueVar
		case IflaVfStatsTxBytes:
			// Corresponds to the number of transmitted bytes by the VF.
			vfstat.TxBytes = valueVar
		case IflaVfStatsMulticast:
			// Corresponds to the number of multicast packets received by the VF.
			vfstat.Multicast = valueVar
		case IflaVfStatsBroadcast:
			// Corresponds to the number of broadcast packets received by the VF.
			vfstat.Broadcast = valueVar
		case IflaVfStatsRxDropped:
			// Corresponds to the number of received packets dropped by the VF.
			vfstat.RxDropped = valueVar
		case IflaVfStatsTxDropped:
			// Corresponds to the number of transmitted packets dropped by the VF.
			vfstat.TxDropped = valueVar
		default:
			// If an unknown attribute type is encountered, return an error.
			return vfstat, errors.New("unknown VfStats type")
		}
	}
