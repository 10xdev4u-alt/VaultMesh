package vault

import (
	"fmt"
	"time"

	"github.com/10xdev4u-alt/VaultMesh/internal/crypto"
)

// Vault represents a shared encrypted storage space.
type Vault struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	OwnerID        string    `json:"owner_id"`
	ParticipantIDs []string  `json:"participant_ids"`
	CreatedAt      time.Time `json:"created_at"`
}

// VaultManager handles the creation and management of vaults.
type VaultManager struct {
	Vaults map[string]*Vault
}

// NewVaultManager creates a new VaultManager.
func NewVaultManager() *VaultManager {
	return &VaultManager{
		Vaults: make(map[string]*Vault),
	}
}

// CreateVault initializes a new collaborative vault.
func (m *VaultManager) CreateVault(id, name, owner string) *Vault {
	v := &Vault{
		ID:        id,
		Name:      name,
		OwnerID:   owner,
		CreatedAt: time.Now(),
	}
	m.Vaults[id] = v
	return v
}

// ReconstructVaultKey uses Shamir's Secret Sharing to rebuild the master key from shards.
func (m *VaultManager) ReconstructVaultKey(shards [][]byte) ([]byte, error) {
	key, err := crypto.CombineShards(shards)
	if err != nil {
		return nil, fmt.Errorf("failed to reconstruct vault key: %w", err)
	}
	return key, nil
}
