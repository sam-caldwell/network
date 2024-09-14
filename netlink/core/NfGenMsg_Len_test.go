package core

import (
	"testing"
)

func TestNfGenMsg_Len(t *testing.T) {
	t.Run("Test NfGenMsg Len method", func(t *testing.T) {
		nfGenMsg := NfGenMsg{}
		expectedLen := NfGenMsgSize

		if nfGenMsg.Len() != expectedLen {
			t.Errorf("Expected Len: %d, got: %d", expectedLen, nfGenMsg.Len())
		}
	})
}
