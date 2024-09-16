package core

import (
	"errors"
	"testing"
)

func TestIpSetError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      IpSetError
		expected string
	}{
		{"IpsetErrPrivate", IpsetErrPrivate, "private"},
		{"IpsetErrProtocol", IpsetErrProtocol, "invalid protocol"},
		{"IpsetErrFindType", IpsetErrFindType, "invalid type"},
		{"IpsetErrMaxSets", IpsetErrMaxSets, "max sets reached"},
		{"IpsetErrBusy", IpsetErrBusy, "busy"},
		{"IpsetErrExistSetname2", IpsetErrExistSetname2, "exist_setname2"},
		{"IpsetErrTypeMismatch", IpsetErrTypeMismatch, "type mismatch"},
		{"IpsetErrExist", IpsetErrExist, "exist"},
		{"IpsetErrInvalidCidr", IpsetErrInvalidCidr, "invalid cidr"},
		{"IpsetErrInvalidNetmask", IpsetErrInvalidNetmask, "invalid netmask"},
		{"IpsetErrInvalidFamily", IpsetErrInvalidFamily, "invalid family"},
		{"IpsetErrTimeout", IpsetErrTimeout, "timeout"},
		{"IpsetErrReferenced", IpsetErrReferenced, "referenced"},
		{"IpsetErrIpaddrIpv4", IpsetErrIpaddrIpv4, "invalid ipv4 address"},
		{"IpsetErrIpaddrIpv6", IpsetErrIpaddrIpv6, "invalid ipv6 address"},
		{"IpsetErrCounter", IpsetErrCounter, "invalid counter"},
		{"IpsetErrComment", IpsetErrComment, "invalid comment"},
		{"IpsetErrInvalidMarkmask", IpsetErrInvalidMarkmask, "invalid markmask"},
		{"IpsetErrSkbinfo", IpsetErrSkbinfo, "skbinfo"},
		{"UnknownError", IpSetError(9999), "errno 9999"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.err.Error()
			expectedErr := errors.New(test.expected)

			if err.Error() != expectedErr.Error() {
				t.Errorf("Expected error %s but got %s", expectedErr.Error(), err.Error())
			}
		})
	}
}
