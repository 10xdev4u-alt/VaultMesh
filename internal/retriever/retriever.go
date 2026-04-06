package retriever

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/10xdev4u-alt/VaultMesh/internal/distributor"
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
	providers, err := r.kdht.FindProviders(ctx, shardHash)
	if err != nil {
		return nil, fmt.Errorf("failed to find providers for shard %s: %w", shardHash, err)
	}

	if len(providers) == 0 {
		return nil, fmt.Errorf("no providers found for shard %s", shardHash)
	}

	return r.downloadFromPeer(ctx, providers[0], shardHash)
}

// RetrieveShardsParallel downloads multiple shards in parallel.
func (r *Retriever) RetrieveShardsParallel(ctx context.Context, shardHashes []string) ([][]byte, error) {
	var wg sync.WaitGroup
	shards := make([][]byte, len(shardHashes))
	errs := make(chan error, len(shardHashes))

	for i, hash := range shardHashes {
		wg.Add(1)
		go func(index int, h string) {
			defer wg.Done()
			data, err := r.RetrieveShard(ctx, h)
			if err != nil {
				errs <- fmt.Errorf("failed to retrieve shard %d (%s): %w", index, h, err)
				return
			}
			shards[index] = data
		}(i, hash)
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return shards, nil
}

// ReassembleShards uses the erasure coder to reconstruct the original chunk data.
func (r *Retriever) ReassembleShards(ctx context.Context, shards [][]byte, dataCount, parityCount int, originalSize int) ([]byte, error) {
	coder, err := distributor.NewErasureCoder(dataCount, parityCount)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize erasure coder for reassembly: %w", err)
	}

	data, err := coder.Reconstruct(shards, originalSize)
	if err != nil {
		return nil, fmt.Errorf("failed to reassemble shards: %w", err)
	}

	return data, nil
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

	if _, err := s.Write([]byte(hash)); err != nil {
		return nil, err
	}

	return io.ReadAll(s)
}
