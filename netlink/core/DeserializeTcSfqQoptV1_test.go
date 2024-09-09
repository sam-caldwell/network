package core

import (
	"testing"
)

// TestDeserializeTcSfqQoptV1 tests the DeserializeTcSfqQoptV1 function.
func TestDeserializeTcSfqQoptV1(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Create a byte slice for TcSfqQoptV1
		validBytes := []byte{
			// TcSfqQoptV1
			// TcSfqQopt
			0x01,                   //                   Quantum uint8
			0x05, 0x04, 0x03, 0x02, // Perturb int32
			0x03, 0x00, 0x00, 0x00, // Limit   uint32
			0x04, //                   Divisor uint8
			0x05, //                   Flows   uint8
			0x00, //                   Padding
			// TcSfqQoptV1
			0x06, 0x00, 0x00, 0x00, // Depth uint32
			0x0F, 0x0E, 0x0D, 0x0C, // HeadDrop uint32
			0x13, 0x12, 0x11, 0x10, // Limit uint32
			0x09, 0x00, 0x00, 0x00, // QthMin uint32
			0x0A, 0x00, 0x00, 0x00, // QthMax uint32
			0x0B,                   // Wlog byte
			0x0C,                   // Plog byte
			0x0D,                   // ScellLog byte
			0x0E,                   // Flag byte
			0x0F, 0x00, 0x00, 0x00, // MaxP uint32
			// TcSfqRedStats
			0x10, 0x00, 0x00, 0x00, // ProbDrop       uint32
			0x2B, 0x2A, 0x29, 0x28, // ForcedDrop     uint32
			0x12, 0x00, 0x00, 0x00, // ProbMark       uint32
			0x13, 0x00, 0x00, 0x00, // ForcedMark     uint32
			0x14, 0x00, 0x00, 0x00, // ProbMarkHead   uint32
			0x00, 0x00, 0x00, 0x00, // padding
			0x3F, 0x3E, 0x3D, 0x3C, // ForcedMarkHead uint32
		}
		if len(validBytes) != SizeOfTcSfqQoptV1 {
			t.Errorf("invalid length.  Want %d, got %d", SizeOfTcSfqQoptV1, len(validBytes))
		}

		// Call DeserializeTcSfqQoptV1
		result, err := DeserializeTcSfqQoptV1(validBytes)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Verify the result (for brevity, checking some fields)
		if result.Quantum != 0x01 {
			t.Errorf("Expected Quantum 0x01, got 0x%02x", result.Quantum)
		}
		if result.Perturb != 0x02030405 {
			t.Errorf("Expected Perturb 0x02030405, got 0x%08x", result.Perturb)
		}
		if result.Depth != 0x0C0D0E0F {
			t.Errorf("Expected Depth 0x0C0D0E0F, got 0x%08x", result.Depth)
		}
		if result.HeadDrop != 0x10111213 {
			t.Errorf("Expected HeadDrop 0x10111213, got 0x%08x", result.HeadDrop)
		}
		if result.TcSfqRedStats.ProbDrop != 0x28292A2B {
			t.Errorf("Expected ProbDrop 0x28292A2B, got 0x%08x", result.TcSfqRedStats.ProbDrop)
		}
		if result.TcSfqRedStats.ForcedMarkHead != 0x3C3D3E3F {
			t.Errorf("Expected ForcedMarkHead 0x3C3D3E3F, got 0x%08x", result.TcSfqRedStats.ForcedMarkHead)
		}
	})
	t.Run("tests DeserializeTcSfqQoptV1 with short input.", func(t *testing.T) {
		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcSfqQoptV1-1)

		// Call DeserializeTcSfqQoptV1
		_, err := DeserializeTcSfqQoptV1(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		expectedErr := "input too short"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
