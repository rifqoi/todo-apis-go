package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Task        string    `json:"task"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}
