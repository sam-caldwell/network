package core

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
)

// See NetlinkRequestData interface--
// 		MockNetlinkRequestData 	- A mock implementation of the NetlinkRequestData interface for testing.
// 		Len 					- Returns the length of the mock data.
// 		Serialize 				- Serializes the mock data into a byte slice.

func TestNetlinkRequest_Serialize(t *testing.T) {
	// Create mock data for NetlinkRequestData
	mockData := &MockNetlinkRequestData{
		data: []byte{0x01, 0x02, 0x03, 0x04},
	}

	// Create a NetlinkRequest with mock data
	req := &NetlinkRequest{
		NlMsghdr: unix.NlMsghdr{
			Type:  uint16(16),
			Flags: uint16(0),
			Seq:   uint32(1),
			Pid:   uint32(1234),
		},
		Data:    []NetlinkRequestData{mockData},
		RawData: []byte{0x05, 0x06},
	}

	// Expected length of the serialized output
	expectedLength := NetlinkMessageHeaderSize + mockData.Len() + len(req.RawData)

	// Perform serialization
	serialized, err := req.Serialize()
	if err != nil {
		t.Fatalf("Unexpected error during serialization: %v", err)
	}

	// Verify the total length of the serialized output
	if len(serialized) != expectedLength {
		t.Errorf("Expected serialized length %d, got %d", expectedLength, len(serialized))
	}

	// Verify the serialized header
	buf := bytes.NewReader(serialized)
	var header unix.NlMsghdr
	if err := binary.Read(buf, NativeEndian, &header); err != nil {
		t.Fatalf("Error reading serialized header: %v", err)
	}

	if header.Len != uint32(expectedLength) {
		t.Errorf("Expected header Len %d, got %d", expectedLength, header.Len)
	}
	if header.Type != req.Type {
		t.Errorf("Expected header Type %d, got %d", req.Type, header.Type)
	}
	if header.Flags != req.Flags {
		t.Errorf("Expected header Flags %d, got %d", req.Flags, header.Flags)
	}
	if header.Seq != req.Seq {
		t.Errorf("Expected header Seq %d, got %d", req.Seq, header.Seq)
	}
	if header.Pid != req.Pid {
		t.Errorf("Expected header Pid %d, got %d", req.Pid, header.Pid)
	}

	// Verify the serialized payload data
	expectedPayload := append(mockData.data, req.RawData...)
	actualPayload := serialized[NetlinkMessageHeaderSize:]

	if !bytes.Equal(expectedPayload, actualPayload) {
		t.Errorf("Expected payload %v, got %v", expectedPayload, actualPayload)
	}
}
