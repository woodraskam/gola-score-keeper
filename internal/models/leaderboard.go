package models

// LeaderboardEntry represents a contestant's position on the leaderboard
type LeaderboardEntry struct {
	ContestantID      int     `json:"contestant_id" db:"contestant_id"`
	Name              string  `json:"name" db:"name"`
	Company           string  `json:"company" db:"company"`
	TotalAttempts     int     `json:"total_attempts" db:"total_attempts"`
	SuccessfulShots   int     `json:"successful_shots" db:"successful_shots"`
	SuccessPercentage float64 `json:"success_percentage" db:"success_percentage"`
	LastUpdated       string  `json:"last_updated" db:"last_updated"`
}

// LeaderboardRequest represents the data structure for leaderboard queries
type LeaderboardRequest struct {
	Filter string `json:"filter" form:"filter"` // "all", "today", "top10"
	Limit  int    `json:"limit" form:"limit"`
}

// LeaderboardResponse represents the data structure for leaderboard API responses
type LeaderboardResponse struct {
	Entries []LeaderboardEntry `json:"entries"`
	Total   int                `json:"total"`
	Filter  string             `json:"filter"`
}
