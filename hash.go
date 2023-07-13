package merkletree

import "crypto/sha256"

// Hash provides a default hash function for Merkle trees that uses truncated SHA256 for 128bit security.
func Hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:16]
}
