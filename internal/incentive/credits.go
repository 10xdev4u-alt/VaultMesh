package incentive

import (
	"sync"
)

// ResourceCredits tracks the balance of a user's resource contributions.
type ResourceCredits struct {
	BandwidthBalance int64 // In bytes
	StorageBalance   int64 // In bytes/hour
}

// CreditManager manages the accounting of resource credits.
type CreditManager struct {
	mu       sync.RWMutex
	balances map[string]*ResourceCredits
}

// NewCreditManager creates a new CreditManager.
func NewCreditManager() *CreditManager {
	return &CreditManager{
		balances: make(map[string]*ResourceCredits),
	}
}

// GetBalance returns the current credit balance for a user.
func (m *CreditManager) GetBalance(userID string) ResourceCredits {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if b, exists := m.balances[userID]; exists {
		return *b
	}
	return ResourceCredits{}
}

// AwardCredits increases a user's balance.
func (m *CreditManager) AwardCredits(userID string, bandwidth int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.balances[userID]; !exists {
		m.balances[userID] = &ResourceCredits{}
	}
	m.balances[userID].BandwidthBalance += bandwidth
}

// SpendCredits decreases a user's balance.
func (m *CreditManager) SpendCredits(userID string, bandwidth int64) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	b, exists := m.balances[userID]
	if !exists || b.BandwidthBalance < bandwidth {
		return false
	}
	b.BandwidthBalance -= bandwidth
	return true
}
