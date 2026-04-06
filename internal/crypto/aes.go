package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log/slog"
)

// AESGCM represents an AES-256-GCM cipher.
type AESGCM struct {
	key []byte
}

// NewAESGCM creates a new AESGCM instance with the given 32-byte key.
func NewAESGCM(key []byte) (*AESGCM, error) {
	if len(key) != 32 {
		return nil, WrapError("NewAESGCM", fmt.Errorf("invalid key size: %d, expected 32 bytes for AES-256", len(key)))
	}
	return &AESGCM{key: key}, nil
}

// Encrypt encrypts the plaintext using AES-256-GCM.
func (a *AESGCM) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, WrapError("AESGCM.Encrypt", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, WrapError("AESGCM.Encrypt", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, WrapError("AESGCM.Encrypt", err)
	}

	slog.Debug("AES-GCM encryption", "nonce_size", len(nonce))

	// Seal appends the ciphertext to the nonce
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts the ciphertext using AES-256-GCM.
func (a *AESGCM) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, WrapError("AESGCM.Decrypt", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, WrapError("AESGCM.Decrypt", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, WrapError("AESGCM.Decrypt", fmt.Errorf("ciphertext too short"))
	}

	nonce, encryptedData := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, WrapError("AESGCM.Decrypt", err)
	}

	return plaintext, nil
}
