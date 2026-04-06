package distributor

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

// PeerLatency represents a peer and its measured connection latency.
type PeerLatency struct {
	ID      peer.ID
	Latency time.Duration
}

// PlacementStrategy handles the selection of peers for data shard distribution.
type PlacementStrategy struct {
	h host.Host
}

// NewPlacementStrategy creates a new PlacementStrategy.
func NewPlacementStrategy(h host.Host) *PlacementStrategy {
	return &PlacementStrategy{h: h}
}

// SelectBestPeers picks the top 'n' peers based on lowest latency.
func (s *PlacementStrategy) SelectBestPeers(ctx context.Context, n int) ([]peer.ID, error) {
	peers := s.h.Network().Peers()
	if len(peers) == 0 {
		return nil, fmt.Errorf("no connected peers available")
	}

	latencies := make([]PeerLatency, 0, len(peers))
	for _, p := range peers {
		// Get latency from libp2p peerstore
		lat := s.h.Peerstore().LatencyByPeer(p)
		latencies = append(latencies, PeerLatency{ID: p, Latency: lat})
	}

	// Sort by latency (lowest first)
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i].Latency < latencies[j].Latency
	})

	// Select top n
	count := n
	if len(latencies) < n {
		count = len(latencies)
	}

	selected := make([]peer.ID, 0, count)
	for i := 0; i < count; i++ {
		selected = append(selected, latencies[i].ID)
	}

	return selected, nil
}
