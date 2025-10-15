# Leaderboard Display - README Documentation

## Overview
Comprehensive README documentation for Leaderboard Display, providing essential information for developers, users, and stakeholders.

## Project Description
Leaderboard Display is a functional component that provides a real-time ranking view of contestants in the Gola Score Keeper application, showing top performers ranked by success rate or total goals with comprehensive filtering and sorting capabilities.

## Features
- **Core Functionality**: Real-time leaderboard displaying top 10 contestants with automatic refresh
- **Key Capabilities**: 
  - Multiple sorting options (success rate, total goals, alphabetical)
  - Filtering by company, date range, or performance metrics
  - Auto-refresh functionality for live updates
  - Responsive design for various display sizes
  - Export capabilities for leaderboard data
- **Integration Points**: Integrates with Score Management Engine, Contestant Database, and Web Interface Layer

## Getting Started

### Prerequisites
- Go 1.21+ installed
- SQLite 3.40+ database
- Web browser with WebSocket support
- Gola Score Keeper core application running

### Installation
1. Ensure the main Gola Score Keeper application is installed
2. Navigate to the project directory: `cd gola-score-keeper`
3. Install dependencies: `go mod tidy`

### Quick Start
```go
// Initialize leaderboard service
leaderboard := NewLeaderboardService(db)

// Get top 10 contestants by success rate
topContestants, err := leaderboard.GetTopContestants(10, "success_rate")
if err != nil {
    log.Fatal(err)
}

// Display leaderboard
for i, contestant := range topContestants {
    fmt.Printf("%d. %s - %s (%.2f%% success)\n", 
        i+1, contestant.Name, contestant.Company, contestant.SuccessRate)
}
```

## Usage

### Basic Usage
1. Access the leaderboard through the main web interface at `/leaderboard`
2. View the current top 10 performers displayed in real-time
3. Use the sorting dropdown to change ranking criteria
4. Apply filters using the filter panel on the left side
5. The display automatically refreshes every 30 seconds

### Advanced Configuration
```go
// Configure leaderboard settings
config := LeaderboardConfig{
    MaxDisplayCount:    10,
    RefreshInterval:    30 * time.Second,
    DefaultSortBy:      "success_rate",
    EnableFiltering:    true,
    EnableExport:       true,
    ShowCompanyLogos:   true,
}

leaderboard := NewLeaderboardService(db, config)
```

### API Reference
- `GetTopContestants(limit int, sortBy string) ([]Contestant, error)`
- `FilterContestants(filters FilterCriteria) ([]Contestant, error)`
- `GetLeaderboardStats() (LeaderboardStats, error)`
- `ExportLeaderboard(format string) ([]byte, error)`

## Configuration

### Environment Variables
- `LEADERBOARD_REFRESH_INTERVAL`: Refresh interval in seconds (default: 30)
- `LEADERBOARD_MAX_DISPLAY`: Maximum contestants to display (default: 10)
- `LEADERBOARD_ENABLE_EXPORT`: Enable export functionality (default: true)

### Settings
```json
{
  "leaderboard": {
    "maxDisplay": 10,
    "refreshInterval": 30,
    "sortOptions": ["success_rate", "total_goals", "attempts", "name"],
    "filterOptions": ["company", "date_range", "min_attempts"],
    "displayColumns": ["rank", "name", "company", "success_rate", "total_goals"],
    "enableAnimations": true,
    "showAvatars": false
  }
}
```

## Examples

### Example 1: Custom Leaderboard Query
```go
// Get top 5 contestants from specific company
filters := FilterCriteria{
    Company:     "TechCorp",
    MinAttempts: 3,
    DateRange:   DateRange{Start: today, End: today},
}

contestants, err := leaderboard.FilterContestants(filters)
if err != nil {
    return err
}

topFive := contestants[:min(5, len(contestants))]
```

### Example 2: WebSocket Real-time Updates
```javascript
// Connect to leaderboard WebSocket
const ws = new WebSocket('ws://localhost:8080/ws/leaderboard');

ws.onmessage = function(event) {
    const leaderboardData = JSON.parse(event.data);
    updateLeaderboardDisplay(leaderboardData);
};

function updateLeaderboardDisplay(data) {
    const tbody = document.getElementById('leaderboard-table-body');
    tbody.innerHTML = '';
    
    data.contestants.forEach((contestant, index) => {
        const row = createLeaderboardRow(contestant, index + 1);
        tbody.appendChild(row);
    });
}
```

## Troubleshooting

### Common Issues
1. **Issue**: Leaderboard not refreshing automatically
   **Solution**: Check WebSocket connection status and ensure the refresh interval is properly configured. Verify network connectivity and restart the service if needed.

2. **Issue**: Incorrect rankings displayed
   **Solution**: Verify the sorting algorithm implementation and check for data consistency in the database. Run `ANALYZE` command on SQLite to update statistics.

3. **Issue**: Slow leaderboard loading with large datasets
   **Solution**: Implement database indexing on ranking columns, add pagination, and consider caching frequently accessed leaderboard data.

### Debugging
- Enable debug logging: `export LOG_LEVEL=debug`
- Check database queries: Monitor SQLite query performance
- Verify WebSocket connections: Use browser developer tools to inspect WebSocket traffic
- Test with sample data: Use the built-in test data generator for debugging

## Contributing
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/leaderboard-enhancement`
3. Follow Go coding standards and include unit tests
4. Ensure all tests pass: `go test ./internal/leaderboard/...`
5. Submit a pull request with detailed description of changes
6. Update documentation as needed

## License
This component is part of the Gola Score Keeper application and follows the same licensing terms as the main project.

## Support
- **Documentation**: See main project documentation at `/docs`
- **Issues**: Report bugs via GitHub issues
- **Email**: Contact project maintainers at support@company.com
- **Slack**: Join #gola-score-keeper channel for real-time support