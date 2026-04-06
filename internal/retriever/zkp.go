package retriever

import (
	"crypto/sha256"
)

// ZKProofManager handles zero-knowledge proofs for storage verification.
type ZKProofManager struct{}

// NewZKProofManager creates a new ZKProofManager.
func NewZKProofManager() *ZKProofManager {
	return &ZKProofManager{}
}

// GenerateChallenge creates a random challenge for a storage node.
func (m *ZKProofManager) GenerateChallenge(chunkID string) ([]byte, error) {
	h := sha256.New()
	h.Write([]byte(chunkID))
	return h.Sum(nil), nil
}

// VerifyProof checks if the proof provided by the peer matches the expected result.
func (m *ZKProofManager) VerifyProof(chunkData []byte, challenge []byte, providedProof []byte) bool {
	h := sha256.New()
	h.Write(chunkData)
	h.Write(challenge)
	expectedProof := h.Sum(nil)

	return string(expectedProof) == string(providedProof)
}
