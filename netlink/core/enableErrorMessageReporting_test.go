package core

import (
	"testing"
)

// TestErrorMessageReporting tests the enabling and disabling of error message reporting
func TestErrorMessageReporting(t *testing.T) {
	// Initially, error message reporting should be disabled by default
	if enableErrorMessageReporting != false {
		t.Errorf("Expected enableErrorMessageReporting to be false, got %v", enableErrorMessageReporting)
	}

	// Enable error message reporting
	EnableErrorMessageReporting()
	if enableErrorMessageReporting != true {
		t.Errorf("Expected enableErrorMessageReporting to be true after enabling, got %v", enableErrorMessageReporting)
	}

	// Disable error message reporting
	DisableErrorMessageReporting()
	if enableErrorMessageReporting != false {
		t.Errorf("Expected enableErrorMessageReporting to be false after disabling, got %v", enableErrorMessageReporting)
	}
}
