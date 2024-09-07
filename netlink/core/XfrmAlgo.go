package core

// XfrmAlgo represents the `xfrm_algo` structure used in the Linux kernel's IPsec subsystem.
// This structure holds information about cryptographic algorithms used in IPsec, such as encryption or authentication.
// It defines the algorithm name, key length, and the actual cryptographic key.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmAlgo struct {
	// AlgName represents the name of the cryptographic algorithm.
	// This field holds the algorithm name as a null-terminated string, with a maximum size of 64 bytes.
	// Examples: "aes", "sha1", "hmac(sha256)"
	AlgName [64]byte

	// AlgKeyLen represents the length of the cryptographic key in bits.
	// This field defines how many bits are used in the key for this particular algorithm.
	// Example: 128 for AES-128, 256 for AES-256.
	AlgKeyLen uint32

	// AlgKey holds the actual cryptographic key used by the algorithm.
	// This field is a variable-length byte slice, and its length corresponds to AlgKeyLen.
	// The key is used to perform encryption, decryption, or authentication operations.
	AlgKey []byte
}
