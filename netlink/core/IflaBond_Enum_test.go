package core

import (
	"testing"
	"unsafe"
)

// TestIflaBondEnum tests the size of the IflaBondAttribute type and the values of its enumerated constants.
func TestIflaBondEnum(t *testing.T) {
	// Verify the size of the IflaBondAttribute type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaBondUnspec); size != 1 {
		t.Errorf("Expected size of IflaBondAttribute to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaBondAttribute
		want  IflaBondAttribute
	}{
		{"IflaBondUnspec", IflaBondUnspec, 0},
		{"IflaBondMode", IflaBondMode, 1},
		{"IflaBondActiveSlave", IflaBondActiveSlave, 2},
		{"IflaBondMiimon", IflaBondMiimon, 3},
		{"IflaBondUpdelay", IflaBondUpdelay, 4},
		{"IflaBondDowndelay", IflaBondDowndelay, 5},
		{"IflaBondUseCarrier", IflaBondUseCarrier, 6},
		{"IflaBondArpInterval", IflaBondArpInterval, 7},
		{"IflaBondArpIpTarget", IflaBondArpIpTarget, 8},
		{"IflaBondArpValidate", IflaBondArpValidate, 9},
		{"IflaBondArpAllTargets", IflaBondArpAllTargets, 10},
		{"IflaBondPrimary", IflaBondPrimary, 11},
		{"IflaBondPrimaryReselect", IflaBondPrimaryReselect, 12},
		{"IflaBondFailOverMac", IflaBondFailOverMac, 13},
		{"IflaBondXmitHashPolicy", IflaBondXmitHashPolicy, 14},
		{"IflaBondResendIgmp", IflaBondResendIgmp, 15},
		{"IflaBondNumPeerNotif", IflaBondNumPeerNotif, 16},
		{"IflaBondAllSlavesActive", IflaBondAllSlavesActive, 17},
		{"IflaBondMinLinks", IflaBondMinLinks, 18},
		{"IflaBondLpInterval", IflaBondLpInterval, 19},
		{"IflaBondPacketsPerSlave", IflaBondPacketsPerSlave, 20},
		{"IflaBondAdLacpRate", IflaBondAdLacpRate, 21},
		{"IflaBondAdSelect", IflaBondAdSelect, 22},
		{"IflaBondAdInfoAttr", IflaBondAdInfoAttr, 23},
		{"IflaBondAdActorSysPrio", IflaBondAdActorSysPrio, 24},
		{"IflaBondAdUserPortKey", IflaBondAdUserPortKey, 25},
		{"IflaBondAdActorSystem", IflaBondAdActorSystem, 26},
		{"IflaBondTlbDynamicLb", IflaBondTlbDynamicLb, 27},
		{"IflaBondPeerNotifDelay", IflaBondPeerNotifDelay, 28},
		{"IflaBondAdLacpActive", IflaBondAdLacpActive, 29},
		{"IflaBondMissedMax", IflaBondMissedMax, 30},
		{"IflaBondNsIp6Target", IflaBondNsIp6Target, 31},
		{"IflaBondCoupledControl", IflaBondCoupledControl, 32},
		{"IflaBondMax", IflaBondMax, 32}, // Max should be 32 based on the iota - 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
