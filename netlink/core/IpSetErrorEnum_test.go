package core

import (
	"errors"
	"strconv"
	"testing"
	"unsafe"
)

func TestIpSetErrorEnum(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = unsafe.Sizeof(uintptr(0)) // Size of uintptr will be architecture dependent
		var err IpSetErrorEnum

		if size := unsafe.Sizeof(err); size != expectedSizeInBytes {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    IpSetErrorEnum
			expected IpSetErrorEnum
		}{
			{"IpsetErrPrivate", IpsetErrPrivate, 4096},
			{"IpsetErrProtocol", IpsetErrProtocol, 4097},
			{"IpsetErrFindType", IpsetErrFindType, 4098},
			{"IpsetErrMaxSets", IpsetErrMaxSets, 4099},
			{"IpsetErrBusy", IpsetErrBusy, 4100},
			{"IpsetErrExistSetname2", IpsetErrExistSetname2, 4101},
			{"IpsetErrTypeMismatch", IpsetErrTypeMismatch, 4102},
			{"IpsetErrExist", IpsetErrExist, 4103},
			{"IpsetErrInvalidCidr", IpsetErrInvalidCidr, 4104},
			{"IpsetErrInvalidNetmask", IpsetErrInvalidNetmask, 4105},
			{"IpsetErrInvalidFamily", IpsetErrInvalidFamily, 4106},
			{"IpsetErrTimeout", IpsetErrTimeout, 4107},
			{"IpsetErrReferenced", IpsetErrReferenced, 4108},
			{"IpsetErrIpaddrIpv4", IpsetErrIpaddrIpv4, 4109},
			{"IpsetErrIpaddrIpv6", IpsetErrIpaddrIpv6, 4110},
			{"IpsetErrCounter", IpsetErrCounter, 4111},
			{"IpsetErrComment", IpsetErrComment, 4112},
			{"IpsetErrInvalidMarkmask", IpsetErrInvalidMarkmask, 4113},
			{"IpsetErrSkbinfo", IpsetErrSkbinfo, 4114},
			{"IpsetErrTypeSpecific", IpsetErrTypeSpecific, 4352},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if test.value != test.expected {
					t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
				}
			})
		}
	})
	t.Run("test ErrorString() method", func(t *testing.T) {
		tests := []struct {
			name     string
			err      IpSetErrorEnum
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
			{"UnknownError", IpSetErrorEnum(9999), "errno " + strconv.Itoa(9999)},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := test.err.ErrorString()
				if result != test.expected {
					t.Errorf("Expected %s but got %s for error %d", test.expected, result, test.err)
				}
			})
		}
	})
	t.Run("test Error() method", func(t *testing.T) {
		tests := []struct {
			name     string
			err      IpSetErrorEnum
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
			{"UnknownError", IpSetErrorEnum(9999), "errno 9999"},
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
	})
}
