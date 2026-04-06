package sync

import (
	"fmt"
)

// ConflictDetector identifies divergences in the version history.
type ConflictDetector struct{}

// NewConflictDetector creates a new ConflictDetector.
func NewConflictDetector() *ConflictDetector {
	return &ConflictDetector{}
}

// Detect identifies if two version nodes represent a conflict.
func (d *ConflictDetector) Detect(v1, v2 *VersionNode) bool {
	return v1.ID != v2.ID && len(v1.ParentIDs) > 0 && len(v2.ParentIDs) > 0
}

// Resolve applies a basic resolution strategy.
func (d *ConflictDetector) Resolve(v1, v2 *VersionNode) *VersionNode {
	if v1.Timestamp.After(v2.Timestamp) {
		return v1
	}
	return v2
}
