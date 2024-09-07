package core

// XfrmAlgoAuth represents the `xfrm_algo_auth` structure in the Linux kernel's IPsec subsystem.
// This structure is used for specifying cryptographic authentication algorithms in IPsec, such as HMAC or SHA-based algorithms.
// It includes the algorithm name, key length, truncation length, and the actual cryptographic key.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmAlgoAuth struct {
	// AlgName represents the name of the cryptographic authentication algorithm.
	// This field holds the algorithm name as a null-terminated string, with a maximum size of 64 bytes.
	// Examples: "hmac(sha1)", "hmac(sha256)"
	AlgName [64]byte

	// AlgKeyLen represents the length of the cryptographic key in bits.
	// This field defines how many bits are used in the key for this particular authentication algorithm.
	// Example: 128 for HMAC-SHA-128, 256 for HMAC-SHA-256.
	AlgKeyLen uint32

	// AlgTruncLen represents the truncation length in bits.
	// Many cryptographic authentication algorithms produce a longer output than is needed, and this field defines
	// the number of bits to be used after truncation.
	// Example: 96 for HMAC-SHA-96 (which truncates the 160-bit output of SHA1 to 96 bits).
	AlgTruncLen uint32

	// AlgKey holds the actual cryptographic key used by the algorithm.
	// This field is a variable-length byte slice, and its length corresponds to AlgKeyLen.
	// The key is used to perform authentication operations, such as HMAC.
	AlgKey []byte
}
