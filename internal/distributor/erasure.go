package distributor

import (
	"bytes"
	"fmt"
	"io"

	"github.com/klauspost/reedsolomon"
)

// ErasureCoder handles the encoding and decoding of data into redundant shards.
type ErasureCoder struct {
	dataShards   int
	parityShards int
	enc          reedsolomon.Encoder
}

// NewErasureCoder creates a new ErasureCoder with the specified shard configuration.
func NewErasureCoder(data, parity int) (*ErasureCoder, error) {
	enc, err := reedsolomon.New(data, parity)
	if err != nil {
		return nil, fmt.Errorf("failed to create reedsolomon encoder: %w", err)
	}
	return &ErasureCoder{
		dataShards:   data,
		parityShards: parity,
		enc:          enc,
	}, nil
}

// Encode splits data into shards and generates parity shards.
func (c *ErasureCoder) Encode(data []byte) ([][]byte, error) {
	shards, err := c.enc.Split(data)
	if err != nil {
		return nil, fmt.Errorf("failed to split data into shards: %w", err)
	}

	if err := c.enc.Encode(shards); err != nil {
		return nil, fmt.Errorf("failed to encode parity shards: %w", err)
	}

	return shards, nil
}

// Reconstruct attempts to rebuild the original data from the available shards and joins them.
func (c *ErasureCoder) Reconstruct(shards [][]byte, originalSize int) ([]byte, error) {
	// Reconstruct missing shards (if any)
	if err := c.enc.Reconstruct(shards); err != nil {
		return nil, fmt.Errorf("failed to reconstruct shards: %w", err)
	}

	// Verify integrity
	ok, err := c.enc.Verify(shards)
	if err != nil || !ok {
		return nil, fmt.Errorf("erasure coding verification failed")
	}

	// Join shards back into original data
	var buf bytes.Buffer
	if err := c.enc.Join(&buf, shards, originalSize); err != nil {
		return nil, fmt.Errorf("failed to join shards: %w", err)
	}

	return buf.Bytes(), nil
}
