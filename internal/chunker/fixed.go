package chunker

import (
	"fmt"
	"io"
	"log/slog"
)

// FixedChunker implements fixed-size chunking.
type FixedChunker struct {
	chunkSize int
}

// NewFixedChunker creates a new FixedChunker with the given size.
func NewFixedChunker(size int) *FixedChunker {
	return &FixedChunker{chunkSize: size}
}

// Split divides the input reader into fixed-size chunks.
// It returns a slice of data chunks.
func (f *FixedChunker) Split(r io.Reader) ([][]byte, error) {
	if f.chunkSize <= 0 {
		return nil, fmt.Errorf("chunk size must be greater than 0: %w", ErrInvalidChunkSize)
	}

	var chunks [][]byte
	buf := make([]byte, f.chunkSize)

	for {
		n, err := io.ReadFull(r, buf)
		if n > 0 {
			chunk := make([]byte, n)
			copy(chunk, buf[:n])
			chunks = append(chunks, chunk)
		}

		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			slog.Error("failed to read from reader", "error", err)
			return nil, fmt.Errorf("failed to split chunks: %w", err)
		}
	}

	return chunks, nil
}

// ErrInvalidChunkSize is returned when the chunk size is invalid.
var ErrInvalidChunkSize = fmt.Errorf("invalid chunk size")
