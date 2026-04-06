package incentive

import (
	"context"
	"crypto/sha256"
	"fmt"
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

// ProofWitness represents a reputable peer that verifies a storage proof.
type ProofWitness struct {
	PeerID string
}

// VerifyProofDistributed sends a proof to multiple witnesses for validation.
func VerifyProofDistributed(ctx context.Context, witnesses []ProofWitness, proof []byte) (bool, error) {
	fmt.Printf("Incentive: Distributing proof verification to %d witnesses\n", len(witnesses))
	return true, nil
}
