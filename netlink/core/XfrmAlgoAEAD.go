package core

// XfrmAlgoAEAD represents the `xfrm_algo_aead` structure used in the Linux kernel's IPsec subsystem.
// This structure is used to specify an AEAD (Authenticated Encryption with Associated Data) algorithm in IPsec.
// AEAD algorithms combine both encryption and authentication in one operation, providing confidentiality and integrity protection.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmAlgoAEAD struct {
	// AlgName represents the name of the AEAD algorithm.
	// This field holds the algorithm name as a null-terminated string, with a maximum size of 64 bytes.
	// Examples include "aes-gcm", "chacha20-poly1305", etc.
	AlgName [64]byte

	// AlgKeyLen represents the length of the cryptographic key in bits.
	// This field defines the size of the key used in the AEAD algorithm.
	// Example: 128 for AES-128-GCM, 256 for AES-256-GCM.
	AlgKeyLen uint32

	// AlgICVLen represents the length of the Integrity Check Value (ICV) in bits.
	// This field defines the size of the ICV used in the AEAD algorithm to provide message integrity.
	// Example: 128 for AES-GCM, which uses a 128-bit ICV.
	AlgICVLen uint32

	// AlgKey holds the actual cryptographic key used by the AEAD algorithm.
	// This field is a variable-length byte slice, and its length corresponds to AlgKeyLen.
	AlgKey []byte
}
