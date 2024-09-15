package core

import (
	"testing"
)

func TestGenericNetlinkMessage_constants(t *testing.T) {
	t.Run("test numeric values", func(t *testing.T) {
		testData := []struct {
			actual int
			expect int
		}{
			{actual: GenlAdminPerm, expect: 0x1},    // unix.GENL_ADMIN_PERM
			{actual: GenlCmdCapDo, expect: 0x2},     // unix.GENL_CMD_CAP_DO
			{actual: GenlCmdCapDump, expect: 0x4},   // unix.GENL_CMD_CAP_DUMP
			{actual: GenlCmdCapHaspol, expect: 0x8}, // unix.GENL_CMD_CAP_HASPOL
			{actual: GenlHdrlen, expect: 0x4},       // unix.GENL_HDRLEN
			{actual: GenlIdCtrl, expect: 0x10},      // unix.GENL_ID_CTRL

			{actual: GenlGtpVersion, expect: 0},

			{actual: GenlCtrlVersion, expect: 2},
			{actual: GenlCtrlCmdGetFamily, expect: 3},
		}
		for _, data := range testData {
			if data.actual != data.expect {
				t.Errorf("got %d, want %d", data.actual, data.expect)
			}
		}
	})
	t.Run("test string values", func(t *testing.T) {
		testData := []struct {
			actual string
			expect string
		}{
			{actual: GenlGtpName, expect: "gtp"},
			{actual: GenlCtrlName, expect: "nlctrl"},
		}
		for _, data := range testData {
			if data.actual != data.expect {
				t.Errorf("got %s, want %s", data.actual, data.expect)
			}
		}
	})
}
