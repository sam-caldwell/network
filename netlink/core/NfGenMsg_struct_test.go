package core

import (
	"testing"
	"unsafe"
)

func TestNfGenMsg(t *testing.T) {
	t.Run("Test NfGenMsg struct size", func(t *testing.T) {
		expectedSize := 4 // Size in bytes: 1 byte for NfgenFamily, 1 byte for Version, 2 bytes for ResId
		actualSize := int(unsafe.Sizeof(NfGenMsg{}))
		if actualSize != expectedSize {
			t.Errorf("Expected struct size: %d, got: %d", expectedSize, actualSize)
		}
	})

	t.Run("Test NfGenMsg field values", func(t *testing.T) {
		nfGenMsg := NfGenMsg{
			NfgenFamily: 2,   // Assuming AF_INET (IPv4)
			Version:     1,   // Netfilter netlink message version
			ResId:       256, // Sample resource identifier
		}

		if nfGenMsg.NfgenFamily != 2 {
			t.Errorf("Expected NfgenFamily: %d, got: %d", 2, nfGenMsg.NfgenFamily)
		}

		if nfGenMsg.Version != 1 {
			t.Errorf("Expected Version: %d, got: %d", 1, nfGenMsg.Version)
		}

		if nfGenMsg.ResId != 256 {
			t.Errorf("Expected ResId: %d, got: %d", 256, nfGenMsg.ResId)
		}
	})
}
