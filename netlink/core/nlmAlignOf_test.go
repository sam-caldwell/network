package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestNlmAlignOf(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, unix.NLMSG_ALIGNTO},
		{2, unix.NLMSG_ALIGNTO},
		{unix.NLMSG_ALIGNTO - 1, unix.NLMSG_ALIGNTO},
		{unix.NLMSG_ALIGNTO, unix.NLMSG_ALIGNTO},
		{unix.NLMSG_ALIGNTO + 1, 2 * unix.NLMSG_ALIGNTO},
		{17, 20}, // assuming NLMSG_ALIGNTO is 4, 17 should align to 20
	}

	for _, tt := range tests {
		t.Run("Aligning message length", func(t *testing.T) {
			result := nlmAlignOf(tt.input)
			if result != tt.expected {
				t.Errorf("nlmAlignOf(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
