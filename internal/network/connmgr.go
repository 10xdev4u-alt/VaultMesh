package network

import (
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
)

// ConnMgrOptions returns libp2p options for managing peer connections.
func ConnMgrOptions(lowWater, highWater int, gracePeriod time.Duration) []libp2p.Option {
	mgr, err := connmgr.NewConnManager(
		lowWater,    // Low water mark
		highWater,   // High water mark
		connmgr.WithGracePeriod(gracePeriod),
	)
	if err != nil {
		return nil
	}
	return []libp2p.Option{
		libp2p.ConnectionManager(mgr),
	}
}
