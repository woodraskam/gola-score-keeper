# Competitor History Database - README Documentation

## Overview
Comprehensive README documentation for Competitor History Database, providing essential information for developers, users, and stakeholders.

## Project Description
Competitor History Database is a local storage component that manages the complete historical record of all contestants participating in the Gola Score Keeper penalty shot competition. It provides persistent storage, efficient retrieval, and comprehensive tracking of contestant attempts, scores, and performance metrics.

## Features
- **Persistent Data Storage**: Complete historical record of all contestant interactions
- **Performance Tracking**: Detailed scoring history with timestamps and attempt sequences  
- **Efficient Indexing**: Optimized database indexes for fast contestant lookup and leaderboard generation
- **Data Integrity**: ACID-compliant transactions ensuring data consistency
- **Export Capabilities**: Support for CSV and JSON data export formats
- **Analytics Support**: Historical data aggregation for performance analytics

## Getting Started

### Prerequisites
- Go 1.21 or higher
- SQLite 3.40+
- macOS Monterey 12.0+ (Apple Silicon optimized)
- CGO enabled for SQLite compilation

### Installation
1. Initialize the database schema using the setup utility
2. Configure database connection parameters in environment variables
3. Run database migrations to ensure latest schema version

### Quick Start
```go
db, err := database.NewCompetitorDB("./data/competitors.db")
if err != nil {
    log.Fatal("Failed to initialize database:", err)
}

// Store penalty shot attempt
attempt := &models.PenaltyShot{
    ContestantID: 1, // Reference to contestants table
    ShotResult: "goal", // Standardized field name
    AttemptNumber: 1,
    Timestamp: time.Now(),
    OperatorID: "operator-001",
    SessionID: "session-123",
}

err = db.StoreAttempt(attempt)
if err != nil {
    log.Error("Failed to store attempt:", err)
}
```

## Usage

### Basic Usage
The Competitor History Database integrates seamlessly with the Gola Score Keeper application to provide persistent storage for all contestant data:

1. **Contestant Registration**: Automatically stores new contestants when badges are scanned
2. **Attempt Recording**: Records each penalty shot attempt with detailed metadata
3. **Score Tracking**: Maintains running totals and performance statistics
4. **Historical Queries**: Supports complex queries for leaderboards and analytics

### Advanced Configuration
Configure database performance and behavior through environment variables and initialization parameters:

```go
config := &database.Config{
    Path: "./data/competitors.db",
    MaxConnections: 10,
    EnableWAL: true,
    CacheSize: 2000,
    BusyTimeout: 30 * time.Second,
}

db, err := database.NewCompetitorDBWithConfig(config)
```

### API Reference
**Core Methods:**
- `StoreAttempt(attempt *CompetitorAttempt) error`: Record new attempt
- `GetContestant(badgeID string) (*Contestant, error)`: Retrieve contestant by badge ID
- `GetLeaderboard(limit int) ([]*LeaderboardEntry, error)`: Generate ranked leaderboard
- `GetAttemptHistory(badgeID string) ([]*CompetitorAttempt, error)`: Get all attempts for contestant
- `ExportData(format string) ([]byte, error)`: Export all data in specified format

## Configuration

### Environment Variables
- `COMPETITOR_DB_PATH`: Database file location (default: "./data/competitors.db")
- `COMPETITOR_DB_CACHE_SIZE`: SQLite cache size in KB (default: 2000)
- `COMPETITOR_DB_BUSY_TIMEOUT`: Database busy timeout in seconds (default: 30)
- `COMPETITOR_DB_ENABLE_WAL`: Enable Write-Ahead Logging (default: true)

### Settings
**Database Schema Configuration:**
- Indexed fields: badge_id, timestamp, score, company
- Foreign key constraints: enabled for data integrity
- Auto-vacuum: incremental mode for space management
- Journal mode: WAL for concurrent read performance

## Examples

### Example 1: Recording Penalty Shot Attempt
```go
// Create new penalty shot attempt
attempt := &models.PenaltyShot{
    ContestantID: 2, // Reference to contestants table
    ShotResult: "miss", // Standardized field name
    AttemptNumber: 1,
    Timestamp: time.Now(),
    OperatorID: "operator-002",
    SessionID: "session-456",
}

// Store in database
if err := db.StoreAttempt(attempt); err != nil {
    log.Printf("Error storing attempt: %v", err)
    return err
}

log.Printf("Successfully recorded attempt for %s", attempt.Name)
```

### Example 2: Generating Leaderboard
```go
// Get top 10 contestants
leaderboard, err := db.GetLeaderboard(10)
if err != nil {
    log.Printf("Error generating leaderboard: %v", err)
    return err
}

// Display results
fmt.Println("Current Leaderboard:")
for i, entry := range leaderboard {
    fmt.Printf("%d. %s (%s) - Score: %d\n", 
        i+1, entry.Name, entry.Company, entry.BestScore)
}
```

## Troubleshooting

### Common Issues
1. **Issue**: Database locked error during concurrent operations
   **Solution**: Ensure WAL mode is enabled and implement proper connection pooling. Check busy timeout configuration.

2. **Issue**: Slow query performance on large datasets
   **Solution**: Verify indexes are created on frequently queried columns (badge_id, timestamp, score). Run ANALYZE command periodically.

3. **Issue**: Database file corruption after unexpected shutdown
   **Solution**: Enable WAL mode and implement proper transaction handling. Consider implementing database backup strategy.

4. **Issue**: Memory usage grows during extended operation
   **Solution**: Configure appropriate cache size and enable auto-vacuum. Implement connection lifecycle management.

### Debugging
Enable detailed logging for database operations:
```go
db.SetLogLevel(database.LogLevelDebug)
```

Monitor database performance:
```go
stats := db.GetStatistics()
log.Printf("Total queries: %d, Average response time: %v", 
    stats.QueryCount, stats.AvgResponseTime)
```

Use SQLite EXPLAIN QUERY PLAN for query optimization:
```sql
EXPLAIN QUERY PLAN SELECT * FROM attempts WHERE badge_id = ? ORDER BY timestamp DESC;
```

## Contributing
1. Follow Go coding standards and include comprehensive tests
2. Ensure database migrations are backwards compatible
3. Update schema documentation for any structural changes
4. Include performance benchmarks for new query operations
5. Test with realistic data volumes (200+ contestants, 1000+ attempts)

## License
Licensed under the same terms as the Gola Score Keeper project. See LICENSE file for details.

## Support
For database-related issues:
- Check application logs for detailed error messages
- Verify database file permissions and disk space
- Review SQLite documentation for advanced configuration
- Contact development team for schema modifications or performance optimization