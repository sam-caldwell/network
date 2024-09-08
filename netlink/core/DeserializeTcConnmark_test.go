package core

import (
	"testing"
)

func TestDeserializeTcConnmark(t *testing.T) {
	//
	// type TcGen struct {
	//    Index   uint32
	//    Capab   uint32
	//    Action  int32
	//    Refcnt  int32
	//    Bindcnt int32
	// }
	//
	// type TcConnmark struct {
	//    TcGen
	//    Zone uint16
	// }
	//
	validTcConnMark := []byte{
		0x01, 0x00, 0x00, 0x00, // Index (uint32)
		0x02, 0x00, 0x00, 0x00, // Capab (uint32)
		0x03, 0x00, 0x00, 0x00, // Action (int32)
		0x04, 0x00, 0x00, 0x00, // Refcnt (int32)
		0x05, 0x00, 0x00, 0x00, // Bindcnt (int32)
		0x0A, 0x00, 0x00, 0x00, // Zone (uint16) + padding
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Call DeserializeTcConnmark
		tcConnmark, err := DeserializeTcConnmark(validTcConnMark)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got '%v' (size: %d/%d)",
				err, len(validTcConnMark), SizeOfTcConnmark)
		}

		// Verify the deserialized values
		if tcConnmark.Index != 1 {
			t.Errorf("Expected Index to be 1, but got %d", tcConnmark.Index)
		}
		if tcConnmark.Capab != 2 {
			t.Errorf("Expected Capab to be 2, but got %d", tcConnmark.Capab)
		}
		if tcConnmark.Action != 3 {
			t.Errorf("Expected Action to be 3, but got %d", tcConnmark.Action)
		}
		if tcConnmark.Refcnt != 4 {
			t.Errorf("Expected Refcnt to be 4, but got %d", tcConnmark.Refcnt)
		}
		if tcConnmark.Bindcnt != 5 {
			t.Errorf("Expected Bindcnt to be 5, but got %d", tcConnmark.Bindcnt)
		}
		if tcConnmark.Zone != 10 {
			t.Errorf("Expected Zone to be 10, but got %d", tcConnmark.Zone)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		buf := make([]byte, SizeOfTcConnmark-2) // Too short

		// Call DeserializeTcConnmark
		tcConnmark, err := DeserializeTcConnmark(buf)

		// Verify the function returns an error
		if err == nil {
			t.Fatalf("Expected error due to short input, but got none")
		}

		// Verify that tcConnmark is nil
		if tcConnmark != nil {
			t.Errorf("Expected tcConnmark to be nil due to short input, but got %v", tcConnmark)
		}
	})
}
