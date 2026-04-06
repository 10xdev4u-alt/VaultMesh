package network

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
)

// HostConfig holds the configuration for initializing a libp2p host.
type HostConfig struct {
	ListenAddrs []string
}

// NewHost creates and initializes a libp2p host with TCP and QUIC transports.
func NewHost(ctx context.Context, cfg HostConfig) (host.Host, error) {
	// If no addresses are provided, listen on all interfaces with random ports
	if len(cfg.ListenAddrs) == 0 {
		cfg.ListenAddrs = []string{
			"/ip4/0.0.0.0/tcp/0",
			"/ip4/0.0.0.0/udp/0/quic-v1",
		}
	}

	h, err := libp2p.New(
		libp2p.ListenAddrStrings(cfg.ListenAddrs...),
		libp2p.ChainOptions(
			libp2p.Transport(tcp.NewTCPTransport),
			libp2p.Transport(quic.NewTransport),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize libp2p host: %w", err)
	}

	return h, nil
}
