package healing

import (
	"context"
	"fmt"
	"time"

	"github.com/10xdev4u-alt/VaultMesh/internal/network"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

// HeartbeatManager handles the periodic heartbeat protocol.
type HeartbeatManager struct {
	host host.Host
}

// NewHeartbeatManager creates a new HeartbeatManager.
func NewHeartbeatManager(h host.Host) *HeartbeatManager {
	return &HeartbeatManager{host: h}
}

// SendHeartbeat sends a ping to a peer to verify its availability.
func (m *HeartbeatManager) SendHeartbeat(ctx context.Context, p peer.ID) (time.Duration, error) {
	s, err := m.host.NewStream(ctx, p, network.ProtocolHealth)
	if err != nil {
		return 0, fmt.Errorf("failed to open heartbeat stream: %w", err)
	}
	defer s.Close()

	start := time.Now()
	if _, err := s.Write([]byte("PING")); err != nil {
		return 0, err
	}

	buf := make([]byte, 4)
	if _, err := s.Read(buf); err != nil {
		return 0, err
	}

	return time.Since(start), nil
}
