//go:build linux

package core

import (
	"bytes"
	"encoding/binary"
)

// DeserializeVfStats - parse byte slice and returns a VfStats struct. It extracts the relevant VF statistics
// attributes from the provided byte slice.
func DeserializeVfStats(b []byte) VfStats {
	var vfstat VfStats
	stats, err := ParseRouteAttr(b)
	if err != nil {
		return vfstat
	}
	var valueVar uint64
	for _, stat := range stats {
		if err := binary.Read(bytes.NewBuffer(stat.Value), nativeEndian, &valueVar); err != nil {
			break
		}
		switch stat.Attr.Type {
		case IFLA_VF_STATS_RX_PACKETS:
			vfstat.RxPackets = valueVar
		case IFLA_VF_STATS_TX_PACKETS:
			vfstat.TxPackets = valueVar
		case IFLA_VF_STATS_RX_BYTES:
			vfstat.RxBytes = valueVar
		case IFLA_VF_STATS_TX_BYTES:
			vfstat.TxBytes = valueVar
		case IFLA_VF_STATS_MULTICAST:
			vfstat.Multicast = valueVar
		case IFLA_VF_STATS_BROADCAST:
			vfstat.Broadcast = valueVar
		case IFLA_VF_STATS_RX_DROPPED:
			vfstat.RxDropped = valueVar
		case IFLA_VF_STATS_TX_DROPPED:
			vfstat.TxDropped = valueVar
		default:
			func() {}() //Do nothing but make the IDE happy
		}
	}
	return vfstat
}
