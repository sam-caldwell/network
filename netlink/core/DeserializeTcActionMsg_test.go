package core

import (
	"bytes"
	"testing"
)

func TestDeserializeTcActionMsg(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Prepare a valid byte slice (size equal to SizeOfTcActionMsg)
		//
		// type TcActionMsg struct {
		//    Family uint8
		//    Pad    [3]byte
		//}
		//
		buf := []byte{
			0x01,             // Family
			0x02, 0x00, 0x00, // Pad
		}

		tcActionMsg, err := DeserializeTcActionMsg(buf)
		if err != nil {
			t.Fatal(err)
		}

		if tcActionMsg.Family != 1 {
			t.Errorf("Expected Family to be 1, but got %d", tcActionMsg.Family)
		}
		if expected := []byte{0x02, 0x00, 0x00}; !bytes.Equal(tcActionMsg.Pad[:], expected) {
			t.Errorf("Expected Pad to be %v but got %v", expected, tcActionMsg.Pad)
		}
	})

	t.Run("input too short", func(t *testing.T) {
		buf := make([]byte, SizeOfTcActionMsg-2) // Too short
		v, err := DeserializeTcActionMsg(buf)
		if err == nil {
			t.Fatal("expected error not found")
		}
		if v != nil {
			t.Fatal("expect error")
		}
		if err.Error() != "input too short" {
			t.Fatalf("expected 'input too short' but got %v", err.Error())
		}
	})
}
