package core

// TcConnmark - Structure that defines Traffic Control (TC) connection mark attributes.
// This structure extends the generic TC action (`TcGen`) and adds connection mark (connmark) related attributes.
// The connmark action allows matching and marking packets based on connection tracking information (netfilter/iptables).
//
// The `Zone` field is used in connection tracking zones, which enable isolating connection tracking within
// different namespaces or contexts.
//
// For more details on TC connection mark:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_connmark.h
type TcConnmark struct {
	// TcGen - Embed the generic Traffic Control action structure.
	TcGen

	// Zone - Connection tracking zone identifier.
	// This field is used to specify the zone in which the connection mark is applied.
	Zone uint16
}
