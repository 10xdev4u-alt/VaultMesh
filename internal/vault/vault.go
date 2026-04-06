package vault

import (
	"time"
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
