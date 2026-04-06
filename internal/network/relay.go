package network

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
)

// RelayManager manages the node's ability to act as a circuit relay.
type RelayManager struct {
	Host host.Host
}

// SetupRelay initializes the relay service on the host.
func (m *RelayManager) SetupRelay() error {
	_, err := relay.New(m.Host)
	if err != nil {
		return fmt.Errorf("failed to start relay service: %w", err)
	}
	return nil
}

// RelayOptions returns options for enabling relay client functionality.
func RelayOptions() []libp2p.Option {
	return []libp2p.Option{
		libp2p.EnableRelay(), // Enable the ability to use relays
	}
}
