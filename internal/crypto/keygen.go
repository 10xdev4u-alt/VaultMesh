// Package crypto provides cryptographic primitives for VaultMesh.
package crypto

import (
	"crypto/sha256"
	"io"
	"log/slog"

	"golang.org/x/crypto/hkdf"
)

// DeriveKey derives a unique key from a master key and a salt/info.
// It uses HKDF with SHA-256.
func DeriveKey(masterKey, salt, info []byte, length int) ([]byte, error) {
	slog.Debug("Deriving key using HKDF", "info", string(info), "length", length)
	
	hash := sha256.New
	kdf := hkdf.New(hash, masterKey, salt, info)
	
	derivedKey := make([]byte, length)
	if _, err := io.ReadFull(kdf, derivedKey); err != nil {
		return nil, WrapError("DeriveKey", err)
	}
	
	return derivedKey, nil
}

// DeriveChunkKeys derives per-chunk unique keys for layered encryption.
// It returns an AES key and a ChaCha20 key.
func DeriveChunkKeys(masterKey []byte, chunkID []byte) (aesKey, chachaKey []byte, err error) {
	slog.Debug("Deriving chunk-specific keys", "chunk_id", string(chunkID))

	// Derive AES-256 key (32 bytes)
	aesKey, err = DeriveKey(masterKey, chunkID, []byte("vaultmesh-aes-256-gcm"), 32)
	if err != nil {
		return nil, nil, WrapError("DeriveChunkKeys (AES)", err)
	}

	// Derive ChaCha20 key (32 bytes)
	chachaKey, err = DeriveKey(masterKey, chunkID, []byte("vaultmesh-chacha20-poly1305"), 32)
	if err != nil {
		return nil, nil, WrapError("DeriveChunkKeys (ChaCha20)", err)
	}

	return aesKey, chachaKey, nil
}
