package core

// TcaBpfEnum - Enumeration of Traffic Control BPF (Berkeley Packet Filter) actions.
//
// This enumeration defines various parameters and options for the BPF actions within the Linux Traffic Control (TC) system.
// BPF programs can be used to filter network traffic and implement various rules on the packets flowing through the network.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
type TcaBpfEnum uint8

const (
	// TcaBpfUnspec - TCA_BPF_UNSPEC - Unspecified BPF action.
	// This is used when no specific action is provided.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfUnspec TcaBpfEnum = iota

	// TcaBpfAct - TCA_BPF_ACT - Specifies the action associated with the BPF filter.
	// Defines the action to be taken on matching packets.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfAct TcaBpfEnum = iota

	// TcaBpfPolice - TCA_BPF_POLICE - Defines policing rules for the BPF action.
	// Policing is used to enforce traffic rate limiting on packets.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfPolice TcaBpfEnum = iota

	// TcaBpfClassid - TCA_BPF_CLASSID - Class identifier for the BPF action.
	// Specifies the traffic class this BPF action is associated with.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfClassid TcaBpfEnum = iota

	// TcaBpfOpsLen - TCA_BPF_OPS_LEN - Length of the BPF operations.
	// This field defines the size of the BPF operations in bytes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfOpsLen TcaBpfEnum = iota

	// TcaBpfOps - TCA_BPF_OPS - BPF program operations.
	// Contains the actual BPF bytecode or instructions to be executed.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfOps TcaBpfEnum = iota

	// TcaBpfFd - TCA_BPF_FD - File descriptor for BPF program.
	// Refers to the file descriptor of the BPF program in the Linux system.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfFd TcaBpfEnum = iota

	// TcaBpfName - TCA_BPF_NAME - Name of the BPF program.
	// This is the human-readable name associated with the BPF program.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfName TcaBpfEnum = iota

	// TcaBpfFlags - TCA_BPF_FLAGS - Flags associated with the BPF program.
	// These flags modify the behavior of the BPF program execution.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfFlags TcaBpfEnum = iota

	// TcaBpfFlagsGen - TCA_BPF_FLAGS_GEN - General flags for the BPF program.
	// This is an additional set of flags that control specific BPF features.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfFlagsGen TcaBpfEnum = iota

	// TcaBpfTag - TCA_BPF_TAG - Tag for the BPF program.
	// A unique tag or identifier for the BPF program, used to differentiate programs.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfTag TcaBpfEnum = iota

	// TcaBpfId - TCA_BPF_ID - Identifier for the BPF program.
	// This ID is used for identifying the BPF program within the system.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfId TcaBpfEnum = iota

	// TcaBpfMax - TCA_BPF_MAX - Maximum valid BPF action value.
	// Represents the highest value in this enumeration.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaBpfMax = iota - 1
)
