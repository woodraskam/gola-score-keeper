package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Setting up Gola Score Keeper database...")

	// Get database path from environment or use default
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/score_keeper.db"
	}

	// Create data directory if it doesn't exist
	if err := os.MkdirAll("./data", 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Enable WAL mode for better concurrency
	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		log.Fatalf("Failed to enable WAL mode: %v", err)
	}

	// Create tables
	if err := createTables(db); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Create indexes
	if err := createIndexes(db); err != nil {
		log.Fatalf("Failed to create indexes: %v", err)
	}

	log.Println("Database setup completed successfully!")
}

func createTables(db *sql.DB) error {
	// Create contestants table
	contestantsSQL := `
	CREATE TABLE IF NOT EXISTS contestants (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		badge_id TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		company TEXT,
		email TEXT,
		phone TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(contestantsSQL); err != nil {
		return fmt.Errorf("failed to create contestants table: %v", err)
	}

	// Create penalty_shots table
	penaltyShotsSQL := `
	CREATE TABLE IF NOT EXISTS penalty_shots (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		contestant_id INTEGER NOT NULL,
		shot_result TEXT NOT NULL CHECK (shot_result IN ('goal', 'miss')),
		attempt_number INTEGER NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		operator_id TEXT,
		session_id TEXT,
		FOREIGN KEY (contestant_id) REFERENCES contestants(id)
	);`

	if _, err := db.Exec(penaltyShotsSQL); err != nil {
		return fmt.Errorf("failed to create penalty_shots table: %v", err)
	}

	// Create leaderboard_cache table
	leaderboardCacheSQL := `
	CREATE TABLE IF NOT EXISTS leaderboard_cache (
		contestant_id INTEGER PRIMARY KEY,
		total_attempts INTEGER DEFAULT 0,
		successful_shots INTEGER DEFAULT 0,
		success_percentage REAL DEFAULT 0.0,
		last_updated DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (contestant_id) REFERENCES contestants(id)
	);`

	if _, err := db.Exec(leaderboardCacheSQL); err != nil {
		return fmt.Errorf("failed to create leaderboard_cache table: %v", err)
	}

	return nil
}

func createIndexes(db *sql.DB) error {
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_contestants_badge_id ON contestants(badge_id);",
		"CREATE INDEX IF NOT EXISTS idx_penalty_shots_contestant_id ON penalty_shots(contestant_id);",
		"CREATE INDEX IF NOT EXISTS idx_penalty_shots_timestamp ON penalty_shots(timestamp);",
		"CREATE INDEX IF NOT EXISTS idx_leaderboard_success_percentage ON leaderboard_cache(success_percentage DESC);",
	}

	for _, indexSQL := range indexes {
		if _, err := db.Exec(indexSQL); err != nil {
			return fmt.Errorf("failed to create index: %v", err)
		}
	}

	return nil
}
