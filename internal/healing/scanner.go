package healing

import (
	"context"

	"github.com/10xdev4u-alt/VaultMesh/internal/storage"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

// AvailabilityScanner audits the availability of data shards in the network.
type AvailabilityScanner struct {
	kdht *dht.IpfsDHT
}

// NewAvailabilityScanner creates a new AvailabilityScanner.
func NewAvailabilityScanner(kdht *dht.IpfsDHT) *AvailabilityScanner {
	return &AvailabilityScanner{kdht: kdht}
}

// AuditManifest checks if all shards in a manifest are currently provided by at least one peer.
func (s *AvailabilityScanner) AuditManifest(ctx context.Context, m *storage.Manifest) ([]string, error) {
	var missingShards []string

	for _, hash := range m.ChunkHashes {
		providers, err := s.kdht.FindProviders(ctx, hash)
		if err != nil || len(providers) == 0 {
			missingShards = append(missingShards, hash)
		}
	}

	return missingShards, nil
}
