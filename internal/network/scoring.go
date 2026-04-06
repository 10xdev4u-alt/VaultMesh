package network

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/peer"
)

// PeerScoreManager tracks reputation and connection status for network peers.
type PeerScoreManager struct {
	mu        sync.RWMutex
	scores    map[peer.ID]int
	blacklist map[peer.ID]bool
}

// NewPeerScoreManager creates a new PeerScoreManager.
func NewPeerScoreManager() *PeerScoreManager {
	return &PeerScoreManager{
		scores:    make(map[peer.ID]int),
		blacklist: make(map[peer.ID]bool),
	}
}

// IncrementScore increases the score of a peer (e.g., successful upload).
func (m *PeerScoreManager) IncrementScore(p peer.ID, delta int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.scores[p] += delta
}

// DecrementScore decreases the score of a peer (e.g., failed integrity check).
func (m *PeerScoreManager) DecrementScore(p peer.ID, delta int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.scores[p] -= delta

	// Automatically blacklist if score falls too low (e.g., -100)
	if m.scores[p] < -100 {
		m.blacklist[p] = true
	}
}

// IsBlacklisted checks if a peer is currently blacklisted.
func (m *PeerScoreManager) IsBlacklisted(p peer.ID) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.blacklist[p]
}

// Blacklist adds a peer to the blacklist explicitly.
func (m *PeerScoreManager) Blacklist(p peer.ID) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.blacklist[p] = true
}
