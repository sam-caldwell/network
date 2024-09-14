package core

import (
	"testing"
)

func TestNlaPolicy(t *testing.T) {
	t.Run("Test NlaPolicy struct initialization and fields", func(t *testing.T) {
		// Define sample values for testing
		expectedType := NlaType(1)                   // Replace with a valid NlaType value
		expectedLen := 4                             // Example length
		expectedStrictStartType := IflaGeneveEnum(2) // Replace with a valid IflaGeneveEnum value

		// Initialize NlaPolicy
		policy := NlaPolicy{
			Type:            expectedType,
			Len:             expectedLen,
			StrictStartType: expectedStrictStartType,
		}

		// Validate Type
		if policy.Type != expectedType {
			t.Errorf("Expected Type to be %v, got %v", expectedType, policy.Type)
		}

		// Validate Len
		if policy.Len != expectedLen {
			t.Errorf("Expected Len to be %d, got %d", expectedLen, policy.Len)
		}

		// Validate StrictStartType
		if policy.StrictStartType != expectedStrictStartType {
			t.Errorf("Expected StrictStartType to be %v, got %v", expectedStrictStartType, policy.StrictStartType)
		}
	})
}
