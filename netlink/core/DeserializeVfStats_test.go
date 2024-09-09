package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TestDeserializeVfStats tests the DeserializeVfStats function.
func TestDeserializeVfStats(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Create a sample VfStats structure for testing
		expectedStats := VfStats{
			RxPackets: 1000,
			TxPackets: 2000,
			RxBytes:   3000,
			TxBytes:   4000,
			Multicast: 5000,
			Broadcast: 6000,
			RxDropped: 7000,
			TxDropped: 8000,
		}

		// Create a byte slice representing these stats using NetlinkRouteAttr structure
		buf := new(bytes.Buffer)

		// Helper function to serialize a stat
		writeStat := func(statType IflaVfStatsEnum, value uint64) {
			_ = binary.Write(buf, NativeEndian, statType)
			_ = binary.Write(buf, NativeEndian, value)
		}

		writeStat(IflaVfStatsRxPackets, expectedStats.RxPackets)
		writeStat(IflaVfStatsTxPackets, expectedStats.TxPackets)
		writeStat(IflaVfStatsRxBytes, expectedStats.RxBytes)
		writeStat(IflaVfStatsTxBytes, expectedStats.TxBytes)
		writeStat(IflaVfStatsMulticast, expectedStats.Multicast)
		writeStat(IflaVfStatsBroadcast, expectedStats.Broadcast)
		writeStat(IflaVfStatsRxDropped, expectedStats.RxDropped)
		writeStat(IflaVfStatsTxDropped, expectedStats.TxDropped)

		serializedData := buf.Bytes()

		// Deserialize the byte slice into a VfStats structure
		deserializedStats, err := DeserializeVfStats(serializedData)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Check if the deserialized structure matches the expected structure
		if deserializedStats != expectedStats {
			t.Fatalf("Deserialized object doesn't match expected.\n"+
				"Expected: %+v\n"+
				"Got: %+v",
				expectedStats, deserializedStats)
		}
	})

	t.Run("empty data", func(t *testing.T) {
		// Test with an empty byte slice
		var emptyData []byte
		deserializedStats, err := DeserializeVfStats(emptyData)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Expect all fields to be zero in the deserialized structure
		expectedEmptyStats := VfStats{}
		if deserializedStats != expectedEmptyStats {
			t.Fatalf("Expected empty VfStats, but got: %+v", deserializedStats)
		}
	})
}
