package network

import (
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
)

// IdentityManager handles the generation and persistence of the node's private key.
type IdentityManager struct {
	keyPath string
}

// NewIdentityManager creates a new IdentityManager.
func NewIdentityManager(keyPath string) *IdentityManager {
	return &IdentityManager{keyPath: keyPath}
}

// LoadOrGenerateKey attempts to load a private key from disk, or generates a new one if it doesn't exist.
func (m *IdentityManager) LoadOrGenerateKey() (crypto.PrivKey, error) {
	if _, err := os.Stat(m.keyPath); os.IsNotExist(err) {
		return m.generateAndSaveKey()
	}

	data, err := os.ReadFile(m.keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read identity key: %w", err)
	}

	priv, err := crypto.UnmarshalPrivateKey(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal identity key: %w", err)
	}

	return priv, nil
}

func (m *IdentityManager) generateAndSaveKey() (crypto.PrivKey, error) {
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, 0, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate keypair: %w", err)
	}

	data, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	if err := os.WriteFile(m.keyPath, data, 0600); err != nil {
		return nil, fmt.Errorf("failed to save identity key: %w", err)
	}

	return priv, nil
}
