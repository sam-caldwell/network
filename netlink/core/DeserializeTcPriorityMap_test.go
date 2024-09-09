package core

import (
	"encoding/binary"
	"testing"
)

// TestDeserializeTcPriorityMap tests the DeserializeTcPriorityMap function.
func TestDeserializeTcPriorityMap(t *testing.T) {
	t.Run("happy test", func(t *testing.T) {
		// Prepare test data for TcPriorityMap
		data := make([]byte, SizeOfTcPriorityMap)

		// Bands field (int32)
		binary.LittleEndian.PutUint32(data[0:4], 1) // Setting Bands to 1

		// Priomap field (16 uint8 values)
		expectedPriomap := [16]uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
		copy(data[4:], expectedPriomap[:])

		// Call DeserializeTcPriorityMap
		result, err := DeserializeTcPriorityMap(data)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Verify the Bands field
		if result.Bands != 1 {
			t.Errorf("Expected Bands to be 1, but got %d", result.Bands)
		}

		// Verify the Priomap field
		for i, val := range expectedPriomap {
			if result.Priomap[i] != val {
				t.Errorf("Expected Priomap[%d] to be %d, but got %d", i, val, result.Priomap[i])
			}
		}
	})
	t.Run("short input", func(t *testing.T) {
		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcPriorityMap-1)

		// Call DeserializeTcPriorityMap
		result, err := DeserializeTcPriorityMap(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}
		if result != nil {
			t.Fatal("Expected nil result")
		}

		const expectedErr = "input too short"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message to start with %q, but got %q", expectedErr, err.Error())
		}
	})
}
