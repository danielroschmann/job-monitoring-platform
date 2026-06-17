package jobs

import (
	"time"
)

type Job struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Company     string
	Location    string
	Description string
	URL         string
	Source      string
	CreatedAt   time.Time
}
