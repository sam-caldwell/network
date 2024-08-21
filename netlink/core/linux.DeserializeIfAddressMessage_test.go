package core

import (
	"golang.org/x/sys/unix"
	"reflect"
	"testing"
)

func TestIfAddrmsgDeserializeSerialize(t *testing.T) {
	var expectedOutput IfAddressMessage
	expectedOutput.Family = unix.AF_INET
	expectedOutput.Prefixlen = 24
	expectedOutput.Flags = 1
	expectedOutput.Scope = 2

	var testInput = []byte{
		uint8(unix.AF_INET),
		uint8(24),
		uint8(1),
		uint8(2),
	}

	if len(testInput) != SizeOfIfAddressMessage {
		t.Fatal("testInput not right size")
	}

	if actualOutput, err := DeserializeIfAddressMessage(testInput); err != nil {
		t.Fatal(err)
	} else {
		if !reflect.DeepEqual(actualOutput, expectedOutput) {
			t.Fatalf("Deserialize failed: expected %v, actual %v", expectedOutput, actualOutput)
		}
	}
}
