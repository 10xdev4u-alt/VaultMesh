package retriever

import (
	"context"
	"fmt"
	"io"

	"github.com/10xdev4u-alt/VaultMesh/internal/network"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

// Retriever handles finding and downloading data shards from the network.
type Retriever struct {
	host host.Host
	kdht *dht.IpfsDHT
}

// NewRetriever creates a new Retriever.
func NewRetriever(h host.Host, kdht *dht.IpfsDHT) *Retriever {
	return &Retriever{
		host: h,
		kdht: kdht,
	}
}

// RetrieveShard finds a shard's location via DHT and downloads it.
func (r *Retriever) RetrieveShard(ctx context.Context, shardHash string) ([]byte, error) {
	// Find providers for this shard hash in the DHT
	providers, err := r.kdht.FindProviders(ctx, shardHash)
	if err != nil {
		return nil, fmt.Errorf("failed to find providers for shard %s: %w", shardHash, err)
	}

	if len(providers) == 0 {
		return nil, fmt.Errorf("no providers found for shard %s", shardHash)
	}

	// For now, try the first provider
	return r.downloadFromPeer(ctx, providers[0], shardHash)
}

func (r *Retriever) downloadFromPeer(ctx context.Context, p peer.AddrInfo, hash string) ([]byte, error) {
	if err := r.host.Connect(ctx, p); err != nil {
		return nil, fmt.Errorf("failed to connect to provider %s: %w", p.ID, err)
	}

	s, err := r.host.NewStream(ctx, p.ID, network.ProtocolDownload)
	if err != nil {
		return nil, fmt.Errorf("failed to open download stream: %w", err)
	}
	defer s.Close()

	// Send the request (shard hash)
	if _, err := s.Write([]byte(hash)); err != nil {
		return nil, err
	}

	// Read the shard data
	return io.ReadAll(s)
}
