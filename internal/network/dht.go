package network

import (
	"context"
	"fmt"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
)

// DHTManager handles the initialization and lifecycle of the Kademlia DHT.
type DHTManager struct {
	Host host.Host
	DHT  *dht.IpfsDHT
}

// NewDHTManager creates a new DHTManager and initializes the Kademlia DHT.
func NewDHTManager(ctx context.Context, h host.Host) (*DHTManager, error) {
	// Initialize DHT in server mode so this node can participate in the routing table
	kdht, err := dht.New(ctx, h, dht.Mode(dht.ModeAutoServer))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize kademlia dht: %w", err)
	}

	if err := kdht.Bootstrap(ctx); err != nil {
		return nil, fmt.Errorf("failed to bootstrap dht: %w", err)
	}

	return &DHTManager{
		Host: h,
		DHT:  kdht,
	}, nil
}

// Close closes the DHT and releases resources.
func (m *DHTManager) Close() error {
	return m.DHT.Close()
}
