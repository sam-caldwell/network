package core

import (
	"testing"
)

func TestDeserializeTcPedit(t *testing.T) {
	// Helper function to create a valid byte slice for TcPeditSel and TcPeditKey
	createValidTcPeditBytes := func(nKeys uint8) []byte {
		// Create a byte slice manually, representing TcPeditSel and TcPeditKey fields explicitly.
		// TcPeditSel (TcGen: Index, Capab, Action, Refcnt, Bindcnt, NKeys, Flags)
		data := []byte{
			0x00, 0x00, 0x00, 0x01, // Index (uint32)
			0x00, 0x00, 0x00, 0x02, // Capab (uint32)
			0x00, 0x00, 0x00, 0x03, // Action (int32)
			0x00, 0x00, 0x00, 0x04, // Refcnt (int32)
			0x00, 0x00, 0x00, 0x05, // Bindcnt (int32)
			nKeys,            // NKeys (uint8)
			0x00, 0x00, 0x00, // Padding for Flags (uint8) + 2 padding bytes
		}

		// TcPeditKey (Mask, Val, Off, At, OffMask, Shift) for each key
		for i := uint8(0); i < nKeys; i++ {
			data = append(data,
				0x00, 0x00, 0x00, 0x01, // Mask (uint32)
				0x00, 0x00, 0x00, 0x02, // Val (uint32)
				0x00, 0x00, 0x00, 0x03, // Off (uint32)
				0x00, 0x00, 0x00, 0x04, // At (uint32)
				0x00, 0x00, 0x00, 0x05, // OffMask (uint32)
				0x00, 0x00, 0x00, 0x06, // Shift (uint32)
			)
		}

		return data
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice with 2 keys
		validBytes := createValidTcPeditBytes(2)

		// Call DeserializeTcPedit
		tcPeditSel, tcPeditKeys, err := DeserializeTcPedit(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify TcPeditSel values
		if tcPeditSel.TcGen.Index != 1 {
			t.Errorf("Expected TcGen.Index to be 1, but got %d", tcPeditSel.TcGen.Index)
		}
		if tcPeditSel.NKeys != 2 {
			t.Errorf("Expected NKeys to be 2, but got %d", tcPeditSel.NKeys)
		}

		// Verify TcPeditKeys
		if len(tcPeditKeys) != 2 {
			t.Fatalf("Expected 2 keys, but got %d", len(tcPeditKeys))
		}

		for i, key := range tcPeditKeys {
			if key.Mask != 1 {
				t.Errorf("Expected Mask for key %d to be 1, but got %d", i, key.Mask)
			}
			if key.Val != 2 {
				t.Errorf("Expected Val for key %d to be 2, but got %d", i, key.Val)
			}
			if key.Off != 3 {
				t.Errorf("Expected Off for key %d to be 3, but got %d", i, key.Off)
			}
			if key.At != 4 {
				t.Errorf("Expected At for key %d to be 4, but got %d", i, key.At)
			}
			if key.OffMask != 5 {
				t.Errorf("Expected OffMask for key %d to be 5, but got %d", i, key.OffMask)
			}
			if key.Shift != 6 {
				t.Errorf("Expected Shift for key %d to be 6, but got %d", i, key.Shift)
			}
		}
	})

	// Test case for invalid input (too short for TcPeditSel)
	t.Run("input too short for TcPeditSel", func(t *testing.T) {
		// Prepare an invalid byte slice (too small for TcPeditSel)
		invalidBytes := make([]byte, SizeOfTcPeditSel-1)

		// Call DeserializeTcPedit
		tcPeditSel, tcPeditKeys, err := DeserializeTcPedit(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected error due to short input, but got none")
		}

		// Verify that tcPeditSel and tcPeditKeys are nil
		if tcPeditSel != nil || tcPeditKeys != nil {
			t.Errorf("Expected nil tcPeditSel and tcPeditKeys, but got %v and %v", tcPeditSel, tcPeditKeys)
		}
	})

	// Test case for invalid input (too short for keys)
	t.Run("input too short for keys", func(t *testing.T) {
		// Create a byte slice for TcPeditSel with 2 keys, but not enough space for 2 full keys
		invalidBytes := createValidTcPeditBytes(2)[:SizeOfTcPeditSel+SizeOfTcPeditKey-1]

		// Call DeserializeTcPedit
		tcPeditSel, tcPeditKeys, err := DeserializeTcPedit(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected error due to short input for keys, but got none")
		}

		// Verify that tcPeditSel and tcPeditKeys are nil
		if tcPeditSel != nil || tcPeditKeys != nil {
			t.Errorf("Expected nil tcPeditSel and tcPeditKeys, but got %v and %v", tcPeditSel, tcPeditKeys)
		}
	})
}
