package core

// SEG6EncapModeString - Converts an integer representing the SEG6 encapsulation mode to a string.
// This function provides a human-readable description of the SEG6 encapsulation mode.
//
// Parameters:
//   - mode: The integer value representing the encapsulation mode.
//
// Returns:
//   - string: A string describing the mode ("inline", "encap", or "unknown").
func SEG6EncapModeString(mode int) string {
	switch mode {
	case SEG6_IPTUN_MODE_INLINE:
		return "inline"
	case SEG6_IPTUN_MODE_ENCAP:
		return "encap"
	}
	return "unknown"
}
