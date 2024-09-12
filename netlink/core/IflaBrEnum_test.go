package core

import (
	"testing"
	"unsafe"
)

// TestIflaBrEnum tests the size of the IflaBrEnum type and the values of its enumerated constants.
func TestIflaBrEnum(t *testing.T) {
	// Verify the size of the IflaBrEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaBrUnspec); size != 1 {
		t.Errorf("Expected size of IflaBrEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaBrEnum
		want  IflaBrEnum
	}{
		{"IflaBrUnspec", IflaBrUnspec, 0},
		{"IflaBrForwardDelay", IflaBrForwardDelay, 1},
		{"IflaBrHelloTime", IflaBrHelloTime, 2},
		{"IflaBrMaxAge", IflaBrMaxAge, 3},
		{"IflaBrAgeingTime", IflaBrAgeingTime, 4},
		{"IflaBrStpState", IflaBrStpState, 5},
		{"IflaBrPriority", IflaBrPriority, 6},
		{"IflaBrVlanFiltering", IflaBrVlanFiltering, 7},
		{"IflaBrVlanProtocol", IflaBrVlanProtocol, 8},
		{"IflaBrGroupFwdMask", IflaBrGroupFwdMask, 9},
		{"IflaBrRootId", IflaBrRootId, 10},
		{"IflaBrBridgeId", IflaBrBridgeId, 11},
		{"IflaBrRootPort", IflaBrRootPort, 12},
		{"IflaBrRootPathCost", IflaBrRootPathCost, 13},
		{"IflaBrTopologyChange", IflaBrTopologyChange, 14},
		{"IflaBrTopologyChangeDetected", IflaBrTopologyChangeDetected, 15},
		{"IflaBrHelloTimer", IflaBrHelloTimer, 16},
		{"IflaBrTcnTimer", IflaBrTcnTimer, 17},
		{"IflaBrTopologyChangeTimer", IflaBrTopologyChangeTimer, 18},
		{"IflaBrGcTimer", IflaBrGcTimer, 19},
		{"IflaBrGroupAddr", IflaBrGroupAddr, 20},
		{"IflaBrFdbFlush", IflaBrFdbFlush, 21},
		{"IflaBrMcastRouter", IflaBrMcastRouter, 22},
		{"IflaBrMcastSnooping", IflaBrMcastSnooping, 23},
		{"IflaBrMcastQueryUseIfaddr", IflaBrMcastQueryUseIfaddr, 24},
		{"IflaBrMcastQuerier", IflaBrMcastQuerier, 25},
		{"IflaBrMcastHashElasticity", IflaBrMcastHashElasticity, 26},
		{"IflaBrMcastHashMax", IflaBrMcastHashMax, 27},
		{"IflaBrMcastLastMemberCnt", IflaBrMcastLastMemberCnt, 28},
		{"IflaBrMcastStartupQueryCnt", IflaBrMcastStartupQueryCnt, 29},
		{"IflaBrMcastLastMemberIntvl", IflaBrMcastLastMemberIntvl, 30},
		{"IflaBrMcastMembershipIntvl", IflaBrMcastMembershipIntvl, 31},
		{"IflaBrMcastQuerierIntvl", IflaBrMcastQuerierIntvl, 32},
		{"IflaBrMcastQueryIntvl", IflaBrMcastQueryIntvl, 33},
		{"IflaBrMcastQueryResponseIntvl", IflaBrMcastQueryResponseIntvl, 34},
		{"IflaBrMcastStartupQueryIntvl", IflaBrMcastStartupQueryIntvl, 35},
		{"IflaBrNfCallIptables", IflaBrNfCallIptables, 36},
		{"IflaBrNfCallIp6tables", IflaBrNfCallIp6tables, 37},
		{"IflaBrNfCallArptables", IflaBrNfCallArptables, 38},
		{"IflaBrVlanDefaultPvid", IflaBrVlanDefaultPvid, 39},
		{"IflaBrPad", IflaBrPad, 40},
		{"IflaBrVlanStatsEnabled", IflaBrVlanStatsEnabled, 41},
		{"IflaBrMcastStatsEnabled", IflaBrMcastStatsEnabled, 42},
		{"IflaBrMcastIgmpVersion", IflaBrMcastIgmpVersion, 43},
		{"IflaBrMcastMldVersion", IflaBrMcastMldVersion, 44},
		{"IflaBrMax", IflaBrMax, 44}, // Since it's the last constant, IflaBrMax equals IflaBrMcastMldVersion
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
