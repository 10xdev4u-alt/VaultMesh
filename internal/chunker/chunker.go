package chunker

import (
	"bytes"
	"fmt"
	"io"

	"github.com/zeebo/blake3"
)

// Chunker is the common interface for different chunking strategies.
type Chunker interface {
	// Split divides the input reader into chunks.
	Split(r io.Reader) ([][]byte, error)
}

// Type represents the type of chunking algorithm to use.
type Type string

const (
	// FixedSize uses a fixed chunk size.
	FixedSize Type = "fixed"
	// CDC uses content-defined chunking (Rabin fingerprint).
	CDC Type = "cdc"
)

// Config holds the configuration for the chunker factory.
type Config struct {
	Type           Type
	FixedSize      int
	MinSize        uint
	MaxSize        uint
}

// NewChunker is a factory function that returns a Chunker based on the configuration.
func NewChunker(cfg Config) (Chunker, error) {
	switch cfg.Type {
	case FixedSize:
		if cfg.FixedSize <= 0 {
			return nil, fmt.Errorf("invalid fixed chunk size: %d", cfg.FixedSize)
		}
		return NewFixedChunker(cfg.FixedSize), nil
	case CDC:
		return NewCDCChunker(cfg.MinSize, cfg.MaxSize), nil
	default:
		return nil, fmt.Errorf("unsupported chunker type: %s", cfg.Type)
	}
}

// Reassemble joins chunks back into a single reader, verifying their integrity against the provided hashes.
func Reassemble(chunks [][]byte, expectedHashes []ChunkHash) (io.Reader, error) {
	if len(chunks) != len(expectedHashes) {
		return nil, fmt.Errorf("mismatch between chunks (%d) and hashes (%d)", len(chunks), len(expectedHashes))
	}

	var buf bytes.Buffer
	for i, chunk := range chunks {
		// Verify integrity using BLAKE3
		hash := blake3.Sum256(chunk)
		actualHash := ChunkHash(fmt.Sprintf("%x", hash[:]))

		if actualHash != expectedHashes[i] {
			return nil, fmt.Errorf("integrity check failed for chunk %d: expected %s, got %s", i, expectedHashes[i], actualHash)
		}

		if _, err := buf.Write(chunk); err != nil {
			return nil, fmt.Errorf("failed to write chunk %d to buffer: %w", i, err)
		}
	}

	return &buf, nil
}
