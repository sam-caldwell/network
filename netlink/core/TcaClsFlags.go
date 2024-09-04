package core

// TcaClsFlags - Traffic Control Classifier Flags.
//
// These flags are used in the Linux kernel traffic control system to specify how filters should be handled.
// They indicate whether a filter should be offloaded to hardware, only used in software, or provide additional
// logging.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
type TcaClsFlags uint32

const (

	// TcaClsFlagsSkipHw - TCA_CLS_FLAGS_SKIP_HW - don't offload filter to hardware.
	// This flag prevents the filter from being sent to hardware for offloading.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaClsFlagsSkipHw TcaClsFlags = 1 << 0

	// TcaClsFlagsSkipSw - TCA_CLS_FLAGS_SKIP_SW - don't use filter in software.
	// This flag prevents the filter from being applied in the software stack.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaClsFlagsSkipSw TcaClsFlags = 1 << 1

	// TcaClsFlagsInHw - TCA_CLS_FLAGS_IN_HW - filter is offloaded to hardware.
	// Indicates that the filter has been successfully offloaded to hardware.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaClsFlagsInHw TcaClsFlags = 1 << 2

	// TcaClsFlagsNotInHw - TCA_CLS_FLAGS_NOT_IN_HW - filter isn't offloaded to hardware.
	// Indicates that the filter has not been offloaded to hardware.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaClsFlagsNotInHw TcaClsFlags = 1 << 3

	// TcaClsFlagsVerbose - TCA_CLS_FLAGS_VERBOSE - verbose logging.
	// Enables detailed logging for traffic control filters.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaClsFlagsVerbose TcaClsFlags = 1 << 4
)
