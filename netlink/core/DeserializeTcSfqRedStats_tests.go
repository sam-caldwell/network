package core

import (
	"encoding/binary"
	"testing"
)

// TestDeserializeTcSfqRedStats tests the DeserializeTcSfqRedStats function.
func TestDeserializeTcSfqRedStats(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Prepare a byte slice for TcSfqRedStats
		data := make([]byte, SizeOfTcSfqRedStats)

		// Fill the byte slice with sample values
		binary.LittleEndian.PutUint32(data[0:], 0x01020304)  // ProbDrop
		binary.LittleEndian.PutUint32(data[4:], 0x05060708)  // ForcedDrop
		binary.LittleEndian.PutUint32(data[8:], 0x090A0B0C)  // ProbMark
		binary.LittleEndian.PutUint32(data[12:], 0x0D0E0F10) // ForcedMark
		binary.LittleEndian.PutUint32(data[16:], 0x11121314) // ProbMarkHead
		binary.LittleEndian.PutUint32(data[20:], 0x15161718) // ForcedMarkHead

		// Call DeserializeTcSfqRedStats
		result, err := DeserializeTcSfqRedStats(data)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Verify the result
		if result.ProbDrop != 0x01020304 {
			t.Errorf("Expected ProbDrop 0x01020304, got 0x%08x", result.ProbDrop)
		}
		if result.ForcedDrop != 0x05060708 {
			t.Errorf("Expected ForcedDrop 0x05060708, got 0x%08x", result.ForcedDrop)
		}
		if result.ProbMark != 0x090A0B0C {
			t.Errorf("Expected ProbMark 0x090A0B0C, got 0x%08x", result.ProbMark)
		}
		if result.ForcedMark != 0x0D0E0F10 {
			t.Errorf("Expected ForcedMark 0x0D0E0F10, got 0x%08x", result.ForcedMark)
		}
		if result.ProbMarkHead != 0x11121314 {
			t.Errorf("Expected ProbMarkHead 0x11121314, got 0x%08x", result.ProbMarkHead)
		}
		if result.ForcedMarkHead != 0x15161718 {
			t.Errorf("Expected ForcedMarkHead 0x15161718, got 0x%08x", result.ForcedMarkHead)
		}
	})

	t.Run("test DeserializeTcSfqRedStats with short input", func(t *testing.T) {
		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcSfqRedStats-1)

		// Call DeserializeTcSfqRedStats
		_, err := DeserializeTcSfqRedStats(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		expectedErr := "input too short"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
