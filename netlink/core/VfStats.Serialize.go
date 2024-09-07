//go:build linux

package core

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
)

// Serialize converts the VfStats struct into a byte slice.
//
// The byte slice contains the VF statistics encapsulated as RtAttr structures.
func (stats *VfStats) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Helper function to create and write an RtAttr to the buffer
	writeStat := func(attrType IflaVfStatsEnum, value uint64) error {
		attr := RtAttr{
			RtAttr: unix.RtAttr{
				Type: uint16(attrType),
				Len:  uint16(binary.Size(value) + SizeOfUnixRtAttr),
			},
			Data: make([]byte, 8),
		}
		NativeEndian.PutUint64(attr.Data, value)
		if err := binary.Write(buf, NativeEndian, &attr.RtAttr); err != nil {
			return err
		}
		if _, err := buf.Write(attr.Data); err != nil {
			return err
		}
		return nil
	}

	// Serialize each VfStats field into an RtAttr and write to the buffer
	if err := writeStat(IflaVfStatsRxPackets, stats.RxPackets); err != nil {
		return nil, err
	}
	if err := writeStat(IflaVfStatsTxPackets, stats.TxPackets); err != nil {
		return nil, err
	}
	if err := writeStat(IflaVfStatsRxBytes, stats.RxBytes); err != nil {
		return nil, err
	}
	if err := writeStat(IflaVfStatsTxBytes, stats.TxBytes); err != nil {
		return nil, err
	}
	if err := writeStat(IflaVfStatsMulticast, stats.Multicast); err != nil {
		return nil, err
	}
	if err := writeStat(IflaVfStatsBroadcast, stats.Broadcast); err != nil {
		return nil, err
	}
	if err := writeStat(IflaVfStatsRxDropped, stats.RxDropped); err != nil {
		return nil, err
	}
	if err := writeStat(IflaVfStatsTxDropped, stats.TxDropped); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
