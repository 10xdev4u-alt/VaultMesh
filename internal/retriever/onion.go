package retriever

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

// OnionRetriever handles multi-hop anonymous retrieval requests.
type OnionRetriever struct {
	host host.Host
}

// NewOnionRetriever creates a new OnionRetriever.
func NewOnionRetriever(h host.Host) *OnionRetriever {
	return &OnionRetriever{host: h}
}

// RetrieveViaProxy forwards a retrieval request through an intermediate peer.
func (o *OnionRetriever) RetrieveViaProxy(ctx context.Context, proxy peer.ID, target peer.AddrInfo, hash string) ([]byte, error) {
	// Connect to the proxy peer
	if err := o.host.Connect(ctx, peer.AddrInfo{ID: proxy}); err != nil {
		return nil, fmt.Errorf("failed to connect to onion proxy: %w", err)
	}

	// Open a stream to the proxy
	s, err := o.host.NewStream(ctx, proxy, "/vaultmesh/onion/1.0.0")
	if err != nil {
		return nil, fmt.Errorf("failed to open onion stream: %w", err)
	}
	defer s.Close()

	// Wrap request (Simplified for this commit)
	request := fmt.Sprintf("%s|%s", target.ID.String(), hash)
	if _, err := s.Write([]byte(request)); err != nil {
		return nil, err
	}

	return nil, nil
}
