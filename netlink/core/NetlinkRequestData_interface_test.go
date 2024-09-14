package core

import (
	"testing"
)

func TestNetlinkRequestDataInterface(t *testing.T) {
	// Create a new instance of MockNetlinkRequestData
	mockData := &MockNetlinkRequestData{
		data: []byte{0x01, 0x02, 0x03, 0x04},
	}

	// Verify that mockData implements the NetlinkRequestData interface
	var _ NetlinkRequestData = mockData

	// Test the Len method
	expectedLen := 4
	if mockData.Len() != expectedLen {
		t.Errorf("Expected length %d, got %d", expectedLen, mockData.Len())
	}

	// Test the Serialize method
	serializedData, err := mockData.Serialize()
	if err != nil {
		t.Errorf("Unexpected error during serialization: %v", err)
	}

	expectedData := []byte{0x01, 0x02, 0x03, 0x04}
	if !equalByteSlices(serializedData, expectedData) {
		t.Errorf("Expected serialized data %v, got %v", expectedData, serializedData)
	}
}

// Helper function to compare two byte slices
func equalByteSlices(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
