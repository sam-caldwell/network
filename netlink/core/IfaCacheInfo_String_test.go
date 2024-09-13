package core

import (
	"fmt"
	"golang.org/x/sys/unix"
	"testing"
)

// Test for the String() method of IfaCacheInfo.
func TestIfaCacheInfoString(t *testing.T) {
	// Test case with predefined values.
	t.Run("String representation", func(t *testing.T) {
		msg := IfaCacheInfo{
			IfaCacheinfo: unix.IfaCacheinfo{
				Valid:    3600,
				Prefered: 1800,
				Cstamp:   100,
				Tstamp:   200,
			},
		}

		// Expected string format.
		expected := fmt.Sprintf("Valid: %d, Prefered: %d, Created: %d, Updated: %d",
			msg.Valid, msg.Prefered, msg.Cstamp, msg.Tstamp)

		result := msg.String()
		if result != expected {
			t.Errorf("Expected: %s, got: %s", expected, result)
		}
	})
}
