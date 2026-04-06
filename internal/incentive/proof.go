package incentive

import (
	"crypto/sha256"
	"time"
)

// StorageChallenge represents a request for a node to prove it has data.
type StorageChallenge struct {
	ChunkID  string
	Salt     []byte
	Deadline time.Time
}

// GenerateProof computes the response to a storage challenge.
func GenerateProof(chunkData []byte, salt []byte) []byte {
	h := sha256.New()
	h.Write(chunkData)
	h.Write(salt)
	return h.Sum(nil)
}

// VerifyProof verifies that the provided proof matches the expected hash.
func VerifyProof(chunkData []byte, salt []byte, providedProof []byte) bool {
	expected := GenerateProof(chunkData, salt)
	return string(expected) == string(providedProof)
}
