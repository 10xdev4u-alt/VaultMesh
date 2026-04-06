package chunker

import (
	"fmt"
	"io"
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
