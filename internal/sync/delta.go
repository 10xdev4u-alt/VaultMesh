package sync

import (
	"github.com/10xdev4u-alt/VaultMesh/internal/chunker"
)

// DeltaSyncEngine calculates the difference between two sets of chunks.
type DeltaSyncEngine struct{}

// NewDeltaSyncEngine creates a new DeltaSyncEngine.
func NewDeltaSyncEngine() *DeltaSyncEngine {
	return &DeltaSyncEngine{}
}

// Diff calculates which chunks in 'newHashes' are not present in 'oldHashes'.
func (e *DeltaSyncEngine) Diff(oldHashes, newHashes []chunker.ChunkHash) []chunker.ChunkHash {
	oldMap := make(map[chunker.ChunkHash]struct{})
	for _, h := range oldHashes {
		oldMap[h] = struct{}{}
	}

	var missing []chunker.ChunkHash
	for _, h := range newHashes {
		if _, exists := oldMap[h]; !exists {
			missing = append(missing, h)
		}
	}

	return missing
}
