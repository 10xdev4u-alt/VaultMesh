package storage

import (
	"context"
	"encoding/json"
	"fmt"
)

// ChunkMetadata holds indexing information for a specific chunk.
type ChunkMetadata struct {
	Hash       string `json:"hash"`
	Size       int64  `json:"size"`
	Version    int    `json:"version"`
	PeerCount  int    `json:"peer_count"`
}

// ChunkStore manages higher-level chunk operations and metadata indexing.
type ChunkStore struct {
	store Store
}

// NewChunkStore creates a new ChunkStore wrapping a base Store.
func NewChunkStore(store Store) *ChunkStore {
	return &ChunkStore{store: store}
}

// PutChunk stores a chunk's data and its metadata.
func (cs *ChunkStore) PutChunk(ctx context.Context, hash string, data []byte) error {
	// Store the actual chunk data
	chunkKey := fmt.Sprintf("chunk:%s", hash)
	if err := cs.store.Put(ctx, []byte(chunkKey), data); err != nil {
		return fmt.Errorf("failed to store chunk data: %w", err)
	}

	// Initialize and store metadata
	meta := ChunkMetadata{
		Hash: hash,
		Size: int64(len(data)),
		Version: 1,
	}
	return cs.PutMetadata(ctx, hash, meta)
}

// GetChunk retrieves a chunk's data by its hash.
func (cs *ChunkStore) GetChunk(ctx context.Context, hash string) ([]byte, error) {
	chunkKey := fmt.Sprintf("chunk:%s", hash)
	return cs.store.Get(ctx, []byte(chunkKey))
}

// PutMetadata stores indexing metadata for a chunk.
func (cs *ChunkStore) PutMetadata(ctx context.Context, hash string, meta ChunkMetadata) error {
	metaKey := fmt.Sprintf("meta:%s", hash)
	data, err := json.Marshal(meta)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}
	return cs.store.Put(ctx, []byte(metaKey), data)
}

// GetMetadata retrieves indexing metadata for a chunk.
func (cs *ChunkStore) GetMetadata(ctx context.Context, hash string) (ChunkMetadata, error) {
	var meta ChunkMetadata
	metaKey := fmt.Sprintf("meta:%s", hash)
	data, err := cs.store.Get(ctx, []byte(metaKey))
	if err != nil {
		return meta, err
	}
	if err := json.Unmarshal(data, &meta); err != nil {
		return meta, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}
	return meta, nil
}

// HasChunk checks if a chunk exists in the store.
func (cs *ChunkStore) HasChunk(ctx context.Context, hash string) (bool, error) {
	chunkKey := fmt.Sprintf("chunk:%s", hash)
	return cs.store.Has(ctx, []byte(chunkKey))
}
