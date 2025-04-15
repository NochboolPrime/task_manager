package models

import "time"

type Project struct {
	ID          int       `json:"id"`
	TeamID      int       `json:"team_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
