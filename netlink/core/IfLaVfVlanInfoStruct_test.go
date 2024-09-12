package core

import (
	"testing"
	"unsafe"
)

// TestIfLaVfVlanInfoStruct tests the IfLaVfVlanInfoStruct struct and its fields.
func TestIfLaVfVlanInfoStruct(t *testing.T) {
	t.Run("verify size", func(t *testing.T) {
		const expectedSizeInBytes = 16
		if sz := int(unsafe.Sizeof(IfLaVfVlanInfoStruct{})); sz != expectedSizeInBytes {
			t.Fatalf("size mismatch")
		}
	})
	t.Run("field check", func(t *testing.T) {
		// Initialize a test instance of IfLaVfVlanInfoStruct
		_ = IfLaVfVlanInfoStruct{
			VfVlan: VfVlan{
				Vf:   uint32(0),
				Vlan: uint32(0),
				Qos:  uint32(0),
			},
			VlanProto: IfLaVfVlanInfoEnum(0),
		}

	})
}
