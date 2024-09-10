package core

import (
	"testing"
	"unsafe"
)

func TestDevlinkCommandEnum(t *testing.T) {
	// Test size of DevlinkCommandEnum
	expectedSize := 1 // DevlinkCommandEnum is based on uint8, which is 1 byte
	actualSize := unsafe.Sizeof(DevlinkCmdUnspec)
	if actualSize != uintptr(expectedSize) {
		t.Errorf("Expected size of DevlinkCommandEnum: %d, got: %d", expectedSize, actualSize)
	}

	// Test correct values for DevlinkCommandEnum constants
	tests := []struct {
		name     string
		cmd      DevlinkCommandEnum
		expected uint8
	}{
		{"DevlinkCmdUnspec", DevlinkCmdUnspec, 0},
		{"DevlinkCmdGet", DevlinkCmdGet, 1},
		{"DevlinkCmdSet", DevlinkCmdSet, 2},
		{"DevlinkCmdNew", DevlinkCmdNew, 3},
		{"DevlinkCmdDel", DevlinkCmdDel, 3},
		{"DevlinkCmdPortGet", DevlinkCmdPortGet, 4},
		{"DevlinkCmdPortSet", DevlinkCmdPortSet, 5},
		{"DevlinkCmdPortNew", DevlinkCmdPortNew, 6},
		{"DevlinkCmdPortDel", DevlinkCmdPortDel, 7},
		{"DevlinkCmdPortSplit", DevlinkCmdPortSplit, 8},
		{"DevlinkCmdPortUnsplit", DevlinkCmdPortUnsplit, 9},
		{"DevlinkCmdSbGet", DevlinkCmdSbGet, 10},
		{"DevlinkCmdSbSet", DevlinkCmdSbSet, 11},
		{"DevlinkCmdSbNew", DevlinkCmdSbNew, 12},
		{"DevlinkCmdSbDel", DevlinkCmdSbDel, 13},
		{"DevlinkCmdSbPoolGet", DevlinkCmdSbPoolGet, 14},
		{"DevlinkCmdSbPoolSet", DevlinkCmdSbPoolSet, 15},
		{"DevlinkCmdSbPoolNew", DevlinkCmdSbPoolNew, 16},
		{"DevlinkCmdSbPoolDel", DevlinkCmdSbPoolDel, 17},
		{"DevlinkCmdSbPortPoolGet", DevlinkCmdSbPortPoolGet, 18},
		{"DevlinkCmdSbPortPoolSet", DevlinkCmdSbPortPoolSet, 19},
		{"DevlinkCmdSbPortPoolNew", DevlinkCmdSbPortPoolNew, 20},
		{"DevlinkCmdSbPortPoolDel", DevlinkCmdSbPortPoolDel, 21},
		{"DevlinkCmdSbTcPoolBindGet", DevlinkCmdSbTcPoolBindGet, 22},
		{"DevlinkCmdSbTcPoolBindSet", DevlinkCmdSbTcPoolBindSet, 23},
		{"DevlinkCmdSbTcPoolBindNew", DevlinkCmdSbTcPoolBindNew, 24},
		{"DevlinkCmdSbTcPoolBindDel", DevlinkCmdSbTcPoolBindDel, 25},
		{"DevlinkCmdSbOccSnapshot", DevlinkCmdSbOccSnapshot, 26},
		{"DevlinkCmdSbOccMaxClear", DevlinkCmdSbOccMaxClear, 27},
		{"DevlinkCmdEswitchGet", DevlinkCmdEswitchGet, 28},
		{"DevlinkCmdEswitchSet", DevlinkCmdEswitchSet, 29},
		{"DevlinkCmdDpipeTableGet", DevlinkCmdDpipeTableGet, 30},
		{"DevlinkCmdDpipeEntriesGet", DevlinkCmdDpipeEntriesGet, 31},
		{"DevlinkCmdDpipeHeadersGet", DevlinkCmdDpipeHeadersGet, 32},
		{"DevlinkCmdDpipeTableCountersSet", DevlinkCmdDpipeTableCountersSet, 33},
		{"DevlinkCmdResourceSet", DevlinkCmdResourceSet, 34},
		{"DevlinkCmdResourceDump", DevlinkCmdResourceDump, 35},
		{"DevlinkCmdReload", DevlinkCmdReload, 36},
		{"DevlinkCmdParamGet", DevlinkCmdParamGet, 37},
		{"DevlinkCmdParamSet", DevlinkCmdParamSet, 38},
		{"DevlinkCmdParamNew", DevlinkCmdParamNew, 39},
		{"DevlinkCmdParamDel", DevlinkCmdParamDel, 40},
		{"DevlinkCmdRegionGet", DevlinkCmdRegionGet, 41},
		{"DevlinkCmdRegionSet", DevlinkCmdRegionSet, 42},
		{"DevlinkCmdRegionNew", DevlinkCmdRegionNew, 43},
		{"DevlinkCmdRegionDel", DevlinkCmdRegionDel, 44},
		{"DevlinkCmdRegionRead", DevlinkCmdRegionRead, 45},
		{"DevlinkCmdPortParamGet", DevlinkCmdPortParamGet, 46},
		{"DevlinkCmdPortParamSet", DevlinkCmdPortParamSet, 47},
		{"DevlinkCmdPortParamNew", DevlinkCmdPortParamNew, 48},
		{"DevlinkCmdPortParamDel", DevlinkCmdPortParamDel, 49},
		{"DevlinkCmdInfoGet", DevlinkCmdInfoGet, 50},
		{"DevlinkCmdHealthReporterGet", DevlinkCmdHealthReporterGet, 51},
		{"DevlinkCmdHealthReporterSet", DevlinkCmdHealthReporterSet, 52},
		{"DevlinkCmdHealthReporterRecover", DevlinkCmdHealthReporterRecover, 53},
		{"DevlinkCmdHealthReporterDiagnose", DevlinkCmdHealthReporterDiagnose, 54},
		{"DevlinkCmdHealthReporterDumpGet", DevlinkCmdHealthReporterDumpGet, 55},
		{"DevlinkCmdHealthReporterDumpClear", DevlinkCmdHealthReporterDumpClear, 56},
		{"DevlinkCmdFlashUpdate", DevlinkCmdFlashUpdate, 57},
		{"DevlinkCmdFlashUpdateEnd", DevlinkCmdFlashUpdateEnd, 58},
		{"DevlinkCmdFlashUpdateStatus", DevlinkCmdFlashUpdateStatus, 59},
		{"DevlinkCmdTrapGet", DevlinkCmdTrapGet, 60},
		{"DevlinkCmdTrapSet", DevlinkCmdTrapSet, 61},
		{"DevlinkCmdTrapNew", DevlinkCmdTrapNew, 62},
		{"DevlinkCmdTrapDel", DevlinkCmdTrapDel, 63},
		{"DevlinkCmdTrapGroupGet", DevlinkCmdTrapGroupGet, 64},
		{"DevlinkCmdTrapGroupSet", DevlinkCmdTrapGroupSet, 65},
		{"DevlinkCmdTrapGroupNew", DevlinkCmdTrapGroupNew, 66},
		{"DevlinkCmdTrapGroupDel", DevlinkCmdTrapGroupDel, 67},
		{"DevlinkCmdTrapPolicerGet", DevlinkCmdTrapPolicerGet, 68},
		{"DevlinkCmdTrapPolicerSet", DevlinkCmdTrapPolicerSet, 69},
		{"DevlinkCmdTrapPolicerNew", DevlinkCmdTrapPolicerNew, 70},
		{"DevlinkCmdTrapPolicerDel", DevlinkCmdTrapPolicerDel, 71},
		{"DevlinkCmdHealthReporterTest", DevlinkCmdHealthReporterTest, 72},
		{"DevlinkCmdRateGet", DevlinkCmdRateGet, 73},
		{"DevlinkCmdRateSet", DevlinkCmdRateSet, 74},
		{"DevlinkCmdRateNew", DevlinkCmdRateNew, 75},
		{"DevlinkCmdRateDel", DevlinkCmdRateDel, 76},
		{"DevlinkCmdLinecardGet", DevlinkCmdLinecardGet, 77},
		{"DevlinkCmdLinecardSet", DevlinkCmdLinecardSet, 78},
		{"DevlinkCmdLinecardNew", DevlinkCmdLinecardNew, 79},
		{"DevlinkCmdLinecardDel", DevlinkCmdLinecardDel, 80},
		{"DevlinkCmdSelftestsGet", DevlinkCmdSelftestsGet, 81},
		{"DevlinkCmdSelftestsRun", DevlinkCmdSelftestsRun, 82},
		{"DevlinkCmdNotifyFilterSet", DevlinkCmdNotifyFilterSet, 83},
	}

	// Iterate over the test cases and check the values
	for _, tt := range tests {
		if tt.cmd != DevlinkCommandEnum(tt.expected) {
			t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.expected, tt.cmd)
		}
	}
}
