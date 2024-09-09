package core

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestTcU32SelSerialize(t *testing.T) {
	// Create a sample TcU32Sel object with keys.
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

	// Call the Serialize method
	serialized, err := msg.Serialize()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Expected serialized output
	expected := make([]byte, SizeOfTcU32Sel+2*SizeOfTcU32Key)

	// Expected values for the fixed portion
	expected[0] = msg.Flags
	expected[1] = msg.Offshift
	expected[2] = msg.Nkeys
	expected[3] = msg.Pad
	binary.BigEndian.PutUint16(expected[4:6], msg.Offmask)
	NativeEndian.PutUint16(expected[6:8], msg.Off)
	NativeEndian.PutUint16(expected[8:10], uint16(msg.Offoff))
	NativeEndian.PutUint16(expected[10:12], uint16(msg.Hoff))
	binary.BigEndian.PutUint32(expected[12:16], msg.Hmask)

	// Serialize each key and append to the expected output
	offset := 16
	for _, key := range msg.Keys {
		keyData, err := key.Serialize()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		copy(expected[offset:], keyData) // Ensures correct copying of key data
		offset += SizeOfTcU32Key
	}

	// Check if the serialized output matches the expected output
	if !reflect.DeepEqual(serialized[:len(expected)], expected) {
		t.Fatalf("Serialized output doesn't match expected.\n"+
			"Expected: %v\n"+
			"     Got: %v", expected, serialized[:len(expected)])
	}
}
