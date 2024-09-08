package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDeserializeTcHtbCopt(t *testing.T) {
	// Helper function to create a valid byte slice for TcHtbCopt
	createValidTcHtbCoptBytes := func() []byte {
		buf := new(bytes.Buffer)

		// Write values for TcHtbCopt fields (Rate, Ceil, Buffer, Cbuffer, Quantum, Level, Prio)
		// Assuming TcRateSpec is 8 bytes
		binary.Write(buf, NativeEndian, TcRateSpec{Rate: 10, CellLog: 1}) // Rate (TcRateSpec)
		binary.Write(buf, NativeEndian, TcRateSpec{Rate: 20, CellLog: 2}) // Ceil (TcRateSpec)
		binary.Write(buf, NativeEndian, uint32(100))                      // Buffer (uint32)
		binary.Write(buf, NativeEndian, uint32(200))                      // Cbuffer (uint32)
		binary.Write(buf, NativeEndian, uint32(300))                      // Quantum (uint32)
		binary.Write(buf, NativeEndian, uint32(400))                      // Level (uint32)
		binary.Write(buf, NativeEndian, uint32(500))                      // Prio (uint32)

		return buf.Bytes()
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcHtbCoptBytes()

		// Call DeserializeTcHtbCopt
		tcHtbCopt, err := DeserializeTcHtbCopt(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcHtbCopt.Rate.Rate != 10 || tcHtbCopt.Rate.CellLog != 1 {
			t.Errorf("Expected Rate to be {10, 1}, but got {%d, %d}", tcHtbCopt.Rate.Rate, tcHtbCopt.Rate.CellLog)
		}
		if tcHtbCopt.Ceil.Rate != 20 || tcHtbCopt.Ceil.CellLog != 2 {
			t.Errorf("Expected Ceil to be {20, 2}, but got {%d, %d}", tcHtbCopt.Ceil.Rate, tcHtbCopt.Ceil.CellLog)
		}
		if tcHtbCopt.Buffer != 100 {
			t.Errorf("Expected Buffer to be 100, but got %d", tcHtbCopt.Buffer)
		}
		if tcHtbCopt.Cbuffer != 200 {
			t.Errorf("Expected Cbuffer to be 200, but got %d", tcHtbCopt.Cbuffer)
		}
		if tcHtbCopt.Quantum != 300 {
			t.Errorf("Expected Quantum to be 300, but got %d", tcHtbCopt.Quantum)
		}
		if tcHtbCopt.Level != 400 {
			t.Errorf("Expected Level to be 400, but got %d", tcHtbCopt.Level)
		}
		if tcHtbCopt.Prio != 500 {
			t.Errorf("Expected Prio to be 500, but got %d", tcHtbCopt.Prio)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := make([]byte, SizeOfTcHtbCopt-10) // Too short

		// Call DeserializeTcHtbCopt
		tcHtbCopt, err := DeserializeTcHtbCopt(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcHtbCopt is nil
		if tcHtbCopt != nil {
			t.Errorf("Expected tcHtbCopt to be nil due to short input, but got %v", tcHtbCopt)
		}
	})
}
