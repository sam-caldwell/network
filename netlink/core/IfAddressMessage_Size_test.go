package core

import (
	"testing"
)

func TestIfAddressMessage_Size(t *testing.T) {
	const expectedSizeInBytes = 8
	if IfAddressMessageSize != expectedSizeInBytes {
		t.Fatalf("IfAddressMessageSize mismatch")
	}
}
