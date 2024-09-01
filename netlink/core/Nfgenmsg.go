//go:build linux

package core

// Nfgenmsg - General form of address family dependent message.
//
//	struct nfgenmsg {
//		__u8  nfgen_family;		/* AF_xxx */
//		__u8  version;		/* nfnetlink version */
//		__be16    res_id;		/* resource id */
//	};
type Nfgenmsg struct {
	NfgenFamily uint8
	Version     uint8
	ResId       uint16 // big endian
}
