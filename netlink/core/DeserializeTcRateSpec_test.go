package core

import (
	"encoding/binary"
	"testing"
)

// TestDeserializeTcRateSpec tests the DeserializeTcRateSpec function.
func TestDeserializeTcRateSpec(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Prepare a byte slice for TcRateSpec
		data := make([]byte, SizeOfTcRateSpec)

		// Populate the byte slice with sample data
		data[0] = 0x01                                 // CellLog (uint8)
		data[1] = 0x02                                 // Linklayer (uint8)
		binary.LittleEndian.PutUint16(data[2:], 0x03)  // Overhead (uint16)
		binary.LittleEndian.PutUint16(data[4:], 0x04)  // CellAlign (int16)
		binary.LittleEndian.PutUint16(data[6:], 0x05)  // Mpu (uint16)
		binary.LittleEndian.PutUint32(data[8:], 0x100) // Rate (uint32)

		// Call DeserializeTcRateSpec
		result := DeserializeTcRateSpec(data)

		// Verify the result
		if result == nil {
			t.Fatal("Expected non-nil TcRateSpec, but got nil")
		}

		if result.CellLog != 0x01 {
			t.Errorf("Expected CellLog 0x01, got 0x%02x", result.CellLog)
		}
		if result.Linklayer != 0x02 {
			t.Errorf("Expected Linklayer 0x02, got 0x%02x", result.Linklayer)
		}
		if result.Overhead != 0x03 {
			t.Errorf("Expected Overhead 0x03, got 0x%02x", result.Overhead)
		}
		if result.CellAlign != 0x04 {
			t.Errorf("Expected CellAlign 0x04, got 0x%02x", result.CellAlign)
		}
		if result.Mpu != 0x05 {
			t.Errorf("Expected Mpu 0x05, got 0x%02x", result.Mpu)
		}
		if result.Rate != 0x100 {
			t.Errorf("Expected Rate 0x100, got 0x%02x", result.Rate)
		}
	})
	t.Run("input too short", func(t *testing.T) {

		// TestDeserializeTcRateSpecShortInput tests that the function returns nil for short input.

		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcRateSpec-1)

		// Call DeserializeTcRateSpec
		result := DeserializeTcRateSpec(shortData)

		if result != nil {
			t.Fatalf("Expected nil for short input, but got non-nil result")
		}
	})
}
