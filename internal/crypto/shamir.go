// Package crypto provides cryptographic primitives for VaultMesh.
package crypto

import (
	"fmt"
	"log/slog"

	"github.com/hashicorp/vault/shamir"
)

// SplitKey splits a key into n shares, requiring k shares to reconstruct.
func SplitKey(key []byte, n, k int) ([][]byte, error) {
	if k > n {
		return nil, WrapError("SplitKey", fmt.Errorf("k cannot be greater than n (k=%d, n=%d)", k, n))
	}

	slog.Info("Splitting key into shares", "n", n, "k", k)
	shares, err := shamir.Split(key, n, k)
	if err != nil {
		return nil, WrapError("SplitKey", err)
	}

	return shares, nil
}

// RecombineShares reconstructs a key from a set of shares.
func RecombineShares(shares [][]byte) ([]byte, error) {
	slog.Info("Recombining key from shares", "num_shares", len(shares))
	key, err := shamir.Combine(shares)
	if err != nil {
		return nil, WrapError("RecombineShares", err)
	}

	return key, nil
}
