//go:build linux

package core

import (
	"testing"
)

func TestError_constants(t *testing.T) {
	testData := []struct {
		actual string
		expect string
	}{
		{actual: ErrNilInput, expect: "nil input"},
		{actual: ErrInputTooShort, expect: "input too short"},
		{actual: ErrInputTooLarge, expect: "input too large"},
		{actual: ErrInvalidAttributeLength, expect: "invalid attribute length"},
		{actual: ErrTruncatedAttribute, expect: "truncated attribute"},
	}
	for _, tt := range testData {
		if tt.actual != tt.expect {
			t.Fatalf("error mismatch. actual:%v, expected:%v", tt.actual, tt.expect)
		}
	}
}
