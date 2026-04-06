package distributor

import (
	"context"
	"fmt"
	"io"
	"sync"
	"sync/atomic"

	"github.com/10xdev4u-alt/VaultMesh/internal/config"
	"github.com/10xdev4u-alt/VaultMesh/internal/crypto"
	"github.com/10xdev4u-alt/VaultMesh/internal/network"
	"github.com/10xdev4u-alt/VaultMesh/internal/storage"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

// ProgressCallback is a function that receives updates on the upload progress.
type ProgressCallback func(uploaded, total int64)

// Distributor coordinates the distribution of data shards across the network.
type Distributor struct {
	cfg       *config.Config
	coder     *ErasureCoder
	host      host.Host
	placement *PlacementStrategy
	// semaphore limits the number of concurrent uploads (backpressure)
	sem chan struct{}
}

// NewDistributor creates a new Distributor with a concurrency limit.
func NewDistributor(cfg *config.Config, h host.Host) (*Distributor, error) {
	coder, err := NewErasureCoder(cfg.Redundancy.DataShards, cfg.Redundancy.ParityShards)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize erasure coder: %w", err)
	}

	// Default to 10 concurrent uploads for backpressure
	concurrency := 10

	return &Distributor{
		cfg:       cfg,
		coder:     coder,
		host:      h,
		placement: NewPlacementStrategy(h, nil),
		sem:       make(chan struct{}, concurrency),
	}, nil
}

// DistributeWithBackpressure uploads shards while respecting the concurrency limit.
func (d *Distributor) DistributeWithBackpressure(ctx context.Context, data []byte, cb ProgressCallback) error {
	shards, err := d.coder.Encode(data)
	if err != nil {
		return err
	}

	peers, err := d.placement.SelectSmartPeers(ctx, len(shards))
	if err != nil {
		return err
	}

	var uploadedShards int64
	totalShards := int64(len(shards))

	var wg sync.WaitGroup
	errs := make(chan error, len(shards))

	for i, shard := range shards {
		if i >= len(peers) {
			break
		}

		// Acquire semaphore (backpressure)
		select {
		case d.sem <- struct{}{}:
		case <-ctx.Done():
			return ctx.Err()
		}

		wg.Add(1)
		go func(p peer.ID, data []byte) {
			defer wg.Done()
			defer func() { <-d.sem }() // Release semaphore

			if err := d.uploadShard(ctx, p, data); err != nil {
				errs <- fmt.Errorf("failed to upload shard to %s: %w", p, err)
				return
			}
			
			atomic.AddInt64(&uploadedShards, 1)
			if cb != nil {
				cb(atomic.LoadInt64(&uploadedShards), totalShards)
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

// VerifyShard requests a hash from a peer to confirm they have the correct shard data.
func (d *Distributor) VerifyShard(ctx context.Context, p peer.ID, expectedHash string) (bool, error) {
	s, err := d.host.NewStream(ctx, p, network.ProtocolHealth)
	if err != nil {
		return false, fmt.Errorf("failed to open verification stream: %w", err)
	}
	defer s.Close()

	if _, err := s.Write([]byte(expectedHash)); err != nil {
		return false, err
	}

	buf := make([]byte, 64)
	n, err := s.Read(buf)
	if err != nil && err != io.EOF {
		return false, err
	}

	actualHash := string(buf[:n])
	return actualHash == expectedHash, nil
}

// PublishManifest saves the file manifest to the Kademlia DHT for global discovery.
func (d *Distributor) PublishManifest(ctx context.Context, kdht *dht.IpfsDHT, fileID string, m *storage.Manifest) error {
	data, err := m.Marshal()
	if err != nil {
		return err
	}

	// Use the fileID (hash of the name or CID) as the key in the DHT
	if err := kdht.PutValue(ctx, "/vaultmesh/manifests/"+fileID, data); err != nil {
		return fmt.Errorf("failed to publish manifest to dht: %w", err)
	}

	return nil
}

// PublishEncryptedManifest encrypts and saves the file manifest to the DHT.
func (d *Distributor) PublishEncryptedManifest(ctx context.Context, kdht *dht.IpfsDHT, fileID string, m *storage.Manifest, masterKey []byte) error {
	data, err := m.Marshal()
	if err != nil {
		return err
	}

	// Encrypt the manifest using the layered encryption pipeline
	pipeline := crypto.NewLayeredEncryption(masterKey)
	encryptedData, err := pipeline.Encrypt(data)
	if err != nil {
		return fmt.Errorf("failed to encrypt manifest: %w", err)
	}

	// Use the fileID as the key in the DHT
	if err := kdht.PutValue(ctx, "/vaultmesh/manifests/"+fileID, encryptedData); err != nil {
		return fmt.Errorf("failed to publish manifest to dht: %w", err)
	}

	return nil
}
