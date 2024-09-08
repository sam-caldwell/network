package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDeserializeTcNetemQopt(t *testing.T) {
	// Helper function to create a valid byte slice for TcNetemQopt
	createValidTcNetemQoptBytes := func() []byte {
		buf := new(bytes.Buffer)

		// Write values for TcNetemQopt fields: Latency, Limit, Loss, Gap, Duplicate, Jitter (all uint32)
		binary.Write(buf, binary.LittleEndian, uint32(1)) // Latency
		binary.Write(buf, binary.LittleEndian, uint32(2)) // Limit
		binary.Write(buf, binary.LittleEndian, uint32(3)) // Loss
		binary.Write(buf, binary.LittleEndian, uint32(4)) // Gap
		binary.Write(buf, binary.LittleEndian, uint32(5)) // Duplicate
		binary.Write(buf, binary.LittleEndian, uint32(6)) // Jitter

		return buf.Bytes()
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcNetemQoptBytes()

		// Call DeserializeTcNetemQopt
		tcNetemQopt, err := DeserializeTcNetemQopt(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcNetemQopt.Latency != 1 {
			t.Errorf("Expected Latency to be 1, but got %d", tcNetemQopt.Latency)
		}
		if tcNetemQopt.Limit != 2 {
			t.Errorf("Expected Limit to be 2, but got %d", tcNetemQopt.Limit)
		}
		if tcNetemQopt.Loss != 3 {
			t.Errorf("Expected Loss to be 3, but got %d", tcNetemQopt.Loss)
		}
		if tcNetemQopt.Gap != 4 {
			t.Errorf("Expected Gap to be 4, but got %d", tcNetemQopt.Gap)
		}
		if tcNetemQopt.Duplicate != 5 {
			t.Errorf("Expected Duplicate to be 5, but got %d", tcNetemQopt.Duplicate)
		}
		if tcNetemQopt.Jitter != 6 {
			t.Errorf("Expected Jitter to be 6, but got %d", tcNetemQopt.Jitter)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := make([]byte, SizeOfTcNetemQopt-10) // Too short

		// Call DeserializeTcNetemQopt
		tcNetemQopt, err := DeserializeTcNetemQopt(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcNetemQopt is nil
		if tcNetemQopt != nil {
			t.Errorf("Expected tcNetemQopt to be nil due to short input, but got %v", tcNetemQopt)
		}
	})
}
