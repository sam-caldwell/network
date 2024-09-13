package core

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
)

func TestIfInfoMsgSerialize(t *testing.T) {
	tests := []struct {
		name         string
		ifInfoMsg    IfInfoMsg
		expectedData []byte
	}{
		{
			name: "Valid IfInfoMsg Serialization",
			ifInfoMsg: IfInfoMsg{
				IfInfomsg: unix.IfInfomsg{
					Family: unix.AF_INET,
					Type:   1,
					Index:  2,
					Flags:  3,
					Change: 4,
				},
			},
			expectedData: func() []byte {
				buf := new(bytes.Buffer)
				_ = binary.Write(buf, NativeEndian, uint8(unix.AF_INET)) // Family
				_ = binary.Write(buf, NativeEndian, uint8(0))            // Padding byte
				_ = binary.Write(buf, NativeEndian, uint16(1))           // Type
				_ = binary.Write(buf, NativeEndian, int32(2))            // Index
				_ = binary.Write(buf, NativeEndian, uint32(3))           // Flags
				_ = binary.Write(buf, NativeEndian, uint32(4))           // Change
				return buf.Bytes()
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Serialize the IfInfoMsg
			serializedData, err := tt.ifInfoMsg.Serialize()
			if err != nil {
				t.Fatalf("Unexpected error during serialization: %v", err)
			}

			// Compare the serialized data with the expected byte slice
			if !bytes.Equal(serializedData, tt.expectedData) {
				t.Errorf("Serialized data does not match expected\nExpected: %v\nGot: %v", tt.expectedData, serializedData)
			}
		})
	}
}
