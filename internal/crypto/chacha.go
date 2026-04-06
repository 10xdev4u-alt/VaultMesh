package crypto

import (
	"crypto/rand"
	"fmt"
	"io"
	"log/slog"

	"golang.org/x/crypto/chacha20poly1305"
)

// ChaCha20Poly1305 represents a ChaCha20-Poly1305 cipher.
type ChaCha20Poly1305 struct {
	key []byte
}

// NewChaCha20Poly1305 creates a new ChaCha20Poly1305 instance with the given 32-byte key.
func NewChaCha20Poly1305(key []byte) (*ChaCha20Poly1305, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, WrapError("NewChaCha20Poly1305", fmt.Errorf("invalid key size: %d, expected 32 bytes", len(key)))
	}
	return &ChaCha20Poly1305{key: key}, nil
}

// Encrypt encrypts the plaintext using ChaCha20-Poly1305.
func (c *ChaCha20Poly1305) Encrypt(plaintext []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(c.key)
	if err != nil {
		return nil, WrapError("ChaCha20Poly1305.Encrypt", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, WrapError("ChaCha20Poly1305.Encrypt", err)
	}

	slog.Debug("ChaCha20-Poly1305 encryption", "nonce_size", len(nonce))

	// Seal appends the ciphertext to the nonce
	return aead.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts the ciphertext using ChaCha20-Poly1305.
func (c *ChaCha20Poly1305) Decrypt(ciphertext []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(c.key)
	if err != nil {
		return nil, WrapError("ChaCha20Poly1305.Decrypt", err)
	}

	nonceSize := aead.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, WrapError("ChaCha20Poly1305.Decrypt", fmt.Errorf("ciphertext too short"))
	}

	nonce, encryptedData := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aead.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, WrapError("ChaCha20Poly1305.Decrypt", err)
	}

	return plaintext, nil
}
