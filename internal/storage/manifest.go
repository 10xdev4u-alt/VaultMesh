package storage

import (
	"encoding/json"
	"time"
)

// Manifest represents the metadata for a complete file stored in VaultMesh.
type Manifest struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	ChunkHashes  []string  `json:"chunk_hashes"`
	DataShards   int       `json:"data_shards"`
	ParityShards int       `json:"parity_shards"`
	CreatedAt    time.Time `json:"created_at"`
}

// Marshal serializes the manifest to JSON bytes.
func (m *Manifest) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalManifest deserializes JSON bytes into a Manifest.
func UnmarshalManifest(data []byte) (*Manifest, error) {
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return &m, nil
}
