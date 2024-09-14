package core

import (
	"testing"
)

func TestNlmsgerrAttrValues(t *testing.T) {
	tests := []struct {
		name     string
		attr     NlmsgerrAttr
		expected uint16
	}{
		{"NlmsgerrAttrUnused", NlmsgerrAttrUnused, 0},
		{"NlmsgerrAttrMsg", NlmsgerrAttrMsg, 1},
		{"NlmsgerrAttrOffs", NlmsgerrAttrOffs, 2},
		{"NlmsgerrAttrCookie", NlmsgerrAttrCookie, 3},
		{"NlmsgerrAttrPolicy", NlmsgerrAttrPolicy, 4},
		{"NlmsgerrAttrMissType", NlmsgerrAttrMissType, 5},
		{"NlmsgerrAttrMissNest", NlmsgerrAttrMissNest, 6},
		{"NlmsgerrAttrMax", NlmsgerrAttrMax, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint16(tt.attr) != tt.expected {
				t.Errorf("Expected %s to be %d, but got %d", tt.name, tt.expected, tt.attr)
			}
		})
	}
}
