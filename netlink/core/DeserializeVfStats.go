//go:build linux

package core

import (
	"bytes"
	"encoding/binary"
)

// DeserializeVfStats - parse byte slice and returns a VfStats struct. It extracts the relevant VF statistics
// attributes from the provided byte slice.
func DeserializeVfStats(b []byte) VfStats {
	var (
		err      error
		vfstat   VfStats
		valueVar uint64
		stats    []NetlinkRouteAttr
	)

	if stats, err = ParseRouteAttr(b); err != nil {
		return vfstat
	}

	for _, stat := range stats {
		if err := binary.Read(bytes.NewBuffer(stat.Value), NativeEndian, &valueVar); err != nil {
			break
		}
		switch IflaVfStatsEnum(stat.Attr.Type) {
		case IflaVfStatsRxPackets:
			vfstat.RxPackets = valueVar
		case IflaVfStatsTxPackets:
			vfstat.TxPackets = valueVar
		case IflaVfStatsRxBytes:
			vfstat.RxBytes = valueVar
		case IflaVfStatsTxBytes:
			vfstat.TxBytes = valueVar
		case IflaVfStatsMulticast:
			vfstat.Multicast = valueVar
		case IflaVfStatsBroadcast:
			vfstat.Broadcast = valueVar
		case IflaVfStatsRxDropped:
			vfstat.RxDropped = valueVar
		case IflaVfStatsTxDropped:
			vfstat.TxDropped = valueVar
		default:
			func() {}() //Do nothing but make the IDE happy
		}
	}
	return vfstat
}
