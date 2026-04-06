package network

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
)

// NATOptions returns a slice of libp2p options specifically for NAT traversal and relay support.
func NATOptions() []libp2p.Option {
	return []libp2p.Option{
		libp2p.NATPortMap(),
		libp2p.EnableAutoRelayWithStaticRelays(client.ProtoID),
		libp2p.EnableHolePunching(),
		libp2p.EnableRelay(), // Added for broad relay support
	}
}
