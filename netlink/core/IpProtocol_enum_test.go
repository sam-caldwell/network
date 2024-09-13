package core

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

// TestIpProtocol tests the values of the IpProtocol constants.
func TestIpProtocol(t *testing.T) {
	t.Run("test type size", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(IpProtocol(0))); sz != expectedSizeInBytes {
			t.Fatal("size mismatch")
		}
	})
	t.Run("Test enumerated values", func(t *testing.T) {
		testData := []struct {
			actual IpProtocol
			expect IpProtocol
		}{
			{actual: IpProtoHOPOPT, expect: unix.IPPROTO_HOPOPTS},
			{actual: IpProtoHOPOPTS, expect: unix.IPPROTO_HOPOPTS},
			{actual: IpProtoIcmp, expect: unix.IPPROTO_ICMP},
			{actual: IpProtoIGMP, expect: 0x02},
			{actual: IpProtoGGP, expect: 0x03},
			{actual: IpProtoIPv4Encapsulation, expect: 0x04},
			{actual: IpProtoST, expect: 0x05},
			{actual: IpProtoTCP, expect: unix.IPPROTO_TCP},
			{actual: IpProtoCBT, expect: 0x07},
			{actual: IpProtoEGP, expect: 0x08},
			{actual: IpProtoIGP, expect: 0x09},
			{actual: IpProtoBBNRCCMON, expect: 0x0A},
			{actual: IpProtoNVP_II, expect: 0x0B},
			{actual: IpProtoPUP, expect: 0x0C},
			{actual: IpProtoARGUS, expect: 0x0D},
			{actual: IpProtoEMCON, expect: 0x0E},
			{actual: IpProtoXNET, expect: 0x0F},
			{actual: IpProtoCHAOS, expect: 0x10},
			{actual: IpProtoUDP, expect: unix.IPPROTO_UDP},
			{actual: IpProtoMUX, expect: 0x12},
			{actual: IpProtoDCNMEAS, expect: 0x13},
			{actual: IpProtoHMP, expect: 0x14},
			{actual: IpProtoPRM, expect: 0x15},
			{actual: IpProtoXNSIDP, expect: 0x16},
			{actual: IpProtoTRUNK1, expect: 0x17},
			{actual: IpProtoTRUNK2, expect: 0x18},
			{actual: IpProtoLEAF1, expect: 0x19},
			{actual: IpProtoLEAF2, expect: 0x1A},
			{actual: IpProtoRDP, expect: 0x1B},
			{actual: IpProtoIRTP, expect: 0x1C},
			{actual: IpProtoISO_TP4, expect: 0x1D},
			{actual: IpProtoNETBLT, expect: 0x1E},
			{actual: IpProtoMFENSP, expect: 0x1F},
			{actual: IpProtoMERIT_INP, expect: 0x20},
			{actual: IpProtoDCCP, expect: unix.IPPROTO_DCCP},
			{actual: IpProtoThirdPC, expect: 0x22},
			{actual: IpProtoIDPR, expect: 0x23},
			{actual: IpProtoXTP, expect: 0x24},
			{actual: IpProtoDDP, expect: 0x25},
			{actual: IpProtoIDPR_CMTP, expect: 0x26},
			{actual: IpProtoTPXX, expect: 0x27},
			{actual: IpProtoIL, expect: 0x28},
			{actual: IpProtoIPv6Encapsulation, expect: 0x29},
			{actual: IpProtoSDRP, expect: 0x2A},
			{actual: IpProtoIPv6Route, expect: 0x2B},
			{actual: IpProtoIPv6Frag, expect: 0x2C},
			{actual: IpProtoIDRP, expect: 0x2D},
			{actual: IpProtoRSVP, expect: 0x2E},
			{actual: IpProtoGRE, expect: 0x2F},
			{actual: IpProtoMHRP, expect: 0x30},
			{actual: IpProtoBNA, expect: 0x31},
			{actual: IpProtoESP, expect: 0x32},
			{actual: IpProtoAH, expect: 0x33},
			{actual: IpProtoINLSP, expect: 0x34},
			{actual: IpProtoSWIPE, expect: 0x35},
			{actual: IpProtoNARP, expect: 0x36},
			{actual: IpProtoMOBILE, expect: 0x37},
			{actual: IpProtoTLSP, expect: 0x38},
			{actual: IpProtoSKIP, expect: 0x39},
			{actual: IpProtoIPv6ICMP, expect: unix.IPPROTO_ICMPV6},
			{actual: IpProtoIPv6NoNxt, expect: 0x3B},
			{actual: IpProtoIPv6Opts, expect: 0x3C},
			{actual: IpProtoAnyHostInternal, expect: 61},
			{actual: IpProtoCFTP, expect: 0x3E},
			{actual: IpProtoAnyLocalNetwork, expect: 63},
			{actual: IpProtoSatExpak, expect: 0x40},
			{actual: IpProtoKRYPTOLAN, expect: 0x41},
			{actual: IpProtoRVD, expect: 0x42},
			{actual: IpProtoIPPC, expect: 0x43},
			{actual: IpProtoAnyDistributedFs, expect: 68},
			{actual: IpProtoSATMON, expect: 0x45},
			{actual: IpProtoVISA, expect: 0x46},
			{actual: IpProtoIpcv, expect: 0x47},
			{actual: IpProtoCPNX, expect: 0x48},
			{actual: IpProtoCPHB, expect: 0x49},
			{actual: IpProtoWSN, expect: 0x4A},
			{actual: IpProtoPVP, expect: 0x4B},
			{actual: IpProtoBRSATMON, expect: 0x4C},
			{actual: IpProtoSUNND, expect: 0x4D},
			{actual: IpProtoWBMON, expect: 0x4E},
			{actual: IpProtoWBEXPAK, expect: 0x4F},
			{actual: IpProtoISOIP, expect: 0x50},
			{actual: IpProtoVMTP, expect: 0x51},
			{actual: IpProtoSECUREVMTP, expect: 0x52},
			{actual: IpProtoVINES, expect: 0x53},
			{actual: IpProtoTTP, expect: 0x54},
			{actual: IpProtoIptm, expect: 0x54},
			{actual: IpProtoNSFNETIGP, expect: 0x55},
			{actual: IpProtoDGP, expect: 0x56},
			{actual: IpProtoTCF, expect: 0x57},
			{actual: IpProtoEIGRP, expect: 0x58},
			{actual: IpProtoOSPF, expect: 0x59},
			{actual: IpProtoSPRITRPC, expect: 0x5A},
			{actual: IpProtoLARP, expect: 0x5B},
			{actual: IpProtoMTP, expect: 0x5C},
			{actual: IpProtoAX25, expect: 0x5D},
			{actual: IpProtoIPIP, expect: 0x5E},
			{actual: IpProtoMICP, expect: 0x5F},
			{actual: IpProtoSCCSP, expect: 0x60},
			{actual: IpProtoETHERIP, expect: 0x61},
			{actual: IpProtoENCAP, expect: 0x62},
			{actual: IpProtoAnyPrivateEncryption, expect: 99},
			{actual: IpProtoGMTP, expect: 0x64},
			{actual: IpProtoIFMP, expect: 0x65},
			{actual: IpProtoPNNI, expect: 0x66},
			{actual: IpProtoPIM, expect: 0x67},
			{actual: IpProtoARIS, expect: 0x68},
			{actual: IpProtoSCPS, expect: 0x69},
			{actual: IpProtoQNX, expect: 0x6A},
			{actual: IpProtoAN, expect: 0x6B},
			{actual: IpProtoIPComp, expect: 0x6C},
			{actual: IpProtoSNP, expect: 0x6D},
			{actual: IpProtoCompaqPeer, expect: 0x6E},
			{actual: IpProtoIpxInIp, expect: 0x6F},
			{actual: IpProtoVRRP, expect: 0x70},
			{actual: IpProtoPGM, expect: 0x71},
			{actual: Any0HopProtocol, expect: 0x72},
			{actual: IpProtoL2TP, expect: 0x73},
			{actual: IpProtoDDX, expect: 0x74},
			{actual: IpProtoIATP, expect: 0x75},
			{actual: IpProtoSTP, expect: 0x76},
			{actual: IpProtoSRP, expect: 0x77},
			{actual: IpProtoUTI, expect: 0x78},
			{actual: IpProtoSMP, expect: 0x79},
			{actual: IpProtoSM, expect: 0x7A},
			{actual: IpProtoPTP, expect: 0x7B},
			{actual: IpProtoIsisOverIpv4, expect: 0x7C},
			{actual: IpProtoFIRE, expect: 0x7D},
			{actual: IpProtoCRTP, expect: 0x7E},
			{actual: IpProtoCRUDP, expect: 0x7F},
			{actual: IpProtoSSCOPMCE, expect: 0x80},
			{actual: IpProtoIPLT, expect: 0x81},
			{actual: IpProtoSPS, expect: 0x82},
			{actual: IpProtoPIPE, expect: 0x83},
			{actual: IpProtoSctp, expect: unix.IPPROTO_SCTP},
			{actual: IpProtoFC, expect: 0x85},
			{actual: IpProtoRsvpE2EIgnore, expect: 134},
			{actual: IpProtoMobilityHeader, expect: 135},
			{actual: IpProtoUdpLite, expect: 136},
			{actual: IpProtoMplsInIp, expect: 137},
			{actual: IpProtoManet, expect: 0x8A},
			{actual: IpProtoHIP, expect: 0x8B},
			{actual: IpProtoShim6, expect: 0x8C},
			{actual: IpProtoWesp, expect: 0x8D},
			{actual: IpProtoRohc, expect: 0x8E},
			{actual: IpProtoEthernet, expect: 0x8F},
			{actual: IpProtoAggFrag, expect: 0x90},
			{actual: IpProtoNsh, expect: 0x91},
			{actual: IpProtoMobility, expect: 0x92},
			{actual: IpProtoRohcTransport, expect: 0x93},
			{actual: IpProtoRohcInternal, expect: 0x94},
			{actual: IpProtoWespInternal, expect: 0x95},
			{actual: IpProtoRohcCompressed, expect: 0x96},
			{actual: IpProtoReserved, expect: 0xFF},
		}

		for _, test := range testData {
			if test.actual != test.expect {
				t.Errorf("Expected %v but got %v", test.expect, test.actual)
			}
		}
	})
	t.Run("test Serialize() method", func(t *testing.T) {

	})
}
