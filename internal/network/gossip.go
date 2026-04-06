package network

import (
	"context"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
)

// GossipManager handles the pub/sub operations using libp2p GossipSub.
type GossipManager struct {
	Host host.Host
	PS   *pubsub.PubSub
}

// NewGossipManager initializes a new GossipSub instance on the given host.
func NewGossipManager(ctx context.Context, h host.Host) (*GossipManager, error) {
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gossipsub: %w", err)
	}

	return &GossipManager{
		Host: h,
		PS:   ps,
	}, nil
}

// JoinTopic joins a specific pub/sub topic.
func (m *GossipManager) JoinTopic(topicName string) (*pubsub.Topic, error) {
	topic, err := m.PS.Join(topicName)
	if err != nil {
		return nil, fmt.Errorf("failed to join topic %s: %w", topicName, err)
	}
	return topic, nil
}
