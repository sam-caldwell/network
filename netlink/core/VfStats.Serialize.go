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
	writeStat := func(attrType uint16, value uint64) error {
		attr := RtAttr{
			RtAttr: unix.RtAttr{
				Type: attrType,
				Len:  uint16(binary.Size(value) + unix.SizeofRtAttr),
			},
			Data: make([]byte, 8),
		}
		binary.LittleEndian.PutUint64(attr.Data, value)
		if err := binary.Write(buf, nativeEndian, &attr.RtAttr); err != nil {
			return err
		}
		if _, err := buf.Write(attr.Data); err != nil {
			return err
		}
		return nil
	}

	// Serialize each VfStats field into an RtAttr and write to the buffer
	if err := writeStat(IFLA_VF_STATS_RX_PACKETS, stats.RxPackets); err != nil {
		return nil, err
	}
	if err := writeStat(IFLA_VF_STATS_TX_PACKETS, stats.TxPackets); err != nil {
		return nil, err
	}
	if err := writeStat(IFLA_VF_STATS_RX_BYTES, stats.RxBytes); err != nil {
		return nil, err
	}
	if err := writeStat(IFLA_VF_STATS_TX_BYTES, stats.TxBytes); err != nil {
		return nil, err
	}
	if err := writeStat(IFLA_VF_STATS_MULTICAST, stats.Multicast); err != nil {
		return nil, err
	}
	if err := writeStat(IFLA_VF_STATS_BROADCAST, stats.Broadcast); err != nil {
		return nil, err
	}
	if err := writeStat(IFLA_VF_STATS_RX_DROPPED, stats.RxDropped); err != nil {
		return nil, err
	}
	if err := writeStat(IFLA_VF_STATS_TX_DROPPED, stats.TxDropped); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
