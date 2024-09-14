package core

import (
	"testing"
)

// See NetlinkRequestData interface--
// 		MockNetlinkRequestData 	- A mock implementation of the NetlinkRequestData interface for testing.
// 		Len 					- Returns the length of the mock data.
// 		Serialize 				- Serializes the mock data into a byte slice.

func TestNetlinkRequest_AddData(t *testing.T) {
	// Create a new NetlinkRequest instance.
	req := &NetlinkRequest{}

	// Create a mock data instance.
	mockData := &MockNetlinkRequestData{
		data: []byte{0x01, 0x02, 0x03, 0x04},
	}

	// Add the mock data to the request.
	req.AddData(mockData)

	// Check if the data was added correctly.
	if len(req.Data) != 1 {
		t.Fatalf("Expected 1 data entry, got %d", len(req.Data))
	}

	// Verify that the added data is the mock data.
	if req.Data[0] != mockData {
		t.Errorf("Expected data %v, got %v", mockData, req.Data[0])
	}

	// Add more data and check if the Data slice grows correctly.
	anotherMockData := &MockNetlinkRequestData{
		data: []byte{0x05, 0x06},
	}
	req.AddData(anotherMockData)

	if len(req.Data) != 2 {
		t.Fatalf("Expected 2 data entries, got %d", len(req.Data))
	}

	if req.Data[1] != anotherMockData {
		t.Errorf("Expected data %v, got %v", anotherMockData, req.Data[1])
	}
}
