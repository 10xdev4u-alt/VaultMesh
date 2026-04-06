package network

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
)

// NATOptions returns a slice of libp2p options specifically for NAT traversal and relay support.
func NATOptions() []libp2p.Option {
	return []libp2p.Option{
		// Enable NAT port mapping via UPnP/NAT-PMP
		libp2p.NATPortMap(),
		// Enable the node to find and use static relays
		libp2p.EnableAutoRelayWithStaticRelays(client.ProtoID),
		// Enable direct connection attempts between nodes behind NAT (DCUtR)
		libp2p.EnableHolePunching(),
	}
}
