package users

import "job-monitoring-platform/api/internal/jobs"

type User struct {
	ID        uint       `gorm:"primaryKey"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Jobs      []jobs.Job `gorm:"many2many:user_jobs"`
}
