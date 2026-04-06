package network

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
)

// HostConfig holds the configuration for initializing a libp2p host.
type HostConfig struct {
	ListenAddrs []string
	PrivKey     crypto.PrivKey
}

// NewHost creates and initializes a libp2p host with the provided configuration and NAT traversal enabled.
func NewHost(ctx context.Context, cfg HostConfig) (host.Host, error) {
	if len(cfg.ListenAddrs) == 0 {
		cfg.ListenAddrs = []string{
			"/ip4/0.0.0.0/tcp/0",
			"/ip4/0.0.0.0/udp/0/quic-v1",
		}
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(cfg.ListenAddrs...),
		libp2p.ChainOptions(
			libp2p.Transport(tcp.NewTCPTransport),
			libp2p.Transport(quic.NewTransport),
		),
	}

	// Add NAT traversal and Relay options
	opts = append(opts, NATOptions()...)

	if cfg.PrivKey != nil {
		opts = append(opts, libp2p.Identity(cfg.PrivKey))
	}

	h, err := libp2p.New(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize libp2p host: %w", err)
	}

	return h, nil
}
