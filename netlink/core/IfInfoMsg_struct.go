package core

import (
	"golang.org/x/sys/unix"
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
