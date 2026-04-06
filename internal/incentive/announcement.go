package incentive

import (
	"encoding/json"
	"time"
)

// StorageAnnouncement represents a node's offer to store data.
type StorageAnnouncement struct {
	PeerID    string    `json:"peer_id"`
	Capacity  int64     `json:"capacity"` // In bytes
	RepScore  float64   `json:"rep_score"`
	Timestamp time.Time `json:"timestamp"`
}

// Marshal serializes the announcement for GossipSub broadcasting.
func (a *StorageAnnouncement) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

// UnmarshalAnnouncement deserializes an incoming announcement.
func UnmarshalAnnouncement(data []byte) (*StorageAnnouncement, error) {
	var a StorageAnnouncement
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}
	return &a, nil
}
