package core

import (
	"golang.org/x/sys/unix"
)

// IpProtocol - represents the `ip_proto` attribute in the Flower classifier. Flower is a flexible, flow-based
// packet classification mechanism used in traffic control.
//
// The `ip_proto` attribute in Flower is used to match the protocol number in the IP header of a packet.
//
// See Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
// See: http://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
type IpProtocol uint8

const (
	// IpProtoHOPOPT - IPPROTO_HOPOPT - represents the Hop-by-Hop Option Protocol (HOPOPT).
	// See https://www.rfc-editor.org/rfc/rfc2292; https://datatracker.ietf.org/doc/html/rfc8200
	//
	IpProtoHOPOPT IpProtocol = unix.IPPROTO_HOPOPTS

	// IpProtoHOPOPTS - IPPROTO_HOPOPTS - alias of IpProtoHopOpt (IPPROTO_HOPOPT)
	IpProtoHOPOPTS IpProtocol = unix.IPPROTO_HOPOPTS

	// IpProtoIcmp - IPPROTO_ICMP - represents the Internet Control Message Protocol (ICMP).
	// See: https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol
	IpProtoIcmp IpProtocol = unix.IPPROTO_ICMP

	// IpProtoIGMP - IPPROTO_IGMP - represents the Internet Group Management Protocol (IGMP).
	// See: https://en.wikipedia.org/wiki/Internet_Group_Management_Protocol
	IpProtoIGMP IpProtocol = unix.IPPROTO_IGMP

	// IpProtoGGP - IPPROTO_GGP - represents the Gateway-to-Gateway Protocol (GGP).
	// See: https://en.wikipedia.org/wiki/Gateway-to-Gateway_Protocol
	IpProtoGGP IpProtocol = 0x03

	// IpProtoIPv4Encapsulation - IPPROTO_IPIP - represents IPv4 encapsulation within IPv4.
	// See: https://en.wikipedia.org/wiki/IP_in_IP
	IpProtoIPv4Encapsulation IpProtocol = unix.IPPROTO_IPIP

	// IpProtoST - IPPROTO_ST - represents the Stream Protocol.
	// See: https://en.wikipedia.org/wiki/Stream_Protocol
	IpProtoST IpProtocol = 0x05

	// IpProtoTCP - IPPROTO_TCP - represents the Transmission Control Protocol (TCP).
	// See: https://en.wikipedia.org/wiki/Transmission_Control_Protocol
	IpProtoTCP IpProtocol = unix.IPPROTO_TCP

	// IpProtoCBT - IPPROTO_CBT - represents the CBT protocol.
	// See: https://en.wikipedia.org/wiki/Core-Based_Trees
	IpProtoCBT IpProtocol = 0x07

	// IpProtoEGP - IPPROTO_EGP - represents the Exterior Gateway Protocol (EGP).
	// See: https://en.wikipedia.org/wiki/Exterior_Gateway_Protocol
	IpProtoEGP IpProtocol = unix.IPPROTO_EGP

	// IpProtoIGP - IPPROTO_IGP - represents the Interior Gateway Protocol (IGP).
	// Note: This is an unspecified protocol used within autonomous systems.
	IpProtoIGP IpProtocol = 0x09

	// IpProtoBBNRCCMON - IPPROTO_BBN_RCC_MON - represents the BBN RCC Monitoring protocol.
	IpProtoBBNRCCMON IpProtocol = 0x0A

	// IpProtoNVP_II - IPPROTO_NVP_II - represents the Network Voice Protocol.
	// See: https://en.wikipedia.org/wiki/Network_Voice_Protocol
	IpProtoNVP_II IpProtocol = 0x0B

	// IpProtoPUP - IPPROTO_PUP - represents the PARC Universal Packet Protocol.
	// See: https://en.wikipedia.org/wiki/PARC_Universal_Packet
	IpProtoPUP IpProtocol = 0x0C

	// IpProtoARGUS - IPPROTO_ARGUS - represents the ARGUS protocol.
	IpProtoARGUS IpProtocol = 0x0D

	// IpProtoEMCON - IPPROTO_EMCON - represents the Emission Control Protocol.
	IpProtoEMCON IpProtocol = 0x0E

	// IpProtoXNET - IPPROTO_XNET - represents the Cross Net Debugger.
	IpProtoXNET IpProtocol = 0x0F

	// IpProtoCHAOS - IPPROTO_CHAOS - represents the CHAOS protocol.
	// See: https://en.wikipedia.org/wiki/Chaosnet
	IpProtoCHAOS IpProtocol = 0x10

	// IpProtoUDP - IPPROTO_UDP - represents the User Datagram Protocol (UDP).
	// See: https://en.wikipedia.org/wiki/User_Datagram_Protocol
	IpProtoUDP IpProtocol = unix.IPPROTO_UDP

	// IpProtoMUX - IPPROTO_MUX - represents the Multiplexing protocol.
	IpProtoMUX IpProtocol = 0x12

	// IpProtoDCNMEAS - IPPROTO_DCN_MEAS - represents the DCN Measurement Subsystems protocol.
	IpProtoDCNMEAS IpProtocol = 0x13

	// IpProtoHMP - IPPROTO_HMP - represents the Host Monitoring Protocol.
	IpProtoHMP IpProtocol = 0x14

	// IpProtoPRM - IPPROTO_PRM - represents the Packet Radio Measurement protocol.
	IpProtoPRM IpProtocol = 0x15

	// IpProtoXNSIDP - IPPROTO_XNS_IDP - represents the Xerox NS IDP protocol.
	IpProtoXNSIDP IpProtocol = 0x16

	// IpProtoTRUNK1 - IPPROTO_TRUNK1 - represents the Trunk-1 protocol.
	IpProtoTRUNK1 IpProtocol = 0x17

	// IpProtoTRUNK2 - IPPROTO_TRUNK2 - represents the Trunk-2 protocol.
	IpProtoTRUNK2 IpProtocol = 0x18

	// IpProtoLEAF1 - IPPROTO_LEAF1 - represents the Leaf-1 protocol.
	IpProtoLEAF1 IpProtocol = 0x19

	// IpProtoLEAF2 - IPPROTO_LEAF2 - represents the Leaf-2 protocol.
	IpProtoLEAF2 IpProtocol = 0x1A

	// IpProtoRDP - IPPROTO_RDP - represents the Reliable Datagram Protocol.
	// See: https://en.wikipedia.org/wiki/Reliable_Datagram_Protocol
	IpProtoRDP IpProtocol = 0x1B

	// IpProtoIRTP - IPPROTO_IRTP - represents the Internet Reliable Transaction Protocol.
	IpProtoIRTP IpProtocol = 0x1C

	// IpProtoISO_TP4 - IPPROTO_ISO_TP4 - represents the ISO Transport Protocol Class 4.
	IpProtoISO_TP4 IpProtocol = 0x1D

	// IpProtoNETBLT - IPPROTO_NETBLT - represents the Bulk Data Transfer Protocol.
	// See: https://en.wikipedia.org/wiki/NETBLT
	IpProtoNETBLT IpProtocol = 0x1E

	// IpProtoMFENSP - IPPROTO_MFE_NSP - represents the MFE Network Services Protocol.
	IpProtoMFENSP IpProtocol = 0x1F

	// IpProtoMERIT_INP - IPPROTO_MERIT_INP - represents the MERIT Inter-Network Protocol.
	IpProtoMERIT_INP IpProtocol = 0x20

	// IpProtoDCCP - IPPROTO_DCCP - represents the Datagram Congestion Control Protocol.
	// See: https://en.wikipedia.org/wiki/Datagram_Congestion_Control_Protocol
	IpProtoDCCP IpProtocol = unix.IPPROTO_DCCP

	// IpProtoThirdPC - IPPROTO_3PC - represents the Third Party Connect protocol.
	IpProtoThirdPC IpProtocol = 0x22

	// IpProtoIDPR - IPPROTO_IDPR - represents the Inter-Domain Policy Routing Protocol.
	IpProtoIDPR IpProtocol = 0x23

	// IpProtoXTP - IPPROTO_XTP - represents the Xpress Transfer Protocol.
	IpProtoXTP IpProtocol = 0x24

	// IpProtoDDP - IPPROTO_DDP - represents the Datagram Delivery Protocol.
	IpProtoDDP IpProtocol = 0x25

	// IpProtoIDPR_CMTP - IPPROTO_IDPR_CMTP - represents the Control Message Transport Protocol for IDPR.
	IpProtoIDPR_CMTP IpProtocol = 0x26

	// IpProtoTPXX - IPPROTO_TPXX - represents the TP++ Transport Protocol.
	IpProtoTPXX IpProtocol = 0x27

	// IpProtoIL - IPPROTO_IL - represents the IL Transport Protocol.
	IpProtoIL IpProtocol = 0x28

	// IpProtoIPv6Encapsulation - IPPROTO_IPV6 - represents IPv6 encapsulation within IPv4.
	IpProtoIPv6Encapsulation IpProtocol = 0x29

	// IpProtoSDRP - IPPROTO_SDRP - represents the Source Demand Routing Protocol.
	IpProtoSDRP IpProtocol = 0x2A

	// IpProtoIPv6Route - IPPROTO_IPV6_ROUTE - represents IPv6 Routing header.
	IpProtoIPv6Route IpProtocol = 0x2B

	// IpProtoIPv6Frag - IPPROTO_IPV6_FRAG - represents IPv6 Fragmentation header.
	IpProtoIPv6Frag IpProtocol = 0x2C

	// IpProtoIDRP - IPPROTO_IDRP - represents the Inter-Domain Routing Protocol.
	IpProtoIDRP IpProtocol = 0x2D

	// IpProtoRSVP - IPPROTO_RSVP - represents the Resource Reservation Protocol.
	IpProtoRSVP IpProtocol = 0x2E

	// IpProtoGRE - IPPROTO_GRE - represents the Generic Routing Encapsulation protocol.
	// See: https://en.wikipedia.org/wiki/Generic_Routing_Encapsulation
	IpProtoGRE IpProtocol = 0x2F

	// IpProtoMHRP - IPPROTO_MHRP - represents the Mobile Host Routing Protocol.
	IpProtoMHRP IpProtocol = 0x30

	// IpProtoBNA - IPPROTO_BNA - represents the BNA protocol.
	IpProtoBNA IpProtocol = 0x31

	// IpProtoESP - IPPROTO_ESP - represents the Encapsulating Security Payload protocol.
	// See: https://en.wikipedia.org/wiki/IPsec#Encapsulating_Security_Payload
	IpProtoESP IpProtocol = 0x32

	// IpProtoAH - IPPROTO_AH - represents the Authentication Header protocol.
	// See: https://en.wikipedia.org/wiki/IPsec#Authentication_Header
	IpProtoAH IpProtocol = 0x33

	// IpProtoINLSP - IPPROTO_INLSP - represents the Integrated Net Layer Security Protocol.
	IpProtoINLSP IpProtocol = 0x34

	// IpProtoSWIPE - IPPROTO_SWIPE - represents the IP with Encryption protocol.
	IpProtoSWIPE IpProtocol = 0x35

	// IpProtoNARP - IPPROTO_NARP - represents the NBMA Address Resolution Protocol.
	IpProtoNARP IpProtocol = 0x36

	// IpProtoMOBILE - IPPROTO_MOBILE - represents the IP Mobility protocol.
	IpProtoMOBILE IpProtocol = 0x37

	// IpProtoTLSP - IPPROTO_TLSP - represents the Transport Layer Security Protocol.
	IpProtoTLSP IpProtocol = 0x38

	// IpProtoSKIP - IPPROTO_SKIP - represents the SKIP protocol.
	IpProtoSKIP IpProtocol = 0x39

	// IpProtoIPv6ICMP - IPPROTO_ICMPV6 - represents the Internet Control Message Protocol for IPv6 (ICMPv6).
	// See: https://en.wikipedia.org/wiki/ICMPv6
	IpProtoIPv6ICMP IpProtocol = unix.IPPROTO_ICMPV6

	// IpProtoIPv6NoNxt - IPPROTO_IPV6_NONXT - represents the No Next Header for IPv6.
	IpProtoIPv6NoNxt IpProtocol = 0x3B

	// IpProtoIPv6Opts - IPPROTO_IPV6_OPTS - represents the Destination Options for IPv6.
	IpProtoIPv6Opts IpProtocol = 0x3C

	// IpProtoAnyHostInternal -
	IpProtoAnyHostInternal IpProtocol = 61

	// IpProtoCFTP - IPPROTO_CFTP - represents the CFTP protocol.
	IpProtoCFTP IpProtocol = 0x3E

	// IpProtoAnyLocalNetwork -
	IpProtoAnyLocalNetwork IpProtocol = 63

	// IpProtoSatExpak - IPPROTO_SAT_EXPAK - represents the SATNET and Backroom EXPAK protocol.
	IpProtoSatExpak IpProtocol = 0x40

	// IpProtoKRYPTOLAN - IPPROTO_KRYPTOLAN - represents the Kryptolan protocol.
	IpProtoKRYPTOLAN IpProtocol = 0x41

	// IpProtoRVD - IPPROTO_RVD - represents the MIT Remote Virtual Disk protocol.
	IpProtoRVD IpProtocol = 0x42

	// IpProtoIPPC - IPPROTO_IPPC - represents the Internet Pluribus Packet Core protocol.
	IpProtoIPPC IpProtocol = 0x43

	// IpProtoAnyDistributedFs -
	IpProtoAnyDistributedFs IpProtocol = 68

	// IpProtoSATMON - IPPROTO_SAT_MON - represents the SATNET Monitoring protocol.
	IpProtoSATMON IpProtocol = 0x45

	// IpProtoVISA - IPPROTO_VISA - represents the VISA protocol.
	IpProtoVISA IpProtocol = 0x46

	// IpProtoIpcv - IPPROTO_IPCV - represents the Internet Packet Core Utility protocol.
	IpProtoIpcv IpProtocol = 0x47

	// IpProtoCPNX - IPPROTO_CPNX - represents the Computer Protocol Network Executive.
	IpProtoCPNX IpProtocol = 0x48

	// IpProtoCPHB - IPPROTO_CPHB - represents the Computer Protocol Heart Beat.
	IpProtoCPHB IpProtocol = 0x49

	// IpProtoWSN - IPPROTO_WSN - represents the Wang Span Network protocol.
	IpProtoWSN IpProtocol = 0x4A

	// IpProtoPVP - IPPROTO_PVP - represents the Packet Video Protocol.
	IpProtoPVP IpProtocol = 0x4B

	// IpProtoBRSATMON - IPPROTO_BR_SAT_MON - represents the Backroom SATNET Monitoring protocol.
	IpProtoBRSATMON IpProtocol = 0x4C

	// IpProtoSUNND - IPPROTO_SUN_ND - represents the SUN ND protocol.
	IpProtoSUNND IpProtocol = 0x4D

	// IpProtoWBMON - IPPROTO_WB_MON - represents the WIDEBAND Monitoring protocol.
	IpProtoWBMON IpProtocol = 0x4E

	// IpProtoWBEXPAK - IPPROTO_WB_EXPAK - represents the WIDEBAND EXPAK protocol.
	IpProtoWBEXPAK IpProtocol = 0x4F

	// IpProtoISOIP - IPPROTO_ISO_IP - represents the ISO Internet Protocol.
	IpProtoISOIP IpProtocol = 0x50

	// IpProtoVMTP - IPPROTO_VMTP - represents the Versatile Message Transaction Protocol.
	// See: https://en.wikipedia.org/wiki/Versatile_Message_Transaction_Protocol
	IpProtoVMTP IpProtocol = 0x51

	// IpProtoSECUREVMTP - IPPROTO_SECURE_VMTP - represents the SECURE-VMTP protocol.
	IpProtoSECUREVMTP IpProtocol = 0x52

	// IpProtoVINES - IPPROTO_VINES - represents the VINES protocol.
	IpProtoVINES IpProtocol = 0x53

	// IpProtoTTP - IPPROTO_TTP - represents the TTP protocol.
	IpProtoTTP IpProtocol = 0x54

	// IpProtoIptm - IPPROTO_IPTM - represents the IPTM protocol.
	IpProtoIptm IpProtocol = 0x54

	// IpProtoNSFNETIGP - IPPROTO_NSFNET_IGP - represents the NSFNET-IGP protocol.
	IpProtoNSFNETIGP IpProtocol = 0x55

	// IpProtoDGP - IPPROTO_DGP - represents the Dissimilar Gateway Protocol.
	IpProtoDGP IpProtocol = 0x56

	// IpProtoTCF - IPPROTO_TCF - represents the TCF protocol.
	IpProtoTCF IpProtocol = 0x57

	// IpProtoEIGRP - IPPROTO_EIGRP - represents the EIGRP protocol.
	// See: https://en.wikipedia.org/wiki/Enhanced_Interior_Gateway_Routing_Protocol
	IpProtoEIGRP IpProtocol = 0x58

	// IpProtoOSPF - IPPROTO_OSPF - represents the Open Shortest Path First protocol.
	// See: https://en.wikipedia.org/wiki/Open_Shortest_Path_First
	IpProtoOSPF IpProtocol = 0x59

	// IpProtoSPRITRPC - IPPROTO_SPRITE_RPC - represents the Sprite RPC protocol.
	IpProtoSPRITRPC IpProtocol = 0x5A

	// IpProtoLARP - IPPROTO_LARP - represents the Locus Address Resolution Protocol.
	IpProtoLARP IpProtocol = 0x5B

	// IpProtoMTP - IPPROTO_MTP - represents the Multicast Transport Protocol.
	IpProtoMTP IpProtocol = 0x5C

	// IpProtoAX25 - IPPROTO_AX25 - represents the AX.25 protocol.
	// See: https://en.wikipedia.org/wiki/AX.25
	IpProtoAX25 IpProtocol = 0x5D

	// IpProtoIPIP - IPPROTO_IPIP - represents IP-in-IP encapsulation.
	// See: https://en.wikipedia.org/wiki/IP_in_IP
	IpProtoIPIP IpProtocol = 0x5E

	// IpProtoMICP - IPPROTO_MICP - represents the Mobile Internetworking Control Protocol.
	IpProtoMICP IpProtocol = 0x5F

	// IpProtoSCCSP - IPPROTO_SCC_SP - represents the Semaphore Communications Secured protocol.
	IpProtoSCCSP IpProtocol = 0x60

	// IpProtoETHERIP - IPPROTO_ETHERIP - represents the Ethernet within IP Encapsulation.
	IpProtoETHERIP IpProtocol = 0x61

	// IpProtoENCAP - IPPROTO_ENCAP - represents encapsulation within IP.
	IpProtoENCAP IpProtocol = 0x62

	// IpProtoAnyPrivateEncryption -
	IpProtoAnyPrivateEncryption IpProtocol = 99

	// IpProtoGMTP - IPPROTO_GMTP - represents the GMTP protocol.
	IpProtoGMTP IpProtocol = 0x64

	// IpProtoIFMP - IPPROTO_IFMP - represents the Ipsilon Flow Management Protocol.
	IpProtoIFMP IpProtocol = 0x65

	// IpProtoPNNI - IPPROTO_PNNI - represents the PNNI protocol.
	IpProtoPNNI IpProtocol = 0x66

	// IpProtoPIM - IPPROTO_PIM - represents the Protocol Independent Multicast protocol.
	// See: https://en.wikipedia.org/wiki/Protocol_Independent_Multicast
	IpProtoPIM IpProtocol = 0x67

	// IpProtoARIS - IPPROTO_ARIS - represents the ARIS protocol.
	IpProtoARIS IpProtocol = 0x68

	// IpProtoSCPS - IPPROTO_SCPS - represents the SCPS protocol.
	IpProtoSCPS IpProtocol = 0x69

	// IpProtoQNX - IPPROTO_QNX - represents the QNX protocol.
	IpProtoQNX IpProtocol = 0x6A

	// IpProtoAN - IPPROTO_A_N - represents the Active Networks protocol.
	IpProtoAN IpProtocol = 0x6B

	// IpProtoIPComp - IPPROTO_IPCOMP - represents IP Payload Compression.
	// See: https://en.wikipedia.org/wiki/IP_payload_compression_protocol
	IpProtoIPComp IpProtocol = 0x6C

	// IpProtoSNP - IPPROTO_SNP - represents the Sitara Networks Protocol.
	IpProtoSNP IpProtocol = 0x6D

	// IpProtoCompaqPeer - IPPROTO_COMPAQ_PEER - represents the Compaq Peer protocol.
	IpProtoCompaqPeer IpProtocol = 0x6E

	// IpProtoIpxInIp - IPPROTO_IPX_IN_IP - represents IPX encapsulation within IP.
	IpProtoIpxInIp IpProtocol = 0x6F

	// IpProtoVRRP - IPPROTO_VRRP - represents the Virtual Router Redundancy Protocol.
	// See: https://en.wikipedia.org/wiki/Virtual_Router_Redundancy_Protocol
	IpProtoVRRP IpProtocol = 0x70

	// IpProtoPGM - IPPROTO_PGM - represents the Pragmatic General Multicast protocol.
	// See: https://en.wikipedia.org/wiki/Pragmatic_General_Multicast
	IpProtoPGM IpProtocol = 0x71

	// Any0HopProtocol -
	Any0HopProtocol IpProtocol = 0x72

	// IpProtoL2TP - IPPROTO_L2TP - represents the Layer Two Tunneling Protocol.
	// See: https://en.wikipedia.org/wiki/Layer_2_Tunneling_Protocol
	IpProtoL2TP IpProtocol = 0x73

	// IpProtoDDX - IPPROTO_DDX - represents the DDX protocol.
	IpProtoDDX IpProtocol = 0x74

	// IpProtoIATP - IPPROTO_IATP - represents the Interactive Agent Transfer Protocol.
	IpProtoIATP IpProtocol = 0x75

	// IpProtoSTP - IPPROTO_STP - represents the Schedule Transfer Protocol.
	IpProtoSTP IpProtocol = 0x76

	// IpProtoSRP - IPPROTO_SRP - represents the SpectraLink Radio Protocol.
	IpProtoSRP IpProtocol = 0x77

	// IpProtoUTI - IPPROTO_UTI - represents the UTI protocol.
	IpProtoUTI IpProtocol = 0x78

	// IpProtoSMP - IPPROTO_SMP - represents the Simple Message Protocol.
	IpProtoSMP IpProtocol = 0x79

	// IpProtoSM - IPPROTO_SM - represents the SM protocol.
	IpProtoSM IpProtocol = 0x7A

	// IpProtoPTP - IPPROTO_PTP - represents the Performance Transparency Protocol.
	IpProtoPTP IpProtocol = 0x7B

	// IpProtoIsisOverIpv4 -
	IpProtoIsisOverIpv4 IpProtocol = 0x7C

	// IpProtoFIRE - IPPROTO_FIRE - represents the FIRE protocol.
	IpProtoFIRE IpProtocol = 0x7D

	// IpProtoCRTP - IPPROTO_CRTP - represents the Combat Radio Transport Protocol.
	IpProtoCRTP IpProtocol = 0x7E

	// IpProtoCRUDP - IPPROTO_CRUDP - represents the Combat Radio User Datagram protocol.
	IpProtoCRUDP IpProtocol = 0x7F

	// IpProtoSSCOPMCE - IPPROTO_SSCOPMCE - represents the SSCOPMCE protocol.
	IpProtoSSCOPMCE IpProtocol = 0x80

	// IpProtoIPLT - IPPROTO_IPLT - represents the IPLT protocol.
	IpProtoIPLT IpProtocol = 0x81

	// IpProtoSPS - IPPROTO_SPS - represents the SPS protocol.
	IpProtoSPS IpProtocol = 0x82

	// IpProtoPIPE - IPPROTO_PIPE - represents the PIPE protocol.
	IpProtoPIPE IpProtocol = 0x83

	// IpProtoSctp - IPPROTO_SCTP - represents the Stream Control Transmission Protocol (SCTP).
	// See: https://en.wikipedia.org/wiki/Stream_Control_Transmission_Protocol
	IpProtoSctp IpProtocol = unix.IPPROTO_SCTP

	// IpProtoFC - IPPROTO_FC - represents the Fibre Channel protocol.
	IpProtoFC IpProtocol = 0x85

	// IpProtoRsvpE2EIgnore -
	IpProtoRsvpE2EIgnore IpProtocol = 134

	// IpProtoMobilityHeader -
	IpProtoMobilityHeader IpProtocol = 135

	// IpProtoUdpLite -
	IpProtoUdpLite IpProtocol = 136

	// IpProtoMplsInIp -
	IpProtoMplsInIp IpProtocol = 137

	// IpProtoManet - IPPROTO_MANET - represents the MANET protocols.
	IpProtoManet IpProtocol = 0x8A

	// IpProtoHIP - IPPROTO_HIP - represents the Host Identity Protocol.
	// See: https://en.wikipedia.org/wiki/Host_Identity_Protocol
	IpProtoHIP IpProtocol = 0x8B

	// IpProtoShim6 - IPPROTO_SHIM6 - represents the Shim6 protocol.
	// See: https://en.wikipedia.org/wiki/SHIM6
	IpProtoShim6 IpProtocol = 0x8C

	// IpProtoWesp -
	IpProtoWesp IpProtocol = 0x8D

	// IpProtoRohc -
	IpProtoRohc IpProtocol = 0x8E

	// IpProtoEthernet -
	IpProtoEthernet IpProtocol = 0x8F

	// IpProtoAggFrag -
	IpProtoAggFrag IpProtocol = 0x90

	// IpProtoNsh -
	IpProtoNsh IpProtocol = 0x91

	// IpProtoMobility -
	IpProtoMobility IpProtocol = 0x92

	// IpProtoRohcTransport -
	IpProtoRohcTransport IpProtocol = 0x93

	// IpProtoRohcInternal -
	IpProtoRohcInternal IpProtocol = 0x94

	// IpProtoWespInternal -
	IpProtoWespInternal IpProtocol = 0x95

	// IpProtoRohcCompressed -
	IpProtoRohcCompressed IpProtocol = 0x96

	// IpProtoReserved -
	IpProtoReserved IpProtocol = 0xFF
)

