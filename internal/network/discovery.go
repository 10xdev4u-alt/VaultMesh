package network

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// discoveryNotifer is a private struct that handles peer discovery events.
type discoveryNotifer struct {
	h host.Host
}

// HandlePeerFound is called by the mDNS service when a new peer is discovered on the local network.
func (n *discoveryNotifer) HandlePeerFound(pi peer.AddrInfo) {
	// Don't connect to ourselves
	if pi.ID == n.h.ID() {
		return
	}

	// Attempt to connect to the discovered peer
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("mDNS: failed to connect to discovered peer %s: %s\n", pi.ID, err)
	}
}

// SetupMDNS initializes the mDNS service for the given host.
func SetupMDNS(h host.Host) error {
	s := mdns.NewMdnsService(h, "vaultmesh-discovery", &discoveryNotifer{h: h})
	return s.Start()
}
