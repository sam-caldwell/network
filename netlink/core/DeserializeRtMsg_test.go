package core

import (
	"testing"
)

func TestDeserializeRtMsg(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Prepare a valid byte slice (12 bytes)
		buf := []byte{
			0x02,                   // Family
			0x10,                   // Dst_len
			0x20,                   // Src_len
			0x30,                   // Tos
			0x40,                   // Table
			0x50,                   // Protocol
			0x60,                   // Scope
			0x70,                   // Type
			0x01, 0x02, 0x03, 0x04, // Flags (4 bytes, little-endian format)
		}

		// Call DeserializeRtMsg
		rtMsg := DeserializeRtMsg(buf)

		// Check that the struct is not nil
		if rtMsg == nil {
			t.Fatalf("Expected rtMsg to be non-nil, but got nil")
		}

		// Check the deserialized values
		if rtMsg.Family != 0x02 {
			t.Errorf("Expected Family to be 0x02, but got %d", rtMsg.Family)
		}
		if rtMsg.Dst_len != 0x10 {
			t.Errorf("Expected Dst_len to be 0x10, but got %d", rtMsg.Dst_len)
		}
		if rtMsg.Src_len != 0x20 {
			t.Errorf("Expected Src_len to be 0x20, but got %d", rtMsg.Src_len)
		}
		if rtMsg.Tos != 0x30 {
			t.Errorf("Expected Tos to be 0x30, but got %d", rtMsg.Tos)
		}
		if rtMsg.Table != 0x40 {
			t.Errorf("Expected Table to be 0x40, but got %d", rtMsg.Table)
		}
		if rtMsg.Protocol != 0x50 {
			t.Errorf("Expected Protocol to be 0x50, but got %d", rtMsg.Protocol)
		}
		if rtMsg.Scope != 0x60 {
			t.Errorf("Expected Scope to be 0x60, but got %d", rtMsg.Scope)
		}
		if rtMsg.Type != 0x70 {
			t.Errorf("Expected Type to be 0x70, but got %d", rtMsg.Type)
		}

		// Check the deserialized Flags value (little-endian conversion)
		expectedFlags := uint32(0x04030201)
		if rtMsg.Flags != expectedFlags {
			t.Errorf("Expected Flags to be 0x%08x, but got 0x%08x", expectedFlags, rtMsg.Flags)
		}
	})

	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than 12 bytes)
		buf := make([]byte, 8) // Too short

		// Call DeserializeRtMsg
		rtMsg := DeserializeRtMsg(buf)

		// Check that the function returns nil for insufficient byte length
		if rtMsg != nil {
			t.Errorf("Expected rtMsg to be nil due to short input, but got %v", rtMsg)
		}
	})
}
