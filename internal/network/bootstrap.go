package network

import (
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

// DefaultBootstrapPeers contains the standard libp2p bootstrap nodes.
var DefaultBootstrapPeers = []string{
	"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTBsPWCcqS2S2s7aPvwVfN2p7rQdEaJs",
	"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2Ecws3N79txbcocFQ977XLeqM6K1Y78T9fG6t4q8G",
	"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMo96F8tA6yHArD9Nn7yS85tshx5G7nQfG7xN9qD",
}

// Bootstrap connects to the given bootstrap peers in parallel.
func Bootstrap(ctx context.Context, h host.Host, bootstrapPeers []string) error {
	var wg sync.WaitGroup
	for _, p := range bootstrapPeers {
		ma, err := multiaddr.NewMultiaddr(p)
		if err != nil {
			fmt.Printf("Bootstrap: failed to parse multiaddr %s: %s\n", p, err)
			continue
		}

		pi, err := peer.AddrInfoFromP2pAddr(ma)
		if err != nil {
			fmt.Printf("Bootstrap: failed to get addr info from %s: %s\n", p, err)
			continue
		}

		wg.Add(1)
		go func(info peer.AddrInfo) {
			defer wg.Done()
			if err := h.Connect(ctx, info); err != nil {
				fmt.Printf("Bootstrap: failed to connect to %s: %s\n", info.ID, err)
			}
		}(*pi)
	}

	wg.Wait()
	return nil
}
