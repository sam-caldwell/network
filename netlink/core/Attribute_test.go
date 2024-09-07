package core

import (
	"math"
	"testing"
)

func TestAttributeStruct(t *testing.T) {
	t.Run("zero test", func(t *testing.T) {
		_ = Attribute{
			Type:  NlaFlags(0),
			Value: []byte(""),
		}
	})
	t.Run("max test", func(t *testing.T) {
		o := Attribute{
			Type:  NlaFlags(math.MaxUint16),
			Value: []byte(""),
		}
		if o.Type != math.MaxUint16 {
			t.Fatalf("max value expected")
		}
	})
}