// String - converts IpProtocol value to its corresponding protocol name as a string.
//
// In the Internet Protocol version 4 (IPv4) [RFC791] there is a field called "Protocol" to identify the
// next level protocol.  This is an 8 bit (uint8) field.  In Internet Protocol version 6 (IPv6) [RFC8200],
// this field is called the "Next Header" field.
//
// Protocol numbers can be found in the IANA Protocol Numbers registry:
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
//
//	http://en.wikipedia.org/wiki/List_of_IP_protocol_numbers,
//	https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
func (p *IpProtocol) String() string {
	switch *p {
	case IpProtoHOPOPT:
		return "HOPOPT"
	case IpProtoIcmp:
		return "IpProtoIcmp"
	case IpProtoIGMP:
		return "IGMP"
	case IpProtoGGP:
		return "GGP"
	case IpProtoIPv4Encapsulation:
		return "IPv4"
	case IpProtoST:
		return "ST"
	case IpProtoTCP:
		return "TCP"
	case IpProtoCBT:
		return "CBT"
	case IpProtoEGP:
		return "EGP"
	case IpProtoIGP:
		return "IGP"
	case IpProtoBBNRCCMON:
		return "BBN-RCC-MON"
	case IpProtoNVP_II:
		return "NVP-II"
	case IpProtoPUP:
		return "PUP"
	case IpProtoARGUS:
		return "ARGUS"
	case IpProtoEMCON:
		return "EMCON"
	case IpProtoXNET:
		return "XNET"
	case IpProtoCHAOS:
		return "CHAOS"
	case IpProtoUDP:
		return "UDP"
	case IpProtoMUX:
		return "MUX"
	case IpProtoDCNMEAS:
		return "DCN-MEAS"
	case IpProtoHMP:
		return "HMP"
	case IpProtoPRM:
		return "PRM"
	case IpProtoXNSIDP:
		return "XNS-IDP"
	case IpProtoTRUNK1:
		return "TRUNK-1"
	case IpProtoTRUNK2:
		return "TRUNK-2"
	case IpProtoLEAF1:
		return "LEAF-1"
	case IpProtoLEAF2:
		return "LEAF-2"
	case IpProtoRDP:
		return "RDP"
	case IpProtoIRTP:
		return "IRTP"
	case IpProtoISO_TP4:
		return "ISO-TP4"
	case IpProtoNETBLT:
		return "NETBLT"
	case IpProtoMFENSP:
		return "MFE-NSP"
	case IpProtoMERIT_INP:
		return "MERIT-INP"
	case IpProtoDCCP:
		return "DCCP"
	case IpProtoThirdPC:
		return "3PC"
	case IpProtoIDPR:
		return "IDPR"
	case IpProtoXTP:
		return "XTP"
	case IpProtoDDP:
		return "DDP"
	case IpProtoIDPR_CMTP:
		return "IDPR-CMTP"
	case IpProtoTPXX:
		return "TP++"
	case IpProtoIL:
		return "IL"
	case IpProtoIPv6Encapsulation:
		return "IPv6"
	case IpProtoSDRP:
		return "SDRP"
	case IpProtoIPv6Route:
		return "IPv6-Route"
	case IpProtoIPv6Frag:
		return "IPv6-Frag"
	case IpProtoIDRP:
		return "IDRP"
	case IpProtoRSVP:
		return "RSVP"
	case IpProtoGRE:
		return "GRE"
	case IpProtoMHRP:
		return "MHRP"
	case IpProtoBNA:
		return "BNA"
	case IpProtoESP:
		return "ESP"
	case IpProtoAH:
		return "AH"
	case IpProtoINLSP:
		return "I-NLSP"
	case IpProtoSWIPE:
		return "SWIPE"
	case IpProtoNARP:
		return "NARP"
	case IpProtoMOBILE:
		return "MOBILE"
	case IpProtoTLSP:
		return "TLSP"
	case IpProtoSKIP:
		return "SKIP"
	case IpProtoIPv6ICMP:
		return "ICMPv6"
	case IpProtoIPv6NoNxt:
		return "IPv6-NoNxt"
	case IpProtoIPv6Opts:
		return "IPv6-Opts"
	case IpProtoAnyHostInternal:
		return "Any host internal protocol"
	case IpProtoCFTP:
		return "CFTP"
	case IpProtoAnyLocalNetwork:
		return "Any local network"
	case IpProtoSatExpak:
		return "SAT-EXPAK"
	case IpProtoKRYPTOLAN:
		return "KRYPTOLAN"
	case IpProtoRVD:
		return "RVD"
	case IpProtoIPPC:
		return "IPPC"
	case IpProtoAnyDistributedFs:
		return "Any distributed file system"
	case IpProtoSATMON:
		return "SAT-MON"
	case IpProtoVISA:
		return "VISA"
	case IpProtoIpcv:
		return "IPCU"
	case IpProtoCPNX:
		return "CPNX"
	case IpProtoCPHB:
		return "CPHB"
	case IpProtoWSN:
		return "WSN"
	case IpProtoPVP:
		return "PVP"
	case IpProtoBRSATMON:
		return "BR-SAT-MON"
	case IpProtoSUNND:
		return "SUN-ND"
	case IpProtoWBMON:
		return "WB-MON"
	case IpProtoWBEXPAK:
		return "WB-EXPAK"
	case IpProtoISOIP:
		return "ISO-IP"
	case IpProtoVMTP:
		return "VMTP"
	case IpProtoSECUREVMTP:
		return "SECURE-VMTP"
	case IpProtoVINES:
		return "VINES"
	case IpProtoTTP:
		return "TTP"
	case IpProtoNSFNETIGP:
		return "NSFNET-IGP"
	case IpProtoDGP:
		return "DGP"
	case IpProtoTCF:
		return "TCF"
	case IpProtoEIGRP:
		return "EIGRP"
	case IpProtoOSPF:
		return "OSPF"
	case IpProtoSPRITRPC:
		return "Sprite-RPC"
	case IpProtoLARP:
		return "LARP"
	case IpProtoMTP:
		return "MTP"
	case IpProtoAX25:
		return "AX.25"
	case IpProtoIPIP:
		return "IPIP"
	case IpProtoMICP:
		return "MICP"
	case IpProtoSCCSP:
		return "SCC-SP"
	case IpProtoETHERIP:
		return "ETHERIP"
	case IpProtoENCAP:
		return "ENCAP"
	case IpProtoAnyPrivateEncryption:
		return "Any private encryption scheme"
	case IpProtoGMTP:
		return "GMTP"
	case IpProtoIFMP:
		return "IFMP"
	case IpProtoPNNI:
		return "PNNI"
	case IpProtoPIM:
		return "PIM"
	case IpProtoARIS:
		return "ARIS"
	case IpProtoSCPS:
		return "SCPS"
	case IpProtoQNX:
		return "QNX"
	case IpProtoAN:
		return "A/N"
	case IpProtoIPComp:
		return "IPComp"
	case IpProtoSNP:
		return "SNP"
	case IpProtoCompaqPeer:
		return "Compaq-Peer"
	case IpProtoIpxInIp:
		return "IPX-in-IP"
	case IpProtoVRRP:
		return "VRRP"
	case IpProtoPGM:
		return "PGM"
	case Any0HopProtocol:
		return "Any 0-hop protocol"
	case IpProtoL2TP:
		return "L2TP"
	case IpProtoDDX:
		return "DDX"
	case IpProtoIATP:
		return "IATP"
	case IpProtoSTP:
		return "STP"
	case IpProtoSRP:
		return "SRP"
	case IpProtoUTI:
		return "UTI"
	case IpProtoSMP:
		return "SMP"
	case IpProtoSM:
		return "SM"
	case IpProtoPTP:
		return "PTP"
	case IpProtoIsisOverIpv4:
		return "ISIS over IPv4"
	case IpProtoFIRE:
		return "FIRE"
	case IpProtoCRTP:
		return "CRTP"
	case IpProtoCRUDP:
		return "CRUDP"
	case IpProtoSSCOPMCE:
		return "SSCOPMCE"
	case IpProtoIPLT:
		return "IPLT"
	case IpProtoSPS:
		return "SPS"
	case IpProtoPIPE:
		return "PIPE"
	case IpProtoSctp:
		return "SCTP"
	case IpProtoFC:
		return "FC"
	case IpProtoRsvpE2EIgnore:
		return "RSVP-E2E-IGNORE"
	case IpProtoMobilityHeader:
		return "Mobility Header"
	case IpProtoUdpLite:
		return "UDPLite"
	case IpProtoMplsInIp:
		return "MPLS-in-IP"
	case IpProtoManet:
		return "manet"
	case IpProtoHIP:
		return "HIP"
	case IpProtoShim6:
		return "Shim6"
	case IpProtoWesp:
		return "WESP"
	case IpProtoRohc:
		return "ROHC"
	case IpProtoEthernet:
		return "Ethernet"
	case IpProtoAggFrag:
		return "AGGFRAG"
	case IpProtoNsh:
		return "NSH"
	case IpProtoMobility:
		return "MOBILITY"
	case IpProtoRohcTransport:
		return "ROHC-TRANSPORT"
	case IpProtoRohcInternal:
		return "ROHC-INTERNAL"
	case IpProtoWespInternal:
		return "WESP-INTERNAL"
	case IpProtoRohcCompressed:
		return "ROHC-COMPRESSED"
	case IpProtoReserved:
		return "Reserved"
	default:
		return "Unknown"
	}
}

// Serialize converts the IPProto enum value to a byte slice.
// This can be useful when serializing network protocols that require the protocol field
// to be transmitted or used in other operations as a single byte.
//
// For example, this could be used when encoding a protocol identifier for network
// packet manipulation.
//
// Returns a byte slice containing the serialized IPProto.
func (p *IpProtocol) Serialize() ([]byte, error) {
	result := make([]byte, 1)
	result[0] = byte(*p)
	return result, nil
}
