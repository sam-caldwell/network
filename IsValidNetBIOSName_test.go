package network

import (
	"testing"
)

func TestIsValidNetBIOSName(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"ValidName1", "NETBIOSNAME", true},
		{"ValidName2", "NETBIOS1234", true},
		{"ValidName3", "NET_BIOS-12", true},
		{"ValidName4", "A$VALIDNAME", true},
		{"InvalidNameTooLong", "THISNAMEISTOOLONGFORNETBIOS", false},
		{"InvalidNameTooShort", "", false},
		{"InvalidNameInvalidChars", "Invalid@Name", false},
		{"InvalidNameStartsWithDigit", "1InvalidName", false},
		{"InvalidNameSpecialChar", "-InvalidName", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidNetBIOSName(test.input)
			if result != test.output {
				t.Errorf("IsValidNetBIOSName(%q) = %v; want %v", test.input, result, test.output)
			}
		})
	}
}
