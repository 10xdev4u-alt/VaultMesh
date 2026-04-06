package chunker

import (
	"encoding/hex"
	"sync"

	"github.com/zeebo/blake3"
)

// ChunkHash is a hex-encoded string of the BLAKE3 hash.
type ChunkHash string

// DedupEngine manages local deduplication of chunks.
type DedupEngine struct {
	mu     sync.RWMutex
	hashes map[ChunkHash]struct{}
}

// NewDedupEngine creates a new deduplication engine.
func NewDedupEngine() *DedupEngine {
	return &DedupEngine{
		hashes: make(map[ChunkHash]struct{}),
	}
}

// Hash returns the BLAKE3 hash of the given data.
func (e *DedupEngine) Hash(data []byte) ChunkHash {
	hash := blake3.Sum256(data)
	return ChunkHash(hex.EncodeToString(hash[:]))
}

// IsDuplicate checks if the chunk has already been processed by this engine.
func (e *DedupEngine) IsDuplicate(data []byte) (ChunkHash, bool) {
	hash := e.Hash(data)
	
	e.mu.RLock()
	defer e.mu.RUnlock()
	
	_, exists := e.hashes[hash]
	return hash, exists
}

// Register adds a chunk's hash to the engine.
func (e *DedupEngine) Register(hash ChunkHash) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.hashes[hash] = struct{}{}
}

// Process takes a list of chunks and returns only the unique ones along with their hashes.
func (e *DedupEngine) Process(chunks [][]byte) ([][]byte, []ChunkHash) {
	var uniqueChunks [][]byte
	var hashes []ChunkHash

	for _, chunk := range chunks {
		hash, isDup := e.IsDuplicate(chunk)
		hashes = append(hashes, hash)
		
		if !isDup {
			uniqueChunks = append(uniqueChunks, chunk)
			e.Register(hash)
		}
	}

	return uniqueChunks, hashes
}
