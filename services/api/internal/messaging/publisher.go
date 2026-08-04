package messaging

import "time"

type CollectJobMessage struct {
	Source      string    `json:"source"`
	RequestedBy uint      `json:"requested_by"`
	RequestedAt time.Time `json:"requested_at"`
}
