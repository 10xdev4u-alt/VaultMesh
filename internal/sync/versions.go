package sync

import (
	"time"
)

// VersionNode represents a single state of a file in the Merkle DAG.
type VersionNode struct {
	ID         string    `json:"id"`
	ParentIDs  []string  `json:"parent_ids"`
	ManifestID string    `json:"manifest_id"`
	Timestamp  time.Time `json:"timestamp"`
	Author     string    `json:"author"`
}

// VersionHistory manages the history of file changes.
type VersionHistory struct {
	Versions map[string]*VersionNode
}

// NewVersionHistory creates a new VersionHistory.
func NewVersionHistory() *VersionHistory {
	return &VersionHistory{
		Versions: make(map[string]*VersionNode),
	}
}

// AddVersion records a new version in the history.
func (h *VersionHistory) AddVersion(node *VersionNode) {
	h.Versions[node.ID] = node
}
