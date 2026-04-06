package storage

import (
	"context"
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

// BadgerStore implements the Store interface using BadgerDB.
type BadgerStore struct {
	db *badger.DB
}

// NewBadgerStore creates a new BadgerStore at the specified path.
func NewBadgerStore(path string) (*BadgerStore, error) {
	opts := badger.DefaultOptions(path).WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to open badger db at %s: %w", path, err)
	}
	return &BadgerStore{db: db}, nil
}

// Put stores a value associated with a key in BadgerDB.
func (s *BadgerStore) Put(ctx context.Context, key []byte, value []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

// Get retrieves a value associated with a key from BadgerDB.
func (s *BadgerStore) Get(ctx context.Context, key []byte) ([]byte, error) {
	var val []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		return err
	})
	if err == badger.ErrKeyNotFound {
		return nil, fmt.Errorf("key not found: %x", key)
	}
	return val, err
}

// Delete removes a key and its associated value from BadgerDB.
func (s *BadgerStore) Delete(ctx context.Context, key []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

// Has checks if a key exists in BadgerDB.
func (s *BadgerStore) Has(ctx context.Context, key []byte) (bool, error) {
	exists := false
	err := s.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err == nil {
			exists = true
			return nil
		}
		if err == badger.ErrKeyNotFound {
			return nil
		}
		return err
	})
	return exists, err
}

// Close closes the BadgerDB instance.
func (s *BadgerStore) Close() error {
	return s.db.Close()
}
