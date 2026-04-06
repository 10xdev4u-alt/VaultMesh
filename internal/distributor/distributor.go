package distributor

import (
	"context"
	"fmt"
	"sync"

	"github.com/10xdev4u-alt/VaultMesh/internal/config"
	"github.com/10xdev4u-alt/VaultMesh/internal/network"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

// Distributor coordinates the distribution of data shards across the network.
type Distributor struct {
	cfg       *config.Config
	coder     *ErasureCoder
	host      host.Host
	placement *PlacementStrategy
}

// NewDistributor creates a new Distributor.
func NewDistributor(cfg *config.Config, h host.Host) (*Distributor, error) {
	coder, err := NewErasureCoder(cfg.Redundancy.DataShards, cfg.Redundancy.ParityShards)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize erasure coder: %w", err)
	}

	return &Distributor{
		cfg:       cfg,
		coder:     coder,
		host:      h,
		placement: NewPlacementStrategy(h),
	}, nil
}

// DistributeParallel splits the data and uploads shards in parallel to selected peers.
func (d *Distributor) DistributeParallel(ctx context.Context, data []byte) error {
	shards, err := d.coder.Encode(data)
	if err != nil {
		return err
	}

	peers, err := d.placement.SelectBestPeers(ctx, len(shards))
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errs := make(chan error, len(shards))

	for i, shard := range shards {
		if i >= len(peers) {
			break
		}

		wg.Add(1)
		go func(p peer.ID, data []byte) {
			defer wg.Done()
			if err := d.uploadShard(ctx, p, data); err != nil {
				errs <- fmt.Errorf("failed to upload shard to %s: %w", p, err)
			}
		}(peers[i], shard)
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

// uploadShard is a private helper to send a shard to a specific peer using the custom protocol.
func (d *Distributor) uploadShard(ctx context.Context, p peer.ID, data []byte) error {
	s, err := d.host.NewStream(ctx, p, network.ProtocolUpload)
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = s.Write(data)
	return err
}
