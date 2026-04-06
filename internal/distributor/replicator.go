package distributor

import (
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
)

// Replicator ensures that data shards meet a target replication factor.
type Replicator struct {
	distributor *Distributor
	target      int
}

// NewReplicator creates a new Replicator.
func NewReplicator(d *Distributor, targetFactor int) *Replicator {
	if targetFactor <= 0 {
		targetFactor = 3 // Default replication factor
	}
	return &Replicator{
		distributor: d,
		target:      targetFactor,
	}
}

// ReplicateShard ensures a single shard is replicated to multiple peers.
func (r *Replicator) ReplicateShard(ctx context.Context, shardData []byte) error {
	peers, err := r.distributor.placement.SelectBestPeers(ctx, r.target)
	if err != nil {
		return fmt.Errorf("failed to select peers for replication: %w", err)
	}

	var wg sync.WaitGroup
	errs := make(chan error, len(peers))

	for _, p := range peers {
		wg.Add(1)
		go func(pid peer.ID) {
			defer wg.Done()
			if err := r.distributor.uploadShard(ctx, pid, shardData); err != nil {
				errs <- err
			}
		}(p)
	}

	wg.Wait()
	close(errs)

	// In a real system, we might be satisfied if at least one upload succeeds,
	// but here we check for any errors.
	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
