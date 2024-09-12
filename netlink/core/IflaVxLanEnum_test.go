package core

import (
	"testing"
)

// TestIflaVxLanEnum tests the values of IflaVxLanEnum constants.
func TestIflaVxLanEnum(t *testing.T) {
	tests := []struct {
		name     string
		value    IflaVxLanEnum
		expected IflaVxLanEnum
	}{
		{"IflaVxlanUnspec", IflaVxlanUnspec, 0},
		{"IflaVxlanId", IflaVxlanId, 1},
		{"IflaVxlanGroup", IflaVxlanGroup, 2},
		{"IflaVxlanLink", IflaVxlanLink, 3},
		{"IflaVxlanLocal", IflaVxlanLocal, 4},
		{"IflaVxlanTtl", IflaVxlanTtl, 5},
		{"IflaVxlanTos", IflaVxlanTos, 6},
		{"IflaVxlanLearning", IflaVxlanLearning, 7},
		{"IflaVxlanAgeing", IflaVxlanAgeing, 8},
		{"IflaVxlanLimit", IflaVxlanLimit, 9},
		{"IflaVxlanPortRange", IflaVxlanPortRange, 10},
		{"IflaVxlanProxy", IflaVxlanProxy, 11},
		{"IflaVxlanRsc", IflaVxlanRsc, 12},
		{"IflaVxlanL2miss", IflaVxlanL2miss, 13},
		{"IflaVxlanL3miss", IflaVxlanL3miss, 14},
		{"IflaVxlanPort", IflaVxlanPort, 15},
		{"IflaVxlanGroup6", IflaVxlanGroup6, 16},
		{"IflaVxlanLocal6", IflaVxlanLocal6, 17},
		{"IflaVxlanUdpCsum", IflaVxlanUdpCsum, 18},
		{"IflaVxlanUdpZeroCsum6Tx", IflaVxlanUdpZeroCsum6Tx, 19},
		{"IflaVxlanUdpZeroCsum6Rx", IflaVxlanUdpZeroCsum6Rx, 20},
		{"IflaVxlanRemcsumTx", IflaVxlanRemcsumTx, 21},
		{"IflaVxlanRemcsumRx", IflaVxlanRemcsumRx, 22},
		{"IflaVxlanGbp", IflaVxlanGbp, 23},
		{"IflaVxlanRemcsumNopartial", IflaVxlanRemcsumNopartial, 24},
		{"IflaVxlanCollectMetadata", IflaVxlanCollectMetadata, 25},
		{"IflaVxlanLabel", IflaVxlanLabel, 26},
		{"IflaVxlanGpe", IflaVxlanGpe, 27},
		{"IflaVxlanTtlInherit", IflaVxlanTtlInherit, 28},
		{"IflaVxlanDf", IflaVxlanDf, 29},
		{"IflaVxlanVnifilter", IflaVxlanVnifilter, 30},
		{"IflaVxlanLocalbypass", IflaVxlanLocalbypass, 31},
		{"IflaVxlanLabelPolicy", IflaVxlanLabelPolicy, 32},
		{"IflaVxlanMax", IflaVxlanMax, 32}, // Max should match the highest value.
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.value != test.expected {
				t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
			}
		})
	}
}
