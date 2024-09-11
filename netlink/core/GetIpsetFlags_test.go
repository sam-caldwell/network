package core

import (
	"testing"

	"golang.org/x/sys/unix"
)

// TestGetIpsetFlags tests the GetIpsetFlags function
func TestGetIpsetFlags(t *testing.T) {
	tests := []struct {
		name     string
		cmd      IpSetCmdEnum
		expected int
	}{
		{
			name:     "IpsetCmdCreate",
			cmd:      IpsetCmdCreate,
			expected: unix.NLM_F_REQUEST | unix.NLM_F_ACK | unix.NLM_F_CREATE,
		},
		{
			name:     "IpsetCmdDestroy",
			cmd:      IpsetCmdDestroy,
			expected: unix.NLM_F_REQUEST | unix.NLM_F_ACK,
		},
		{
			name:     "IpsetCmdList",
			cmd:      IpsetCmdList,
			expected: unix.NLM_F_REQUEST | unix.NLM_F_ACK | unix.NLM_F_ROOT | unix.NLM_F_MATCH | unix.NLM_F_DUMP,
		},
		{
			name:     "IpsetCmdAdd",
			cmd:      IpsetCmdAdd,
			expected: unix.NLM_F_REQUEST | unix.NLM_F_ACK,
		},
		{
			name:     "IpsetCmdHeader",
			cmd:      IpsetCmdHeader,
			expected: unix.NLM_F_REQUEST,
		},
		{
			name:     "Unknown Command",
			cmd:      IpSetCmdEnum(255), // Invalid command
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetIpsetFlags(tt.cmd)
			if result != tt.expected {
				t.Errorf("GetIpsetFlags(%v) = %d; want %d", tt.cmd, result, tt.expected)
			}
		})
	}
}
