package core

// NlMsghdr represents the Netlink message header used in Netlink communication between
// user space and the Linux kernel.
//
// In Netlink communication, each message starts with a header that contains metadata about
// the message, such as its length, type, flags, sequence number, and the sender's process ID.
//
// This Go struct corresponds to the `struct nlmsghdr` defined in the Linux kernel's Netlink
// header files.
//
// Reference:
//   - Linux kernel definition of `struct nlmsghdr`:
//     https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
//
// The structure of `struct nlmsghdr` in the Linux kernel:
//
//	struct nlmsghdr {
//	    __u32 nlmsg_len;   /* Length of message including header */
//	    __u16 nlmsg_type;  /* Message content */
//	    __u16 nlmsg_flags; /* Additional flags */
//	    __u32 nlmsg_seq;   /* Sequence number */
//	    __u32 nlmsg_pid;   /* Sending process port ID */
//	};
//
// Fields:
//   - **Len** (`nlmsg_len`): Total length of the Netlink message, including this header and the payload.
//     This field indicates to the kernel how many bytes to read for this message.
//   - **Type** (`nlmsg_type`): Type of the Netlink message, indicating the message content.
//     Standard message types are defined (e.g., `NLMSG_NOOP`, `NLMSG_ERROR`), and protocol-specific
//     message types start from `NLMSG_MIN_TYPE`.
//   - **Flags** (`nlmsg_flags`): Additional flags that modify the behavior of the message.
//     Flags can be combined using bitwise OR. Common flags include:
//   - `NLM_F_REQUEST`: Must be set on all request messages.
//   - `NLM_F_ACK`: Request an acknowledgment from the kernel.
//   - `NLM_F_MULTI`: Message is part of a multipart message terminated by `NLMSG_DONE`.
//   - **Seq** (`nlmsg_seq`): Sequence number of the message, used to match requests and responses.
//     Applications can use this to track messages.
//   - **Pid** (`nlmsg_pid`): Sending process port ID. In user space, this is usually set to the process ID (`getpid()`).
//     The kernel sets this to `0` when sending messages to user space.
//
// Example Usage:
// ```go
// import (
//
//	"encoding/binary"
//	"bytes"
//	"os"
//
// )
//
//	func createNetlinkMessage(msgType uint16, flags uint16, seq uint32, payload []byte) ([]byte, error) {
//	    nlHeader := NlMsghdr{
//	        Len:   uint32(NlmsgHdrLen + len(payload)),
//	        Type:  msgType,
//	        Flags: flags,
//	        Seq:   seq,
//	        Pid:   uint32(os.Getpid()),
//	    }
//
//	    buf := new(bytes.Buffer)
//	    // Serialize the header
//	    if err := binary.Write(buf, binary.LittleEndian, nlHeader); err != nil {
//	        return nil, err
//	    }
//	    // Append the payload
//	    buf.Write(payload)
//
//	    return buf.Bytes(), nil
//	}
//
// ```
//
// In this example, `createNetlinkMessage` constructs a Netlink message with the given type, flags,
// sequence number, and payload. The header is serialized into bytes in little-endian format, followed
// by the payload.
//
// **Note:** When sending Netlink messages, ensure that the message is aligned to 4-byte boundaries
// as required by Netlink protocol.
//
// **Serialization Considerations:**
// - Netlink uses the host byte order (endianness). On little-endian systems (e.g., x86), this is little-endian.
// - The `binary` package is used for serialization.
// - Ensure that the `Len` field accurately represents the total length of the message.
//
// **References for Serialization:**
// - Netlink message alignment: https://man7.org/linux/man-pages/man7/netlink.7.html
// - Host byte order in Netlink: https://stackoverflow.com/questions/15190228/netlink-message-header-byte-order
type NlMsghdr struct {
	// Len is the total length of the Netlink message, including the header and payload.
	// It corresponds to the `nlmsg_len` field in the kernel's `struct nlmsghdr`.
	Len uint32

	// Type is the message type, indicating the nature of the message content.
	// It corresponds to the `nlmsg_type` field.
	// Standard types are defined in `<linux/netlink.h>`, and protocol-specific types start from `NLMSG_MIN_TYPE`.
	Type uint16

	// Flags are additional flags that modify the behavior of the message.
	// It corresponds to the `nlmsg_flags` field.
	// Flags can be standard (e.g., `NLM_F_REQUEST`, `NLM_F_ACK`) or protocol-specific.
	Flags uint16

	// Seq is the sequence number of the message, used to match requests and responses.
	// It corresponds to the `nlmsg_seq` field.
	// Applications should maintain their own sequence number counter.
	Seq uint32

	// Pid is the port ID of the sending process.
	// It corresponds to the `nlmsg_pid` field.
	// For user-space processes, this is typically set to the process's PID.
	Pid uint32
}
