package core

import (
	"testing"
	"unsafe"
)

func TestIpSetErrorEnum(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = unsafe.Sizeof(IpSetError(0)) // Size of uintptr will be architecture dependent
		var err IpSetError

		if size := unsafe.Sizeof(err); size != expectedSizeInBytes {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    IpSetError
			expected IpSetError
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
}
