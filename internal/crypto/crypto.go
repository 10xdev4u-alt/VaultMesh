// Package crypto provides cryptographic primitives for VaultMesh.
package crypto

import "fmt"

// Cipher defines the interface for encryption and decryption.
type Cipher interface {
	// Encrypt encrypts the plaintext.
	Encrypt(plaintext []byte) ([]byte, error)
	// Decrypt decrypts the ciphertext.
	Decrypt(ciphertext []byte) ([]byte, error)
}

// WrapError provides a common way to wrap errors in the crypto package.
func WrapError(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("crypto %s: %w", op, err)
}
