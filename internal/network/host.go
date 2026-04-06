package network

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	libp2pwebrtc "github.com/libp2p/go-libp2p/p2p/transport/webrtc"
	"github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
)

// HostConfig holds the configuration for initializing a libp2p host.
type HostConfig struct {
	ListenAddrs []string
	PrivKey     crypto.PrivKey
	Bandwidth   *BandwidthManager // Added for bandwidth tracking
}

// NewHost creates and initializes a libp2p host with the provided configuration.
func NewHost(ctx context.Context, cfg HostConfig) (host.Host, error) {
	if len(cfg.ListenAddrs) == 0 {
		cfg.ListenAddrs = []string{
			"/ip4/0.0.0.0/tcp/0",
			"/ip4/0.0.0.0/udp/0/quic-v1",
			"/ip4/0.0.0.0/udp/0/webrtc-direct",
		}
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(cfg.ListenAddrs...),
		libp2p.ChainOptions(
			libp2p.Transport(tcp.NewTCPTransport),
			libp2p.Transport(quic.NewTransport),
			libp2p.Transport(libp2pwebrtc.New),
		),
	}

	// Add NAT and Relay support
	opts = append(opts, NATOptions()...)

	// Add Connection Manager
	opts = append(opts, ConnMgrOptions(20, 50, time.Minute)...)

	// Add Bandwidth Reporter if provided
	if cfg.Bandwidth != nil {
		opts = append(opts, cfg.Bandwidth.Options())
	}

	if cfg.PrivKey != nil {
		opts = append(opts, libp2p.Identity(cfg.PrivKey))
	}

	h, err := libp2p.New(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize libp2p host: %w", err)
	}

	return h, nil
}
