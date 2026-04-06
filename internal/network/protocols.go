package network

import (
	"github.com/libp2p/go-libp2p/core/protocol"
)

const (
	// ProtocolUpload is used for uploading data chunks to peers.
	ProtocolUpload protocol.ID = "/vaultmesh/upload/1.0.0"
	// ProtocolDownload is used for retrieving data chunks from peers.
	ProtocolDownload protocol.ID = "/vaultmesh/download/1.0.0"
	// ProtocolHealth is used for node health checks and status heartbeats.
	ProtocolHealth protocol.ID = "/vaultmesh/health/1.0.0"
)

// SupportedProtocols returns a list of all protocols supported by VaultMesh.
func SupportedProtocols() []protocol.ID {
	return []protocol.ID{
		ProtocolUpload,
		ProtocolDownload,
		ProtocolHealth,
	}
}
