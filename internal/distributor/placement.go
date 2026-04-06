package distributor

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/10xdev4u-alt/VaultMesh/internal/network"
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
	h      host.Host
	scorer *network.PeerScoreManager
}

// NewPlacementStrategy creates a new PlacementStrategy with scoring support.
func NewPlacementStrategy(h host.Host, scorer *network.PeerScoreManager) *PlacementStrategy {
	return &PlacementStrategy{
		h:      h,
		scorer: scorer,
	}
}

// SelectBestPeers picks the top 'n' peers based on lowest latency.
func (s *PlacementStrategy) SelectBestPeers(ctx context.Context, n int) ([]peer.ID, error) {
	return s.SelectSmartPeers(ctx, n)
}

// SelectSmartPeers picks peers based on a combination of latency and reputation score.
func (s *PlacementStrategy) SelectSmartPeers(ctx context.Context, n int) ([]peer.ID, error) {
	peers := s.h.Network().Peers()
	if len(peers) == 0 {
		return nil, fmt.Errorf("no connected peers available")
	}

	type candidate struct {
		id    peer.ID
		score float64
	}

	candidates := make([]candidate, 0, len(peers))
	for _, p := range peers {
		if s.scorer != nil && s.scorer.IsBlacklisted(p) {
			continue // Skip blacklisted peers
		}

		lat := s.h.Peerstore().LatencyByPeer(p)
		
		// Heuristic score: lower is better. 
		// We use latency in ms as a baseline.
		latMs := float64(lat.Milliseconds())
		if latMs == 0 {
			latMs = 500 // Default for unknown latency
		}

		hScore := latMs
		
		candidates = append(candidates, candidate{id: p, score: hScore})
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score < candidates[j].score
	})

	count := n
	if len(candidates) < n {
		count = len(candidates)
	}

	selected := make([]peer.ID, 0, count)
	for i := 0; i < count; i++ {
		selected = append(selected, candidates[i].id)
	}

	return selected, nil
}
