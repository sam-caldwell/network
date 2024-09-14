package core

import (
	"bytes"
	"testing"
)

func TestNetlinkRequest_AddRawData(t *testing.T) {
	// Create a new NetlinkRequest instance.
	req := &NetlinkRequest{}

	// Define some raw data to add.
	rawData1 := []byte{0x01, 0x02, 0x03, 0x04}
	rawData2 := []byte{0x05, 0x06}

	// Add the first raw data to the request.
	req.AddRawData(rawData1)

	// Check if the raw data was added correctly.
	if !bytes.Equal(req.RawData, rawData1) {
		t.Errorf("Expected raw data %v, got %v", rawData1, req.RawData)
	}

	// Add more raw data to the request.
	req.AddRawData(rawData2)

	// Create the expected result by combining the first and second raw data.
	expectedData := append(rawData1, rawData2...)

	// Check if the combined raw data was appended correctly.
	if !bytes.Equal(req.RawData, expectedData) {
		t.Errorf("Expected combined raw data %v, got %v", expectedData, req.RawData)
	}
}
