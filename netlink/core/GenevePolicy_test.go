//go:build geneve

package core

import (
	"testing"
	"unsafe"
)

// TestGenevePolicy verifies the correct configuration of the GenevePolicy attributes.
func TestGenevePolicy(t *testing.T) {
	tests := []struct {
		index         IflaGeneveEnum
		expectedType  NlaType
		expectedLen   int
		expectedStart IflaGeneveEnum // For StrictStartType
	}{
		{IflaGeneveUnspec, 0, 0, IflaGeneveInnerProtoInherit},
		{IflaGeneveId, NlaU32, 0, 0},
		{IflaGeneveRemote, 0, int(unsafe.Sizeof(uint32(0))), 0},
		{IflaGeneveRemote6, 0, int(unsafe.Sizeof([16]byte{})), 0},
		{IflaGeneveTtl, NlaU8, 0, 0},
		{IflaGeneveTos, NlaU8, 0, 0},
		{IflaGeneveLabel, NlaU32, 0, 0},
		{IflaGenevePort, NlaU16, 0, 0},
		{IflaGeneveCollectMetadata, NlaFlag, 0, 0},
		{IflaGeneveUdpCsum, NlaU8, 0, 0},
		{IflaGeneveUdpZeroCsum6Tx, NlaU8, 0, 0},
		{IflaGeneveUdpZeroCsum6Rx, NlaU8, 0, 0},
		{IflaGeneveTtlInherit, NlaU8, 0, 0},
		{IflaGeneveDf, NlaU8, 0, 0},
		{IflaGeneveInnerProtoInherit, NlaFlag, 0, 0},
	}

	for _, tt := range tests {
		t.Run("Testing GenevePolicy", func(t *testing.T) {
			attr := GenevePolicy[tt.index]

			// Check type
			if attr.Type != tt.expectedType {
				t.Errorf("Expected Type %v for index %d, but got %v", tt.expectedType, tt.index, attr.Type)
			}

			// Check length
			if attr.Len != tt.expectedLen {
				t.Errorf("Expected Len %d for index %d, but got %d", tt.expectedLen, tt.index, attr.Len)
			}

			// Check StrictStartType
			if attr.StrictStartType != tt.expectedStart {
				t.Errorf("Expected StrictStartType %v for index %d, but got %v", tt.expectedStart, tt.index, attr.StrictStartType)
			}
		})
	}
}
