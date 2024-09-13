package core

import (
	"encoding/binary"
	"testing"
)

func TestGenericNetlinkMessage_Len(t *testing.T) {
	// Create an instance of GenericNetlinkMessage with some values
	msg := &GenericNetlinkMessage{
		Command: 0x10, // Example command value
		Version: 0x01, // Example version value
	}

	// Call Len() to get the length
	length := msg.Len()

	// Calculate the expected length using binary.Size directly
	expectedLength := binary.Size(*msg)

	// Check if the calculated length matches the expected length
	if length != expectedLength {
		t.Errorf("Expected length to be %d, but got %d", expectedLength, length)
	}
}
