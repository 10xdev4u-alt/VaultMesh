package network

import (
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

// PeerConnectionInfo represents a summary of a peer's connection status.
type PeerConnectionInfo struct {
	ID        peer.ID
	Addresses []multiaddr.Multiaddr
	Protocols []string
}

// TopologyMapper provides methods to query the current network structure.
type TopologyMapper struct {
	Host host.Host
}

// NewTopologyMapper creates a new TopologyMapper.
func NewTopologyMapper(h host.Host) *TopologyMapper {
	return &TopologyMapper{Host: h}
}

// GetConnectedPeers returns a list of all currently connected peers and their addresses.
func (m *TopologyMapper) GetConnectedPeers() []PeerConnectionInfo {
	peers := m.Host.Network().Peers()
	info := make([]PeerConnectionInfo, 0, len(peers))

	for _, p := range peers {
		conns := m.Host.Network().ConnsToPeer(p)
		addrs := make([]multiaddr.Multiaddr, 0, len(conns))
		for _, c := range conns {
			addrs = append(addrs, c.RemoteMultiaddr())
		}

		protocols, _ := m.Host.Peerstore().GetProtocols(p)

		info = append(info, PeerConnectionInfo{
			ID:        p,
			Addresses: addrs,
			Protocols: protocols,
		})
	}

	return info
}
