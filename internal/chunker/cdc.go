package chunker

import (
	"fmt"
	"io"

	"github.com/restic/chunker"
)

// CDCChunker implements content-defined chunking using Rabin fingerprints.
type CDCChunker struct {
	minSize uint
	maxSize uint
}

// NewCDCChunker creates a new CDCChunker with min and max chunk sizes.
func NewCDCChunker(min, max uint) *CDCChunker {
	if min == 0 {
		min = 512 * 1024 // 512KB default min
	}
	if max == 0 {
		max = 8 * 1024 * 1024 // 8MB default max
	}
	return &CDCChunker{
		minSize: min,
		maxSize: max,
	}
}

// Split divides the input reader into content-defined chunks.
func (c *CDCChunker) Split(r io.Reader) ([][]byte, error) {
	var chunks [][]byte
	
	// We use a polynomial for the Rabin fingerprint. 
	// This one is a standard choice.
	pol := chunker.Pol(0x3DA3358B4DC173)
	
	ch := chunker.New(r, pol)
	// Note: restic/chunker defaults are used for min/max if not explicitly handled here
	// The restic/chunker library handles the chunking logic internally.
	
	for {
		chunk, err := ch.Next(nil)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cdc split failed: %w", err)
		}

		// Copy chunk data
		data := make([]byte, len(chunk.Data))
		copy(data, chunk.Data)
		chunks = append(chunks, data)
	}

	return chunks, nil
}
