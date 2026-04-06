package storage

import "context"

// Store defines the interface for key-value storage of chunks and metadata.
type Store interface {
	// Put stores a value associated with a key.
	Put(ctx context.Context, key []byte, value []byte) error
	// Get retrieves a value associated with a key.
	Get(ctx context.Context, key []byte) ([]byte, error)
	// Delete removes a key and its associated value.
	Delete(ctx context.Context, key []byte) error
	// Has checks if a key exists in the store.
	Has(ctx context.Context, key []byte) (bool, error)
	// Close closes the store and releases resources.
	Close() error
}
