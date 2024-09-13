package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestBridgeVlanInfo_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		input   BridgeVlanInfo
		want    []byte
		wantErr bool
	}{
		{
			name: "Valid BridgeVlanInfo",
			input: BridgeVlanInfo{
				Flags: 0x0100, // Example flag value
				Vid:   100,    // Example VLAN ID (100)
			},
			// Expected result is little-endian representation of Flags and Vid.
			want: func() []byte {
				buf := new(bytes.Buffer)
				_ = binary.Write(buf, NativeEndian, uint16(0x0100)) // Serialize Flags
				_ = binary.Write(buf, NativeEndian, uint16(100))    // Serialize Vid
				return buf.Bytes()
			}(),
			wantErr: false,
		},
		{
			name: "Zero BridgeVlanInfo",
			input: BridgeVlanInfo{
				Flags: 0x0000, // No flags
				Vid:   0,      // VLAN ID = 0 (undefined VLAN ID)
			},
			// Expected result for zero fields.
			want: func() []byte {
				buf := new(bytes.Buffer)
				_ = binary.Write(buf, NativeEndian, uint16(0x0000)) // Serialize Flags
				_ = binary.Write(buf, NativeEndian, uint16(0))      // Serialize Vid
				return buf.Bytes()
			}(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.Serialize()

			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !bytes.Equal(got, tt.want) {
				t.Errorf("Serialize() got = %v, want %v", got, tt.want)
			}
		})
	}
}
