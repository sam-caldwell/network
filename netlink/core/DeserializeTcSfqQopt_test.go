package core

import (
	"encoding/binary"
	"testing"
)

// TestDeserializeTcSfqQopt tests the DeserializeTcSfqQopt function.
func TestDeserializeTcSfqQopt(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Prepare a byte slice for TcSfqQopt
		data := make([]byte, SizeOfTcSfqQopt)

		// Populate the byte slice with sample data
		data[0] = 0x01                                      // Quantum (uint8)
		binary.LittleEndian.PutUint32(data[1:], 0x02030405) // Perturb (int32)
		binary.LittleEndian.PutUint32(data[5:], 0x06070809) // Limit (uint32)
		data[9] = 0x0A                                      // Divisor (uint8)
		data[10] = 0x0B                                     // Flows (uint8)

		// Call DeserializeTcSfqQopt
		result, err := DeserializeTcSfqQopt(data)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Verify the result
		if result.Quantum != 0x01 {
			t.Errorf("Expected Quantum 0x01, got 0x%02x", result.Quantum)
		}
		if result.Perturb != 0x02030405 {
			t.Errorf("Expected Perturb 0x02030405, got 0x%08x", result.Perturb)
		}
		if result.Limit != 0x06070809 {
			t.Errorf("Expected Limit 0x06070809, got 0x%08x", result.Limit)
		}
		if result.Divisor != 0x0A {
			t.Errorf("Expected Divisor 0x0A, got 0x%02x", result.Divisor)
		}
		if result.Flows != 0x0B {
			t.Errorf("Expected Flows 0x0B, got 0x%02x", result.Flows)
		}
	})

	t.Run("TestDeserializeTcSfqQoptShortInput tests DeserializeTcSfqQopt with a short input.", func(t *testing.T) {

		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcSfqQopt-1)

		// Call DeserializeTcSfqQopt
		_, err := DeserializeTcSfqQopt(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		expectedErr := "input too short"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
