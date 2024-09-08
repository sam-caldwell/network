package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDeserializeTcPedit(t *testing.T) {
	t.Skip("disabled")
	// Helper function to create a valid byte slice for TcPeditSel and TcPeditKey
	createValidTcPeditBytes := func(nKeys uint8) []byte {
		buf := new(bytes.Buffer)

		// TcPeditSel: TcGen (Index, Capab, Action, Refcnt, Bindcnt), NKeys, Flags
		_ = binary.Write(buf, binary.BigEndian, uint32(1)) // TcGen.Index
		_ = binary.Write(buf, binary.BigEndian, uint32(2)) // TcGen.Capab
		_ = binary.Write(buf, binary.BigEndian, int32(3))  // TcGen.Action
		_ = binary.Write(buf, binary.BigEndian, int32(4))  // TcGen.Refcnt
		_ = binary.Write(buf, binary.BigEndian, int32(5))  // TcGen.Bindcnt
		_ = binary.Write(buf, binary.BigEndian, nKeys)     // NKeys
		_ = binary.Write(buf, binary.BigEndian, uint8(0))  // Flags

		// TcPeditKey: Mask, Val, Off, At, OffMask, Shift for each key
		for i := uint8(0); i < nKeys; i++ {
			_ = binary.Write(buf, binary.BigEndian, uint32(1)) // Mask
			_ = binary.Write(buf, binary.BigEndian, uint32(2)) // Val
			_ = binary.Write(buf, binary.BigEndian, uint32(3)) // Off
			_ = binary.Write(buf, binary.BigEndian, uint32(4)) // At
			_ = binary.Write(buf, binary.BigEndian, uint32(5)) // OffMask
			_ = binary.Write(buf, binary.BigEndian, uint32(6)) // Shift
		}

		return buf.Bytes()
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
