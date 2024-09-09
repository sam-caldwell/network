package core

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
)

// TestDeserializeTcU32Sel tests the DeserializeTcU32Sel function.
func TestDeserializeTcU32Sel(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Create a byte slice that represents a serialized TcU32Sel object
		expectedSel := TcU32Sel{
			Flags:    0x01,
			Offshift: 0x02,
			Nkeys:    2,
			Pad:      0x00,
			Offmask:  0x0405,
			Off:      0x0506,
			Offoff:   0x0708,
			Hoff:     0x090A,
			Hmask:    0x0B0C0D0E,
			Keys: []TcU32Key{
				{
					Mask:    0x01020304,
					Val:     0x05060708,
					Off:     0x090A0B0C,
					OffMask: 0x0D0E0F10,
				},
				{
					Mask:    0x11121314,
					Val:     0x15161718,
					Off:     0x191A1B1C,
					OffMask: 0x1D1E1F20,
				},
			},
		}

		// Serialize the expected object into a byte slice
		var buf bytes.Buffer

		// Serialize the fixed portion
		buf.WriteByte(expectedSel.Flags)
		buf.WriteByte(expectedSel.Offshift)
		buf.WriteByte(expectedSel.Nkeys)
		buf.WriteByte(expectedSel.Pad)
		_ = binary.Write(&buf, binary.BigEndian, expectedSel.Offmask)
		_ = binary.Write(&buf, NativeEndian, expectedSel.Off)
		_ = binary.Write(&buf, NativeEndian, expectedSel.Offoff)
		_ = binary.Write(&buf, NativeEndian, expectedSel.Hoff)
		_ = binary.Write(&buf, binary.BigEndian, expectedSel.Hmask)

		// Serialize the keys
		for _, key := range expectedSel.Keys {
			keyData, _ := key.Serialize()
			buf.Write(keyData)
		}

		// Deserialize the byte slice back into a TcU32Sel object
		serializedData := buf.Bytes()
		deserializedSel, err := DeserializeTcU32Sel(serializedData)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Check if the deserialized object matches the expected object
		if !reflect.DeepEqual(deserializedSel, &expectedSel) {
			t.Fatalf("Deserialized object doesn't match expected.\nExpected: %+v\nGot: %+v", expectedSel, deserializedSel)
		}
	})
	t.Run("short input", func(t *testing.T) {
		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcU32Sel-1)

		// Call DeserializeTcU32Sel
		_, err := DeserializeTcU32Sel(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		expectedErr := "input too short"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
	t.Run("serialize-deserialize agreement", func(t *testing.T) {
		msg := TcU32Sel{
			Flags:    0x01,
			Offshift: 0x02,
			Nkeys:    2,
			Pad:      0x00,
			Offmask:  0x0405,
			Off:      0x0506,
			Offoff:   0x0708,
			Hoff:     0x090A,
			Hmask:    0x0B0C0D0E,
			Keys: []TcU32Key{
				{
					Mask:    0x01020304,
					Val:     0x05060708,
					Off:     0x090A0B0C,
					OffMask: 0x0D0E0F10,
				},
				{
					Mask:    0x11121314,
					Val:     0x15161718,
					Off:     0x191A1B1C,
					OffMask: 0x1D1E1F20,
				},
			},
		}

		serializedOutput, err := msg.Serialize()
		if err != nil {
			t.Fatalf("Unexpected error during serialization: %v", err)
		}

		deserializedOutput, err := DeserializeTcU32Sel(serializedOutput)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Manually compare fields if needed
		if !reflect.DeepEqual(&msg, deserializedOutput) {
			t.Fatalf("Serialized/Deserialized object mismatch.\nExpected: %+v\nGot: %+v", msg, *deserializedOutput)
		}
	})
}
