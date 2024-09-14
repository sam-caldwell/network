package core

import (
	"testing"
)

func TestNlaTypeEnum(t *testing.T) {
	tests := []struct {
		name     string
		expected NlaType
	}{
		{"NlaUnspec", NlaUnspec},
		{"NlaU8", NlaU8},
		{"NlaU16", NlaU16},
		{"NlaU32", NlaU32},
		{"NlaU64", NlaU64},
		{"NlaString", NlaString},
		{"NlaFlag", NlaFlag},
		{"NlaMsecs", NlaMsecs},
		{"NlaNested", NlaNested},
		{"NlaNestedArray", NlaNestedArray},
		{"NlaNulString", NlaNulString},
		{"NlaBinary", NlaBinary},
		{"NlaS8", NlaS8},
		{"NlaS16", NlaS16},
		{"NlaS32", NlaS32},
		{"NlaS64", NlaS64},
		{"NlaBitfield32", NlaBitfield32},
		{"NlaReject", NlaReject},
		{"NlaBe16", NlaBe16},
		{"NlaBe32", NlaBe32},
		{"NlaSint", NlaSint},
		{"NlaUint", NlaUint},
		{"NlaTypeMax", NlaTypeMax},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if NlaType(i) != tt.expected {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, i, tt.expected)
			}
		})
	}
}
