// Package crypto provides cryptographic primitives for VaultMesh.
package crypto

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
)

// KeyStore manages the master key and other cryptographic materials.
type KeyStore struct {
	mu        sync.RWMutex
	MasterKey []byte `json:"master_key"`
}

// NewKeyStore generates a new KeyStore with a random 32-byte master key.
func NewKeyStore() (*KeyStore, error) {
	masterKey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, masterKey); err != nil {
		return nil, WrapError("NewKeyStore", err)
	}

	return &KeyStore{
		MasterKey: masterKey,
	}, nil
}

// Save encrypts and saves the KeyStore to a file.
func (ks *KeyStore) Save(path string, password string) error {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	data, err := json.Marshal(ks)
	if err != nil {
		return WrapError("KeyStore.Save", err)
	}

	// Use password to derive a key for AES-GCM
	key, err := DeriveKey([]byte(password), []byte("vaultmesh-keystore-salt"), []byte("keystore-encryption"), 32)
	if err != nil {
		return WrapError("KeyStore.Save", err)
	}

	cipher, err := NewAESGCM(key)
	if err != nil {
		return WrapError("KeyStore.Save", err)
	}

	encrypted, err := cipher.Encrypt(data)
	if err != nil {
		return WrapError("KeyStore.Save", err)
	}

	if err := os.WriteFile(path, encrypted, 0600); err != nil {
		return WrapError("KeyStore.Save", err)
	}

	slog.Info("Keystore saved successfully", "path", path)
	return nil
}

// LoadKeyStore loads and decrypts a KeyStore from a file.
func LoadKeyStore(path string, password string) (*KeyStore, error) {
	encrypted, err := os.ReadFile(path)
	if err != nil {
		return nil, WrapError("LoadKeyStore", err)
	}

	// Use password to derive the same key
	key, err := DeriveKey([]byte(password), []byte("vaultmesh-keystore-salt"), []byte("keystore-encryption"), 32)
	if err != nil {
		return nil, WrapError("LoadKeyStore", err)
	}

	cipher, err := NewAESGCM(key)
	if err != nil {
		return nil, WrapError("LoadKeyStore", err)
	}

	data, err := cipher.Decrypt(encrypted)
	if err != nil {
		return nil, WrapError("LoadKeyStore", fmt.Errorf("failed to decrypt keystore (wrong password?): %w", err))
	}

	var ks KeyStore
	if err := json.Unmarshal(data, &ks); err != nil {
		return nil, WrapError("LoadKeyStore", err)
	}

	slog.Info("Keystore loaded successfully", "path", path)
	return &ks, nil
}
