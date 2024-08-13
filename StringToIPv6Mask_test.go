package network

import (
	"net"
	"reflect"
	"testing"
)

func TestStringToIPv6Mask(t *testing.T) {
	tests := []struct {
		cidr         string
		expectedMask net.IPMask
		expectingErr bool
	}{
		{
			cidr:         "[::]/0",
			expectedMask: net.CIDRMask(0, 128),
			expectingErr: false,
		},
		{
			cidr:         "2001:db8::/32",
			expectedMask: net.CIDRMask(32, 128),
			expectingErr: false,
		},
		{
			cidr:         "invalid_cidr",
			expectedMask: nil,
			expectingErr: true,
		},
	}

	for _, test := range tests {
		mask, err := StringToIPv6Mask(test.cidr)

		if test.expectingErr {
			if err == nil {
				t.Fatalf("expected an error for CIDR '%s', but got none", test.cidr)
			}
		} else {
			if err != nil {
				t.Fatalf("unexpected error for CIDR '%s'\n"+
					" mask: %v\n"+
					"error: %v", test.cidr, mask, err)
			}
			if !reflect.DeepEqual(mask, test.expectedMask) {
				t.Fatalf("expected mask '%v' for CIDR '%s', but got '%v'", test.expectedMask, test.cidr, mask)
			}
		}
	}
}
