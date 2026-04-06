package healing

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
)

// NodeScore tracks the reliability metrics for a specific peer.
type NodeScore struct {
	LastSeen    time.Time
	MissedPings int
	LatencyAvg  time.Duration
	FailureProb float64
}

// Scorer manages failure probability scores for all known nodes.
type Scorer struct {
	mu     sync.RWMutex
	scores map[peer.ID]*NodeScore
}

// NewScorer creates a new Scorer.
func NewScorer() *Scorer {
	return &Scorer{
		scores: make(map[peer.ID]*NodeScore),
	}
}

// UpdateScore updates a node's reliability metrics based on a heartbeat result.
func (s *Scorer) UpdateScore(p peer.ID, success bool, latency time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	score, exists := s.scores[p]
	if !exists {
		score = &NodeScore{}
		s.scores[p] = score
	}

	if success {
		score.LastSeen = time.Now()
		score.MissedPings = 0
		score.LatencyAvg = (score.LatencyAvg + latency) / 2
		score.FailureProb *= 0.9
	} else {
		score.MissedPings++
		score.FailureProb += 0.1 * float64(score.MissedPings)
	}
}

// ShouldEvacuate checks if a node is deemed too unreliable to hold shards.
func (s *Scorer) ShouldEvacuate(p peer.ID) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	score, exists := s.scores[p]
	if !exists {
		return false
	}

	return score.FailureProb > 0.7
}
