package core

import (
	"testing"
)

func TestNullExecuteIterFunc(t *testing.T) {
	if !nullExecuteIterFunc(nil) {
		t.Fatal("expected nullExecuteIterFunc to always return true")
	}
}
