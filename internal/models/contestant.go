package models

import (
	"time"
)

// Contestant represents a registered contestant in the system
type Contestant struct {
	ID        int       `json:"id" db:"id"`
	BadgeID   string    `json:"badge_id" db:"badge_id"`
	Name      string    `json:"name" db:"name"`
	Company   string    `json:"company" db:"company"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ContestantRequest represents the data structure for creating a new contestant
type ContestantRequest struct {
	BadgeID string `json:"badge_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

// ContestantResponse represents the data structure for API responses
type ContestantResponse struct {
	ID        int    `json:"id"`
	BadgeID   string `json:"badge_id"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}
