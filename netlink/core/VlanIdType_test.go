package core

import (
	"testing"
	"unsafe"
)

func TestVlanIdType(t *testing.T) {
	t.Run("test type size", func(t *testing.T) {
		o := VlanIdType(0)
		if unsafe.Sizeof(o) != 2 {
			t.Fatalf("size mismatch")
		}
	})
	t.Run("vlanId.ToInt() test", func(t *testing.T) {
		var vlanId VlanIdType = 100

		result := vlanId.ToInt()

		if result != 100 {
			t.Errorf("ToInt() = %d; want 100", result)
		}
	})
	t.Run("vlanId.FromInt() test", func(t *testing.T) {
		var vlanId VlanIdType

		// Test with a valid VLAN ID
		err := vlanId.FromInt(100)
		if err != nil {
			t.Errorf("FromInt(100) returned an unexpected error: %v", err)
		}
		if vlanId.ToInt() != 100 {
			t.Errorf("FromInt(100) = %d; want 100", vlanId.ToInt())
		}

		// Test with a minimum valid VLAN ID
		err = vlanId.FromInt(1)
		if err != nil {
			t.Errorf("FromInt(1) returned an unexpected error: %v", err)
		}
		if vlanId.ToInt() != 1 {
			t.Errorf("FromInt(1) = %d; want 1", vlanId.ToInt())
		}

		// Test with a maximum valid VLAN ID
		err = vlanId.FromInt(4094)
		if err != nil {
			t.Errorf("FromInt(4094) returned an unexpected error: %v", err)
		}
		if vlanId.ToInt() != 4094 {
			t.Errorf("FromInt(4094) = %d; want 4094", vlanId.ToInt())
		}

		// Test with an invalid VLAN ID (too low)
		err = vlanId.FromInt(0)
		if err == nil {
			t.Errorf("FromInt(0) expected error but got nil")
		}

		// Test with an invalid VLAN ID (too high)
		err = vlanId.FromInt(4095)
		if err == nil {
			t.Errorf("FromInt(4095) expected error but got nil")
		}
	})
}
