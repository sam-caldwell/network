package core

import (
	"testing"
	"unsafe"
)

func TestConnectorMessageOperation(t *testing.T) {
	t.Run("ConnectorMessageOperation struct", func(t *testing.T) {
		t.Run("test ConnectorMessageOperation fields", func(t *testing.T) {
			_ = ConnectorMessageOperation{
				ConnectorMessage: ConnectorMessage{
					ID: CbID{
						Idx: uint32(0),
						Val: uint32(0),
					},
					Seq:    uint32(0),
					Ack:    uint32(0),
					Length: uint16(0),
					Flags:  uint16(0),
				},
				Operation: uint32(0),
			}
		})
		t.Run("size check", func(t *testing.T) {
			var o ConnectorMessageOperation
			t.Run("Size of ConnectorMessageOperation", func(t *testing.T) {
				if unsafe.Sizeof(o.ConnectorMessage) != unsafe.Sizeof(ConnectorMessage{}) {
					t.Fatalf("ConnectorMessage size verification")
				}
				if unsafe.Sizeof(o.Operation) != unsafe.Sizeof(uint32(0)) {
					t.Fatalf("Operation size verification")
				}
			})
		})
	})
}
