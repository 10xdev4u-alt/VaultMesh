package network

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/metrics"
)

// BandwidthManager wraps the libp2p bandwidth reporter to provide easy access to stats.
type BandwidthManager struct {
	Reporter *metrics.BandwidthCounter
}

// NewBandwidthManager creates a new BandwidthManager.
func NewBandwidthManager() *BandwidthManager {
	return &BandwidthManager{
		Reporter: metrics.NewBandwidthCounter(),
	}
}

// Options returns the libp2p option to enable bandwidth reporting.
func (m *BandwidthManager) Options() libp2p.Option {
	return libp2p.BandwidthReporter(m.Reporter)
}

// GetStats returns the total bytes in and out.
func (m *BandwidthManager) GetStats() (int64, int64) {
	stats := m.Reporter.GetBandwidthTotals()
	return stats.TotalIn, stats.TotalOut
}
