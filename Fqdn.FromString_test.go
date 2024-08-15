package network

import (
	"fmt"
	"testing"
)

func TestFqdnFromString(t *testing.T) {
	var f Fqdn

	testData := map[string]error{
		"verify.this.tld":          nil,
		"":                         fmt.Errorf("FQDN must be valid"),
		"127.0.0.1":                nil,
		"192.168.1.1":              nil,
		"fe80::1fa8:4cd:8210:5e87": nil,
		"ver%ify.this.tld":         fmt.Errorf("FQDN must be valid"),
	}

	for k, v := range testData {
		if err := f.FromString(&k); err != nil && err.Error() != v.Error() {
			t.Fatalf("on %s expected: '%v', got: '%v'", k, v, err)
		}
	}
}
