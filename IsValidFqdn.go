package network

import (
	"regexp"
	"strings"
)

// IsValidFqdn - reports whether host is a valid hostname that can be matched or
// matched against according to RFC 6125 2.2, with some leniency to accommodate
// legacy values.
func IsValidFqdn(s string) bool {
	// Ensure the length of the string is within the allowed range
	if len(s) < 1 || len(s) > 255 {
		return false
	}

	// Split the string into labels (trim trailing dot)
	labels := strings.Split(strings.TrimSuffix(s, "."), ".")

	// Ensure the first and last labels meet the length requirements
	if len(labels[0]) < 1 || len(labels[len(labels)-1]) < 2 || len(labels[len(labels)-1]) > 63 {
		return false
	}

	if (labels[0] == "-") || (labels[len(labels)-1] == "-") {
		return false
	}

	// Regular expression for valid label characters
	validLabel := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)

	// Check if each label is valid
	for _, label := range labels {
		label = strings.TrimSuffix(label, ".")
		if !validLabel.MatchString(label) {
			return false
		}
		if strings.Contains(label, "--") || strings.Contains(label, "..") {
			return false
		}
	}
	return true
}
