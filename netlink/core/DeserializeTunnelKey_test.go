package core

import (
	"testing"
)

// TestDeserializeTunnelKey tests the DeserializeTunnelKey function.
func TestDeserializeTunnelKey(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		// Prepare a byte slice for TcTunnelKey
		data := make([]byte, SizeOfTcTunnelKey)

		// Fill the byte slice with test values
		NativeEndian.PutUint32(data[0:4], 0x01020304)   // Index
		NativeEndian.PutUint32(data[4:8], 0x05060708)   // Capab
		NativeEndian.PutUint32(data[8:12], 0x090A0B0C)  // Action
		NativeEndian.PutUint32(data[12:16], 0x0D0E0F10) // Refcnt
		NativeEndian.PutUint32(data[16:20], 0x11121314) // Bindcnt

		// Call DeserializeTunnelKey
		result, err := DeserializeTunnelKey(data)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Verify the result
		if result.Index != 0x01020304 {
			t.Errorf("Expected Index 0x01020304, got 0x%08x", result.Index)
		}
		if result.Capab != 0x05060708 {
			t.Errorf("Expected Capab 0x05060708, got 0x%08x", result.Capab)
		}
		if result.Action != 0x090A0B0C {
			t.Errorf("Expected Action 0x090A0B0C, got 0x%08x", result.Action)
		}
		if result.Refcnt != 0x0D0E0F10 {
			t.Errorf("Expected Refcnt 0x0D0E0F10, got 0x%08x", result.Refcnt)
		}
		if result.Bindcnt != 0x11121314 {
			t.Errorf("Expected Bindcnt 0x11121314, got 0x%08x", result.Bindcnt)
		}
	})

	t.Run("tests DeserializeTunnelKey with short input.", func(t *testing.T) {
		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcTunnelKey-1)

		// Call DeserializeTunnelKey and expect an error
		_, err := DeserializeTunnelKey(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		expectedErr := "input too short"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
