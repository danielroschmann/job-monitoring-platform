package jobs

import (
	"time"
)

type Job struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Source      string    `json:"source"`
	CreatedAt   time.Time `json:"created_at"`
}
