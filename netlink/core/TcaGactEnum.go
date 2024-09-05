package core

// TcaGactEnum - tc-gact - Enumeration of Generic Actions (GACT) for Traffic Control in Linux.
//
// GACT is a type of action that can be applied to network packets in the Linux Traffic Control system. These actions
// help determine what should be done with a packet, such as dropping it, passing it through, or performing probabilistic actions.
//
// The following constants are used in Linux's Traffic Control system for generic actions.
//
// See:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_gact.h
// - https://man7.org/linux/man-pages/man8/tc-gact.8.html
type TcaGactEnum uint8

const (
	// TcaGactUnspec - TCA_GACT_UNSPEC - Unspecified generic action.
	// This constant is used when no specific action is defined for GACT.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_gact.h
	TcaGactUnspec TcaGactEnum = iota

	// TcaGactTm - TCA_GACT_TM - Timing information for the action.
	// This is used to represent timing metadata for the GACT action.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_gact.h
	TcaGactTm TcaGactEnum = iota

	// TcaGactParms - TCA_GACT_PARMS - Action parameters.
	// This constant represents the parameters for the generic action, such as what action to perform.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_gact.h
	TcaGactParms TcaGactEnum = iota

	// TcaGactProb - TCA_GACT_PROB - Probability-based action.
	// This represents a probabilistic action, where the decision to drop or pass a packet is based on probability.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_gact.h
	TcaGactProb TcaGactEnum = iota

	// TcaGactMax - TCA_GACT_MAX - Maximum valid GACT value.
	// This constant marks the highest valid value for the TcaGactEnum enumeration.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_gact.h
	TcaGactMax = iota - 1
)
