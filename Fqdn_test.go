package network

import "testing"

func TestFqdnType(t *testing.T) {
	t.Run("test string cast", func(t *testing.T) {
		_ = Fqdn("test") //just make sure a string can cast evenly.
	})
	t.Run("test comparable", func(t *testing.T) {
		a := Fqdn("test")
		b := Fqdn("test")
		if a != b {
			t.Fatal("expect comparison operators to work")
		}
	})
}
