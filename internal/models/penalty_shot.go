package models

import (
	"time"
)

// PenaltyShot represents a penalty shot attempt
type PenaltyShot struct {
	ID           int       `json:"id" db:"id"`
	ContestantID int       `json:"contestant_id" db:"contestant_id"`
	ShotResult   string    `json:"shot_result" db:"shot_result"` // "goal" or "miss"
	AttemptNumber int      `json:"attempt_number" db:"attempt_number"`
	Timestamp    time.Time `json:"timestamp" db:"timestamp"`
	OperatorID   string    `json:"operator_id" db:"operator_id"`
	SessionID    string    `json:"session_id" db:"session_id"`
}

// ShotRequest represents the data structure for recording a penalty shot
type ShotRequest struct {
	ContestantID int    `json:"contestant_id" binding:"required"`
	ShotResult   string `json:"shot_result" binding:"required,oneof=goal miss"`
	AttemptNumber int   `json:"attempt_number" binding:"required,min=1"`
	OperatorID   string `json:"operator_id"`
	SessionID    string `json:"session_id"`
}

// ShotResponse represents the data structure for API responses
type ShotResponse struct {
	ID           int    `json:"id"`
	ContestantID int    `json:"contestant_id"`
	ShotResult   string `json:"shot_result"`
	AttemptNumber int   `json:"attempt_number"`
	Timestamp    string `json:"timestamp"`
	OperatorID   string `json:"operator_id"`
	SessionID    string `json:"session_id"`
}
