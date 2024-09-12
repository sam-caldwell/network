package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
)

// IfInfoMsg - wrapper for unix.IfInfoMsg (rtnetlink message for links and list requests)
//
// Creating this wrapper to protect against changes to the underlying golang code and such.
//
// Expectation from golang:
//
//	 type IfInfoMsg struct {
//		  Family uint8
//		  _      uint8
//	      Type   uint16
//		  Index  int32
//		  Flags  uint32
//		  Change uint32
//	 }
//
// In C (https://linux.die.net/man/7/rtnetlink):
//
//	 struct ifinfomsg {
//	   unsigned char  ifi_family; /* AF_UNSPEC */
//	   unsigned short ifi_type;   /* Device type */
//	   int            ifi_index;  /* Interface index */
//	   unsigned int   ifi_flags;  /* Device flags  */
//	   unsigned int   ifi_change; /* change mask */
//	};
type IfInfoMsg struct {
	unix.IfInfomsg
}

// SizeOfIfInfoMsg -  Size of the IfInfoMsg struct
const SizeOfIfInfoMsg = int(unsafe.Sizeof(IfInfoMsg{}))

// Len - Return IfInfoMsg length
func (msg *IfInfoMsg) Len() int {
	return SizeOfIfInfoMsg
}

// Serialize - Serialize the IfInfoMsg structure into a byte slice in a safe way.
func (msg *IfInfoMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Write each field to the buffer in the correct order and format
	if err := binary.Write(buf, NativeEndian, msg.Family); err != nil {
		return nil, err
	}
	// Placeholder for the second byte, typically padding
	if err := binary.Write(buf, NativeEndian, uint8(0)); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Type); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Flags); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Change); err != nil {
		return nil, err
	}
	// Return the byte slice
	return buf.Bytes(), nil
}

// Deserialize - Deserialize byte slice into IfInfoMsg struct
func (msg *IfInfoMsg) Deserialize(b []byte) (err error) {
	var result *IfInfoMsg
	if result, err = DeserializeIfInfoMsg(b); err != nil {
		return err
	}
	*msg = *result
	return nil
}

// DeserializeIfInfoMsg - Deserialize byte slice into IfInfoMsg
//
//	type IfInfoMsg struct {
//	   Family uint8
//	   _      uint8
//	   Type   uint16
//	   Index  int32
//	   Flags  uint32
//	   Change uint32
//	}
func DeserializeIfInfoMsg(b []byte) (*IfInfoMsg, error) {
	if len(b) < SizeOfIfInfoMsg {
		return nil, errors.New("IfInfoMsg to short")
	}
	result := IfInfoMsg{
		unix.IfInfomsg{
			// byte 0
			Family: b[0],
			// byte 1 is skipped
			// bytes 2-3: Type
			Type: NativeEndian.Uint16(b[2:4]),
			// bytes 4-7: Index
			Index: int32(NativeEndian.Uint32(b[4:8])),
			// bytes 8-11: Flags
			Flags: NativeEndian.Uint32(b[8:12]),
			// bytes 12-15: Change
			Change: NativeEndian.Uint32(b[12:16]),
		},
	}
	return &result, nil
}

