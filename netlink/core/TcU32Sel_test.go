package core

import (
	"reflect"
	"testing"
	"unsafe"
)

// TestTcU32SelStruct verifies the presence, types, and size of the TcU32Sel struct fields.
func TestTcU32SelStruct(t *testing.T) {
	t.Run("verify fields and types", func(t *testing.T) {
		// Create an instance of TcU32Sel to test its fields.
		var sel TcU32Sel

		// Use reflection to check the struct fields and their types.
		selType := reflect.TypeOf(sel)

		// Define the expected fields and their types.
		expectedFields := []struct {
			Name string
			Type reflect.Type
		}{
			{"Flags", reflect.TypeOf(uint8(0))},
			{"Offshift", reflect.TypeOf(uint8(0))},
			{"Nkeys", reflect.TypeOf(uint8(0))},
			{"Pad", reflect.TypeOf(uint8(0))},
			{"Offmask", reflect.TypeOf(uint16(0))},
			{"Off", reflect.TypeOf(uint16(0))},
			{"Offoff", reflect.TypeOf(int16(0))},
			{"Hoff", reflect.TypeOf(int16(0))},
			{"Hmask", reflect.TypeOf(uint32(0))},
			{"Keys", reflect.TypeOf([]TcU32Key{})},
		}

		// Check if the number of fields matches.
		if selType.NumField() != len(expectedFields) {
			t.Fatalf("Expected %d fields, but got %d", len(expectedFields), selType.NumField())
		}

		// Verify each field's name and type.
		for i, expectedField := range expectedFields {
			field := selType.Field(i)

			if field.Name != expectedField.Name {
				t.Errorf("Expected field name %q, but got %q", expectedField.Name, field.Name)
			}

			if field.Type != expectedField.Type {
				t.Errorf("Expected field %q to be of type %v, but got %v", field.Name, expectedField.Type, field.Type)
			}
		}

		// Verify the size of the TcU32Sel struct.
		expectedSize := uintptr(unsafe.Sizeof(sel))
		actualSize := uintptr(unsafe.Sizeof(sel))

		// You can modify the expected size according to your system architecture or assumptions.
		if actualSize != expectedSize {
			t.Errorf("Expected size of TcU32Sel to be %d, but got %d", expectedSize, actualSize)
		}
	})
}
