package incentive

import (
	"sync"
	"time"
)

// ReputationScore tracks the long-term trustworthiness of a peer.
type ReputationScore struct {
	TotalUptime           float64
	SuccessfulRetrievals  int
	FailedIntegrityChecks int
	Score                 float64
}

// ReputationManager manages the reputation database.
type ReputationManager struct {
	mu     sync.RWMutex
	scores map[string]*ReputationScore
}

// NewReputationManager creates a new ReputationManager.
func NewReputationManager() *ReputationManager {
	return &ReputationManager{
		scores: make(map[string]*ReputationScore),
	}
}

// GetScore returns the current score for a peer.
func (m *ReputationManager) GetScore(peerID string) float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if s, exists := m.scores[peerID]; exists {
		return s.Score
	}
	return 0.5
}

// RecordSuccess improves a peer's reputation.
func (m *ReputationManager) RecordSuccess(peerID string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, exists := m.scores[peerID]
	if !exists {
		s = &ReputationScore{Score: 0.5}
		m.scores[peerID] = s
	}
	s.SuccessfulRetrievals++
	s.Score += 0.01
	if s.Score > 1.0 {
		s.Score = 1.0
	}
}

// RecordSybilCheck ensures that a peer's reputation growth is gated by time.
func (m *ReputationManager) RecordSybilCheck(peerID string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, exists := m.scores[peerID]
	if !exists {
		return false
	}

	// New peers cannot exceed 0.6 score in their first 24 hours.
	if s.TotalUptime < float64(24*time.Hour) && s.Score > 0.6 {
		s.Score = 0.6
		return true
	}
	return false
}