// EncapType - Return L2 Encapsulation type
func (msg *IfInfoMsg) EncapType() string {
	switch msg.Type {
	case 0:
		return "generic"
	case unix.ARPHRD_ETHER:
		return "ether"
	case unix.ARPHRD_EETHER:
		return "eether"
	case unix.ARPHRD_AX25:
		return "ax25"
	case unix.ARPHRD_PRONET:
		return "pronet"
	case unix.ARPHRD_CHAOS:
		return "chaos"
	case unix.ARPHRD_IEEE802:
		return "ieee802"
	case unix.ARPHRD_ARCNET:
		return "arcnet"
	case unix.ARPHRD_APPLETLK:
		return "atalk"
	case unix.ARPHRD_DLCI:
		return "dlci"
	case unix.ARPHRD_ATM:
		return "atm"
	case unix.ARPHRD_METRICOM:
		return "metricom"
	case unix.ARPHRD_IEEE1394:
		return "ieee1394"
	case unix.ARPHRD_INFINIBAND:
		return "infiniband"
	case unix.ARPHRD_SLIP:
		return "slip"
	case unix.ARPHRD_CSLIP:
		return "cslip"
	case unix.ARPHRD_SLIP6:
		return "slip6"
	case unix.ARPHRD_CSLIP6:
		return "cslip6"
	case unix.ARPHRD_RSRVD:
		return "rsrvd"
	case unix.ARPHRD_ADAPT:
		return "adapt"
	case unix.ARPHRD_ROSE:
		return "rose"
	case unix.ARPHRD_X25:
		return "x25"
	case unix.ARPHRD_HWX25:
		return "hwx25"
	case unix.ARPHRD_PPP:
		return "ppp"
	case unix.ARPHRD_HDLC:
		return "hdlc"
	case unix.ARPHRD_LAPB:
		return "lapb"
	case unix.ARPHRD_DDCMP:
		return "ddcmp"
	case unix.ARPHRD_RAWHDLC:
		return "rawhdlc"
	case unix.ARPHRD_TUNNEL:
		return "ipip"
	case unix.ARPHRD_TUNNEL6:
		return "tunnel6"
	case unix.ARPHRD_FRAD:
		return "frad"
	case unix.ARPHRD_SKIP:
		return "skip"
	case unix.ARPHRD_LOOPBACK:
		return "loopback"
	case unix.ARPHRD_LOCALTLK:
		return "ltalk"
	case unix.ARPHRD_FDDI:
		return "fddi"
	case unix.ARPHRD_BIF:
		return "bif"
	case unix.ARPHRD_SIT:
		return "sit"
	case unix.ARPHRD_IPDDP:
		return "ip/ddp"
	case unix.ARPHRD_IPGRE:
		return "gre"
	case unix.ARPHRD_PIMREG:
		return "pimreg"
	case unix.ARPHRD_HIPPI:
		return "hippi"
	case unix.ARPHRD_ASH:
		return "ash"
	case unix.ARPHRD_ECONET:
		return "econet"
	case unix.ARPHRD_IRDA:
		return "irda"
	case unix.ARPHRD_FCPP:
		return "fcpp"
	case unix.ARPHRD_FCAL:
		return "fcal"
	case unix.ARPHRD_FCPL:
		return "fcpl"
	case unix.ARPHRD_FCFABRIC:
		return "fcfb0"
	case unix.ARPHRD_FCFABRIC + 1:
		return "fcfb1"
	case unix.ARPHRD_FCFABRIC + 2:
		return "fcfb2"
	case unix.ARPHRD_FCFABRIC + 3:
		return "fcfb3"
	case unix.ARPHRD_FCFABRIC + 4:
		return "fcfb4"
	case unix.ARPHRD_FCFABRIC + 5:
		return "fcfb5"
	case unix.ARPHRD_FCFABRIC + 6:
		return "fcfb6"
	case unix.ARPHRD_FCFABRIC + 7:
		return "fcfb7"
	case unix.ARPHRD_FCFABRIC + 8:
		return "fcfb8"
	case unix.ARPHRD_FCFABRIC + 9:
		return "fcfb9"
	case unix.ARPHRD_FCFABRIC + 10:
		return "fcfb10"
	case unix.ARPHRD_FCFABRIC + 11:
		return "fcfb11"
	case unix.ARPHRD_FCFABRIC + 12:
		return "fcfb12"
	case unix.ARPHRD_IEEE802_TR:
		return "tr"
	case unix.ARPHRD_IEEE80211:
		return "ieee802.11"
	case unix.ARPHRD_IEEE80211_PRISM:
		return "ieee802.11/prism"
	case unix.ARPHRD_IEEE80211_RADIOTAP:
		return "ieee802.11/radiotap"
	case unix.ARPHRD_IEEE802154:
		return "ieee802.15.4"

	case 65534:
		return "none"
	case 65535:
		return "void"
	}
	return fmt.Sprintf("unknown%d", msg.Type)
}
