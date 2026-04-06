package network

import (
	"context"
	"encoding/json"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

const PEXTopicName = "vaultmesh-pex"

// PEXManager handles peer exchange messages via GossipSub.
type PEXManager struct {
	Host  host.Host
	Topic *pubsub.Topic
	Sub   *pubsub.Subscription
}

// NewPEXManager creates and initializes a PEXManager.
func NewPEXManager(ctx context.Context, h host.Host, ps *pubsub.PubSub) (*PEXManager, error) {
	topic, err := ps.Join(PEXTopicName)
	if err != nil {
		return nil, fmt.Errorf("failed to join pex topic: %w", err)
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to pex topic: %w", err)
	}

	return &PEXManager{
		Host:  h,
		Topic: topic,
		Sub:   sub,
	}, nil
}

// BroadcastPeers sends the list of currently connected peers to the PEX topic.
func (m *PEXManager) BroadcastPeers(ctx context.Context) error {
	peers := m.Host.Network().Peers()
	peerAddrs := make([]peer.AddrInfo, 0, len(peers))

	for _, p := range peers {
		peerAddrs = append(peerAddrs, m.Host.Peerstore().PeerInfo(p))
	}

	data, err := json.Marshal(peerAddrs)
	if err != nil {
		return fmt.Errorf("failed to marshal pex data: %w", err)
	}

	return m.Topic.Publish(ctx, data)
}

// ListenForPeers listens for incoming PEX messages and connects to discovered peers.
func (m *PEXManager) ListenForPeers(ctx context.Context) {
	for {
		msg, err := m.Sub.Next(ctx)
		if err != nil {
			return
		}

		if msg.ReceivedFrom == m.Host.ID() {
			continue
		}

		var discoveredPeers []peer.AddrInfo
		if err := json.Unmarshal(msg.Data, &discoveredPeers); err != nil {
			continue
		}

		for _, p := range discoveredPeers {
			if p.ID != m.Host.ID() {
				m.Host.Connect(ctx, p)
			}
		}
	}
}
