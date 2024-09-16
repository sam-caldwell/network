package core

import (
	"testing"
	"unsafe"
)

// TestIflaIpTunEnum tests the size of the IflaIpTun type and the values of its enumerated constants.
func TestIflaIpTunEnum(t *testing.T) {
	// Verify that the size of the IflaIpTun type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaIptunUnspec); size != 1 {
		t.Errorf("Expected size of IflaIpTun to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaIpTun
		want  IflaIpTun
	}{
		{"IflaIptunUnspec", IflaIptunUnspec, 0},
		{"IflaIptunLink", IflaIptunLink, 1},
		{"IflaIptunLocal", IflaIptunLocal, 2},
		{"IflaIptunRemote", IflaIptunRemote, 3},
		{"IflaIptunTtl", IflaIptunTtl, 4},
		{"IflaIptunTos", IflaIptunTos, 5},
		{"IflaIptunEncapLimit", IflaIptunEncapLimit, 6},
		{"IflaIptunFlowinfo", IflaIptunFlowinfo, 7},
		{"IflaIptunFlags", IflaIptunFlags, 8},
		{"IflaIptunProto", IflaIptunProto, 9},
		{"IflaIptunPmtudisc", IflaIptunPmtudisc, 10},
		{"IflaIptun6rdPrefix", IflaIptun6rdPrefix, 11},
		{"IflaIptun6rdRelayPrefix", IflaIptun6rdRelayPrefix, 12},
		{"IflaIptun6rdPrefixlen", IflaIptun6rdPrefixlen, 13},
		{"IflaIptun6rdRelayPrefixlen", IflaIptun6rdRelayPrefixlen, 14},
		{"IflaIptunEncapType", IflaIptunEncapType, 15},
		{"IflaIptunEncapFlags", IflaIptunEncapFlags, 16},
		{"IflaIptunEncapSport", IflaIptunEncapSport, 17},
		{"IflaIptunEncapDport", IflaIptunEncapDport, 18},
		{"IflaIptunCollectMetadata", IflaIptunCollectMetadata, 19},
		{"IflaIptunMax", IflaIptunMax, 19},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
