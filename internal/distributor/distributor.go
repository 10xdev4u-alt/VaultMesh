package distributor

import (
	"fmt"

	"github.com/10xdev4u-alt/VaultMesh/internal/config"
)

// Distributor coordinates the distribution of data shards across the network.
type Distributor struct {
	cfg   *config.Config
	coder *ErasureCoder
}

// NewDistributor creates a new Distributor using the provided configuration.
func NewDistributor(cfg *config.Config) (*Distributor, error) {
	coder, err := NewErasureCoder(cfg.Redundancy.DataShards, cfg.Redundancy.ParityShards)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize erasure coder: %w", err)
	}

	return &Distributor{
		cfg:   cfg,
		coder: coder,
	}, nil
}
