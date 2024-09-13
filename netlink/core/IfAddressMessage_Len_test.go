package core

import (
	"testing"
)

func TestIfAddressMessage_Len(t *testing.T) {
	const expectedSizeInBytes = IfAddressMessageSize
	var msg IfAddressMessage
	if msg.Len() != expectedSizeInBytes {
		t.Fatalf("Len() method did not return expected size")
	}
}
