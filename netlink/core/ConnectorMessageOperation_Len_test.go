package core

import (
	"testing"
	"unsafe"
)

func TestConnectorMessageOperation_Len(t *testing.T) {
	var o ConnectorMessageOperation
	if o.Len() != int(unsafe.Sizeof(ConnectorMessageOperation{})) {
		t.Fatalf("ConnectorMessageOperation.Len() method size mismatch")
	}
}
