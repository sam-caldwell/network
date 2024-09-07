package core

import (
	"testing"
	"unsafe"
)

func TestBridgePortEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		o := BridgePortEnum(0)
		if unsafe.Sizeof(o) != expectedSizeInBytes {
			t.Fatalf("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		if value := BridgePortEnum(0); value != IflaBrPortUnspec {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(1); value != IflaBrPortState {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(2); value != IflaBrPortPriority {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(3); value != IflaBrPortCost {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(4); value != IflaBrPortMode {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(5); value != IflaBrPortGuard {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(6); value != IflaBrPortProtect {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(7); value != IflaBrPortFastLeave {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(8); value != IflaBrPortLearning {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(9); value != IflaBrPortUnicastFlood {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(10); value != IflaBrPortProxyArp {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(11); value != IflaBrPortLearningSync {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(12); value != IflaBrPortProxyArpWifi {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(13); value != IflaBrPortRootId {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(14); value != IflaBrPortBridgeId {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(15); value != IflaBrPortDesignatedPort {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(16); value != IflaBrPortDesignatedCost {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(17); value != IflaBrPortId {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(18); value != IflaBrPortNo {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(19); value != IflaBrPortTopologyChangeAck {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(20); value != IflaBrPortConfigPending {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(21); value != IflaBrPortMessageAgeTimer {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(22); value != IflaBrPortForwardDelayTimer {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(23); value != IflaBrPortHoldTimer {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(24); value != IflaBrPortFlush {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(25); value != IflaBrPortMulticastRouter {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(26); value != IflaBrPortPad {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(27); value != IflaBrPortMcastFlood {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(28); value != IflaBrPortMcastToUcast {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(29); value != IflaBrPortVlanTunnel {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(30); value != IflaBrPortBcastFlood {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(31); value != IflaBrPortGroupFwdMask {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(32); value != IflaBrPortNeighSuppress {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(33); value != IflaBrPortIsolated {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(34); value != IflaBrPortBackupPort {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(35); value != IflaBrPortMrpRingOpen {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(36); value != IflaBrPortMrpInOpen {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(37); value != IflaBrPortMcastEhtHostsLimit {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(38); value != IflaBrPortMcastEhtHostsCnt {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(39); value != IflaBrPortLocked {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(40); value != IflaBrPortMab {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(41); value != IflaBrPortMcastNGroups {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(42); value != IflaBrPortMcastMaxGroups {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(43); value != IflaBrPortNeighVlanSuppress {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(44); value != IflaBrPortBackupNhid {
			t.Fatalf("value check failed %d", value)
		}
		if value := BridgePortEnum(44); value != IflaBrPortMax {
			t.Fatalf("value check failed %d", value)
		}
	})
}
