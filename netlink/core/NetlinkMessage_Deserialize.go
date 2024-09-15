package core

// DeserializeNetlinkMessage deserializes a Netlink message header and data from a byte slice.
//
// This function parses a Netlink message from the provided byte slice `b`. It extracts the Netlink
// message header (`NlMsghdr`) and returns the remaining data, the aligned message length, and any
// error encountered during the process.
//
// **Function Workflow:**
//
// 1. **Input Size Check:**
//   - Verifies that the input byte slice `b` is at least the size of a Netlink message header
//     (`NetlinkMessageHeaderSize`).
//   - Uses `checkInputSize` to perform this validation.
//   - If the check fails, the function returns an error.
//
// 2. **Header Deserialization:**
//   - Calls `DeserializeNetlinkMessageHeader` to deserialize the Netlink message header from the
//     byte slice `b`.
//   - If deserialization fails, the function returns an error.
//
// 3. **Aligned Message Length Calculation:**
//   - Calculates the aligned message length using `nlmAlignOf(int(header.Len))`.
//   - Netlink messages are aligned to 4-byte boundaries for proper memory alignment.
//
// 4. **Header Length and Data Alignment Validation:**
//   - Checks if the header length (`header.Len`) is less than the minimum header size.
//   - Ensures that the length of `b` is a multiple of 4 bytes to satisfy alignment requirements.
//   - If validation fails, the function returns an `EINVAL` error indicating an invalid argument.
//
// 5. **Return Values:**
//   - On success, returns the deserialized header, the remaining data after the header,
//     the aligned message length, and `nil` for the error.
//
// **Parameters:**
//
// - `b []byte`:
//   - The byte slice containing the serialized Netlink message.
//
// **Returns:**
//
// - `header *NlMsghdr`:
//   - Pointer to the deserialized Netlink message header.
//
// - `remainingData []byte`:
//   - The remaining data in the message after the header.
//
// - `messageLength int`:
//   - The aligned length of the Netlink message.
//
// - `err error`:
//   - Any error encountered during deserialization.
//
// **Related Linux Kernel Sources:**
//
// - Definition of `struct nlmsghdr` in `<linux/netlink.h>`:
//   - [GitHub Source Link](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L50)
//   - This structure defines the format of Netlink message headers used in communication between
//     user space and the kernel.
//
// **Notes:**
//
// - **Netlink Message Structure:**
//   - A Netlink message consists of a `nlmsghdr` followed by the payload data.
//   - The `nlmsghdr` includes fields like message length, type, flags, sequence number, and process ID.
//
// - **Alignment:**
//   - Netlink messages must be aligned to 4-byte boundaries.
//   - The function uses `nlmAlignOf` to ensure proper alignment.
//
// - **Byte Order:**
//   - Netlink messages use the host byte order (endianness).
//   - Serialization and deserialization should consider the system's endianness.
//
// **References:**
//
// - **Netlink Man Page:**
//   - [netlink(7) - Linux manual page](https://man7.org/linux/man-pages/man7/netlink.7.html)
//   - Provides detailed information about Netlink sockets, message formats, and communication.
//
// - **Understanding Netlink Message Headers:**
//   - [Linux Kernel Documentation](https://www.kernel.org/doc/Documentation/networking/netlink.txt)
//   - Offers insights into Netlink message structures and usage patterns.
//
// **Example Usage:**
//
// ```go
// data := /* byte slice containing Netlink messages */
// header, remainingData, messageLength, err := DeserializeNetlinkMessage(data)
//
//	if err != nil {
//	    // Handle error
//	}
//
// // Process the header and remainingData
// ```
//
// **Error Handling:**
//
// - The function returns `EINVAL` if the header length is invalid or the data is not properly aligned.
// - Other errors may be returned from `checkInputSize` or `DeserializeNetlinkMessageHeader`.
//
// **Function Definition:**
func DeserializeNetlinkMessage(b []byte) (header *NlMsghdr, remainingData []byte, messageLength int, err error) {

	// Step 1: Check if the input data is at least the size of the Netlink message header.
	if err = checkInputSize(b, NetlinkMessageHeaderSize, disableSizeCheck); err != nil {
		return nil, nil, 0, err
	}

	// Step 2: Deserialize the Netlink message header from the input byte slice.
	if header, err = DeserializeNetlinkMessageHeader(b); err != nil {
		return nil, nil, 0, err
	}

	// Step 3: Calculate the aligned length of the Netlink message.
	// Netlink messages are aligned to 4-byte boundaries.
	messageLength = nlmAlignOf(int(header.Len))

	// Step 4: Validate the header length and data alignment.
	// - The header length should be at least the size of the Netlink message header.
	// - The length of the input data should be a multiple of 4 bytes.
	if int(header.Len) < NetlinkMessageHeaderSize || len(b)%4 != 0 {
		return nil, nil, 0, EINVAL // EINVAL indicates an invalid argument.
	}

	// Step 5: Return the deserialized header, remaining data, and message length.
	return header, b[NetlinkMessageHeaderSize:], messageLength, err
}
