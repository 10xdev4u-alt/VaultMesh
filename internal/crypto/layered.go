// Package crypto provides cryptographic primitives for VaultMesh.
package crypto

import (
	"log/slog"
)

// LayeredCipher implements a dual encryption pipeline.
// It encrypts the data first with AES-256-GCM and then with ChaCha20-Poly1305.
type LayeredCipher struct {
	aes    *AESGCM
	chacha *ChaCha20Poly1305
}

// NewLayeredCipher creates a new LayeredCipher instance with the given keys.
func NewLayeredCipher(aesKey, chachaKey []byte) (*LayeredCipher, error) {
	aes, err := NewAESGCM(aesKey)
	if err != nil {
		return nil, WrapError("NewLayeredCipher", err)
	}

	chacha, err := NewChaCha20Poly1305(chachaKey)
	if err != nil {
		return nil, WrapError("NewLayeredCipher", err)
	}

	return &LayeredCipher{
		aes:    aes,
		chacha: chacha,
	}, nil
}

// Encrypt encrypts the plaintext using the layered pipeline.
func (l *LayeredCipher) Encrypt(plaintext []byte) ([]byte, error) {
	slog.Debug("Layered encryption: starting AES layer")
	intermediate, err := l.aes.Encrypt(plaintext)
	if err != nil {
		return nil, WrapError("LayeredCipher.Encrypt (AES)", err)
	}

	slog.Debug("Layered encryption: starting ChaCha20 layer")
	ciphertext, err := l.chacha.Encrypt(intermediate)
	if err != nil {
		return nil, WrapError("LayeredCipher.Encrypt (ChaCha20)", err)
	}

	return ciphertext, nil
}

// Decrypt decrypts the ciphertext using the layered pipeline.
func (l *LayeredCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	slog.Debug("Layered decryption: starting ChaCha20 layer")
	intermediate, err := l.chacha.Decrypt(ciphertext)
	if err != nil {
		return nil, WrapError("LayeredCipher.Decrypt (ChaCha20)", err)
	}

	slog.Debug("Layered decryption: starting AES layer")
	plaintext, err := l.aes.Decrypt(intermediate)
	if err != nil {
		return nil, WrapError("LayeredCipher.Decrypt (AES)", err)
	}

	return plaintext, nil
}
