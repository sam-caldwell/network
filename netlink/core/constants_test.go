package core

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestConstantsStructureSize(t *testing.T) {
	type TestDataStruct struct {
		actual   int
		expected int
	}
	testData := []TestDataStruct{
		{expected: int(unsafe.Sizeof(NfGenMsg{})), actual: NfGenMsgSize},
		{expected: int(unsafe.Sizeof(NetlinkMessageHeader{})), actual: NetlinkMessageHeaderSize},
		{expected: int(unsafe.Sizeof(RtNexthop{})), actual: SizeOfRtNextHop},
		{expected: int(unsafe.Sizeof(TcNetemRate{})), actual: SizeOfTcNetemRate},
		{expected: int(unsafe.Sizeof(TcActionMsg{})), actual: SizeOfTcActionMsg},
		{expected: int(unsafe.Sizeof(TcConnmark{})), actual: SizeOfTcConnmark},
		{expected: int(unsafe.Sizeof(TcCsum{})), actual: SizeOfTcCsum},
		{expected: int(unsafe.Sizeof(TcGen{})), actual: SizeOfTcGen},
		{expected: int(unsafe.Sizeof(TcHtbCopt{})), actual: SizeOfTcHtbCopt},
		{expected: int(unsafe.Sizeof(TcHtbGlob{})), actual: SizeOfTcHtbGlob},
		{expected: int(unsafe.Sizeof(IfInfoMsg{})), actual: IfInfoMsgSize},
		{expected: int(unsafe.Sizeof(TcMsg{})), actual: SizeOfTcMsg},
		{expected: int(unsafe.Sizeof(TcMirred{})), actual: SizeOfTcMirred},
		{expected: int(unsafe.Sizeof(TcNetemCorr{})), actual: SizeOfTcNetemCorr},
		{expected: int(unsafe.Sizeof(TcNetemCorrupt{})), actual: SizeOfTcNetemCorrupt},
		{expected: int(unsafe.Sizeof(TcNetemQopt{})), actual: SizeOfTcNetemQopt},
		{expected: int(unsafe.Sizeof(TcNetemReorder{})), actual: SizeOfTcNetemReorder},
		{expected: int(unsafe.Sizeof(TcPeditKey{})), actual: SizeOfTcPeditKey},
		{expected: int(unsafe.Sizeof(TcPeditSel{})), actual: SizeOfTcPeditSel},
		{expected: int(unsafe.Sizeof(TcPolice{})), actual: SizeOfTcPolice},
		{expected: int(unsafe.Sizeof(TcPriorityMap{})), actual: SizeOfTcPriorityMap},
		{expected: int(unsafe.Sizeof(TcRateSpec{})), actual: SizeOfTcRateSpec},
		{expected: int(unsafe.Sizeof(TcSfqQopt{})), actual: SizeOfTcSfqQopt},
		{expected: int(unsafe.Sizeof(TcSfqQoptV1{})), actual: SizeOfTcSfqQoptV1},
		{expected: int(unsafe.Sizeof(TcSfqRedStats{})), actual: SizeOfTcSfqRedStats},
		{expected: int(unsafe.Sizeof(TcSkbEdit{})), actual: SizeOfTcSkbEdit},
		{expected: int(unsafe.Sizeof(TcTbfQopt{})), actual: SizeOfTcTbfQopt},
		{expected: 16, actual: SizeOfTcU32Sel},
		{expected: 16, actual: SizeOfTcU32Key},
		{expected: int(unsafe.Sizeof(TcTunnelKey{})), actual: SizeOfTcTunnelKey},
		{expected: int(unsafe.Sizeof(Tcf{})), actual: SizeOfTcf},
		{expected: int(unsafe.Sizeof(Uint32Bitfield{})), actual: SizeOfUint32Bitfield},
		{expected: int(unsafe.Sizeof(unix.RtAttr{})), actual: SizeOfUnixRtAttr},
		{expected: int(unsafe.Sizeof(unix.RtMsg{})), actual: SizeOfUnixRtMsg},
		{expected: int(unsafe.Sizeof(VfGUID{})), actual: SizeOfVfGUID},
		{expected: int(unsafe.Sizeof(VfLinkState{})), actual: SizeOfVfLinkState},
		{expected: int(unsafe.Sizeof(VfMac{})), actual: SizeOfVfMac},
		{expected: int(unsafe.Sizeof(VfRate{})), actual: SizeOfVfRate},
		{expected: int(unsafe.Sizeof(VfRssQueryEn{})), actual: SizeOfVfRssQueryEn},
		{expected: int(unsafe.Sizeof(VfSpoofchk{})), actual: SizeOfVfSpoofchk},
		{expected: int(unsafe.Sizeof(VfTrust{})), actual: SizeOfVfTrust},
		{expected: int(unsafe.Sizeof(VfTxRate{})), actual: SizeOfVfTxRate},
		{expected: int(unsafe.Sizeof(VfVlan{})), actual: SizeOfVfVlan},
		{expected: int(unsafe.Sizeof(IfLaVfVlanInfoStruct{})), actual: SizeOfVfVlanInfo},
		{expected: int(4 * 32), actual: SizeOfXfrmAddress},
		{expected: int(unsafe.Sizeof(XfrmAlgo{})), actual: SizeOfXfrmAlgo},
		{expected: int(unsafe.Sizeof(XfrmAlgoAEAD{})), actual: SizeOfXfrmAlgoAEAD},
		{expected: int(unsafe.Sizeof(XfrmAlgoAuth{})), actual: SizeOfXfrmAlgoAuth},
		{expected: int(unsafe.Sizeof(XfrmEncapTmpl{})), actual: SizeOfXfrmEncapTmpl},
		{expected: int(unsafe.Sizeof(XfrmId{})), actual: SizeOfXfrmId},
		{expected: int(unsafe.Sizeof(XfrmLifetimeCfg{})), actual: SizeOfXfrmLifetimeCfg},
		{expected: int(unsafe.Sizeof(XfrmLifetimeCur{})), actual: SizeOfXfrmLifetimeCur},
		{expected: int(unsafe.Sizeof(XfrmMark{})), actual: SizeOfXfrmMark},
		{expected: int(unsafe.Sizeof(XfrmUserPolicyId{})), actual: SizeOfXfrmUserPolicyId},
		{expected: int(unsafe.Sizeof(XfrmUserPolicyInfo{})), actual: SizeOfXfrmUserPolicyInfo},
		{expected: int(unsafe.Sizeof(XfrmReplayState{})), actual: SizeOfXfrmReplayState},
		{expected: int(unsafe.Sizeof(XfrmReplayStateEsn{})), actual: SizeOfXfrmReplayStateEsn},
		{expected: int(unsafe.Sizeof(XfrmSelector{})), actual: SizeOfXfrmSelector},
		{expected: int(unsafe.Sizeof(XfrmStats{})), actual: SizeOfXfrmStats},
		{expected: int(unsafe.Sizeof(XfrmUserExpire{})), actual: SizeOfXfrmUserExpire},
		{expected: int(unsafe.Sizeof(XfrmUserSaFlush{})), actual: SizeOfXfrmUserSaFlush},
		{expected: int(unsafe.Sizeof(XfrmUserSaId{})), actual: SizeOfXfrmUserSaId},
		{expected: int(unsafe.Sizeof(XfrmUserSaInfo{})), actual: SizeOfXfrmUserSaInfo},
		{expected: int(unsafe.Sizeof(XfrmUserSpiInfo{})), actual: SizeOfXfrmUserSpiInfo},
		{expected: int(unsafe.Sizeof(XfrmUserTmpl{})), actual: SizeOfXfrmUserTmpl},
	}
	for _, row := range testData {
		if row.actual != row.expected {
			t.Fatalf("value mismatch (actual %d, expected %d)", row.actual, row.expected)
		}
	}
}
