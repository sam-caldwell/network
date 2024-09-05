package core

// TcaCsumEnum - Enumeration for Checksum Action (CSUM) attributes.
// These attributes define the different checksum actions that can be used in traffic control.
// They are used to configure how checksums are handled in packets.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
type TcaCsumEnum uint8

const (
	// TcaCsumUnspec - TCA_CSUM_UNSPEC - Unspecified checksum action.
	// This is a placeholder for unspecified or unknown checksum actions.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
	TcaCsumUnspec TcaCsumEnum = iota

	// TcaCsumParms - TCA_CSUM_PARMS - Parameters for checksum action.
	// This is used to specify the parameters for the checksum action.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
	TcaCsumParms TcaCsumEnum = iota

	// TcaCsumTm - TCA_CSUM_TM - Time information for checksum action.
	// This is used to manage timing for checksum actions.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
	TcaCsumTm TcaCsumEnum = iota

	// TcaCsumPad - TCA_CSUM_PAD - Padding for checksum action.
	// This is reserved for padding purposes to ensure proper alignment.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
	TcaCsumPad TcaCsumEnum = iota

	// TcaCsumMax - TCA_CSUM_MAX - Maximum attribute for checksum action.
	// This defines the maximum attribute in the checksum action enumeration.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
	TcaCsumMax = iota - 1
)
