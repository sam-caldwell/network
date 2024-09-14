package core

import (
	"testing"
)

func TestNlaTypeEnum(t *testing.T) {
	tests := []struct {
		name   string
		actual NlaType
		expect NlaType
	}{
		{name: "NlaUnspec", actual: NlaType(0), expect: NlaUnspec},
		{name: "NlaU8", actual: NlaType(1), expect: NlaU8},
		{name: "NlaU16", actual: NlaType(2), expect: NlaU16},
		{name: "NlaU32", actual: NlaType(3), expect: NlaU32},
		{name: "NlaU64", actual: NlaType(4), expect: NlaU64},
		{name: "NlaString", actual: NlaType(5), expect: NlaString},
		{name: "NlaFlag", actual: NlaType(6), expect: NlaFlag},
		{name: "NlaMsecs", actual: NlaType(7), expect: NlaMsecs},
		{name: "NlaNested", actual: NlaType(8), expect: NlaNested},
		{name: "NlaNestedArray", actual: NlaType(9), expect: NlaNestedArray},
		{name: "NlaNulString", actual: NlaType(10), expect: NlaNulString},
		{name: "NlaBinary", actual: NlaType(11), expect: NlaBinary},
		{name: "NlaS8", actual: NlaType(12), expect: NlaS8},
		{name: "NlaS16", actual: NlaType(13), expect: NlaS16},
		{name: "NlaS32", actual: NlaType(14), expect: NlaS32},
		{name: "NlaS64", actual: NlaType(15), expect: NlaS64},
		{name: "NlaBitfield32", actual: NlaType(16), expect: NlaBitfield32},
		{name: "NlaReject", actual: NlaType(17), expect: NlaReject},
		{name: "NlaBe16", actual: NlaType(18), expect: NlaBe16},
		{name: "NlaBe32", actual: NlaType(19), expect: NlaBe32},
		{name: "NlaSint", actual: NlaType(20), expect: NlaSint},
		{name: "NlaUint", actual: NlaType(21), expect: NlaUint},
		{name: "NlaTypeMax", actual: NlaType(NlaUint), expect: NlaTypeMax},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expect {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.actual, tt.expect)
			}
		})
	}
}
