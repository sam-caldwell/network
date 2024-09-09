package core

import (
	"testing"
	"unsafe"
)

func TestTcU32Enum(t *testing.T) {
	const expectedSizeInBytes = 1
	if sz := unsafe.Sizeof(TcU32Enum(0)); sz != expectedSizeInBytes {
		t.Fatalf("expected TcU32Enum to be %d byte", expectedSizeInBytes)
	}
	if v := TcU32Enum(1); v != TcU32Terminal {
		t.Fatalf("expected TcU32Enum to be TcU32Terminal")
	}
	if v := TcU32Enum(2); v != TcU32Offset {
		t.Fatalf("expected TcU32Enum to be TcU32Offset")
	}
	if v := TcU32Enum(4); v != TcU32Varoffset {
		t.Fatalf("expected TcU32Enum to be TcU32Varoffset")
	}
	if v := TcU32Enum(8); v != TcU32Eat {
		t.Fatalf("expected TcU32Enum to be TcU32Eat")
	}
}
