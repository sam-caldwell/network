package core

import (
	"testing"
)

func TestBridgeVlanInfo_Size(t *testing.T) {
	const expectedSizeInBytes = 4
	if BridgeVlanInfoSize != expectedSizeInBytes {
		t.Fatalf("BridgeVlanInfoSize mismatch")
	}
}
