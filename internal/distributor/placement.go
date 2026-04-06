package distributor

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/10xdev4u-alt/VaultMesh/internal/incentive"
	"github.com/10xdev4u-alt/VaultMesh/internal/network"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

// PlacementStrategy handles the selection of peers for data shard distribution.
type PlacementStrategy struct {
	h          host.Host
	scorer     *network.PeerScoreManager
	repManager *incentive.ReputationManager
}

// NewPlacementStrategy creates a new PlacementStrategy with reputation support.
func NewPlacementStrategy(h host.Host, scorer *network.PeerScoreManager, rm *incentive.ReputationManager) *PlacementStrategy {
	return &PlacementStrategy{
		h:          h,
		scorer:     scorer,
		repManager: rm,
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
			continue
		}

		lat := s.h.Peerstore().LatencyByPeer(p)
		latMs := float64(lat.Milliseconds())
		if latMs == 0 {
			latMs = 500
		}

		// Apply reputation boost
		repScore := 0.5
		if s.repManager != nil {
			repScore = s.repManager.GetScore(p.String())
		}

		// Higher reputation = lower (better) score
		hScore := latMs / (1.0 + repScore)
		
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
