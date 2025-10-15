# Penalty Shot Scoring Interface - README Documentation

## Overview
Comprehensive README documentation for Penalty Shot Scoring Interface, providing essential information for developers, users, and stakeholders.

## Project Description
Penalty Shot Scoring Interface is a functional component that provides an intuitive web-based interface for recording penalty shot results (goal/miss) for contestants in the Gola Score Keeper application. This interface enables booth operators to quickly capture shot outcomes with simple input controls while displaying current contestant information.

## Features
- **Core Functionality**: Real-time penalty shot result recording with goal/miss buttons
- **Key Capabilities**: 
  - Quick shot result entry (<2 seconds per entry)
  - Current contestant information display
  - Real-time score updates
  - Touch-friendly interface for booth environments
  - Error handling and validation
- **Integration Points**: Connects with Score Management Engine, Contestant Database, and Real-time Display systems

## Getting Started

### Prerequisites
- Go 1.21+ installed
- SQLite 3.40+ database
- Modern web browser with JavaScript enabled
- Active contestant session in Gola Score Keeper

### Installation
1. Ensure the main Gola Score Keeper application is running
2. Navigate to the web interface at `http://localhost:8080`
3. Access the penalty shot scoring module from the main dashboard

### Quick Start
```go
// Initialize scoring interface
scoringInterface := &ScoringInterface{
    ContestantID: currentContestant.ID,
    SessionID:    activeSession.ID,
}

// Record a goal
result := scoringInterface.RecordShot(ShotResult{
    Type:      "goal",
    Timestamp: time.Now(),
})
```

## Usage

### Basic Usage
1. **Select Active Contestant**: Ensure a contestant is currently active in the system
2. **Display Interface**: The scoring interface automatically shows current contestant details
3. **Record Shot**: Click either "GOAL" or "MISS" button to record the penalty shot result
4. **Confirm Entry**: System provides visual feedback and updates scores in real-time
5. **Continue or Finish**: Proceed with additional shots or complete the contestant's session

### Advanced Configuration
```javascript
// Configure scoring interface options
const scoringConfig = {
    autoAdvance: true,          // Auto-advance to next contestant
    confirmationDialog: false,  // Skip confirmation for quick entry
    soundEffects: true,         // Enable audio feedback
    touchOptimized: true,       // Optimize for touch screens
    displayTimeout: 30000       // Auto-hide after 30 seconds of inactivity
};
```

### API Reference
- `POST /api/penalty-shots`: Record penalty shot result (standardized endpoint)
- `GET /api/contestants/current`: Retrieve current contestant information
- `PUT /api/penalty-shots/{id}`: Update existing shot record
- `DELETE /api/penalty-shots/{id}`: Remove shot record (admin only)

## Configuration

### Environment Variables
- `SCORING_TIMEOUT`: Maximum time to wait for shot entry (default: 60s)
- `AUTO_ADVANCE`: Automatically advance to next contestant (default: false)
- `TOUCH_MODE`: Enable touch-optimized interface (default: true)
- `CONFIRMATION_REQUIRED`: Require confirmation before recording (default: false)

### Settings
```json
{
  "interface": {
    "buttonSize": "large",
    "colorScheme": "high-contrast",
    "animationSpeed": "fast",
    "feedbackType": "visual+audio"
  },
  "scoring": {
    "allowUndo": true,
    "maxUndoTime": 10,
    "requireConfirmation": false
  },
  "display": {
    "showContestantPhoto": true,
    "showCompanyInfo": true,
    "showCurrentStats": true
  }
}
```

## Examples

### Example 1: Recording a Goal
```javascript
// JavaScript frontend code
async function recordGoal() {
    const response = await fetch('/api/shots/record', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            contestant_id: currentContestant.id,
            shot_result: 'goal',
            timestamp: new Date().toISOString()
        })
    });
    
    if (response.ok) {
        showSuccessMessage('Goal recorded!');
        updateScoreDisplay();
    }
}
```

### Example 2: Backend Shot Processing
```go
func (h *ScoringHandler) RecordShot(c *gin.Context) {
    var shotData ShotRequest
    if err := c.ShouldBindJSON(&shotData); err != nil {
        c.JSON(400, gin.H{"error": "Invalid shot data"})
        return
    }
    
    shot := &models.PenaltyShot{
        ContestantID: shotData.ContestantID,
        ShotResult:   shotData.ShotResult, // Standardized field name
        Timestamp:    time.Now(),
        OperatorID:   shotData.OperatorID,
        SessionID:    shotData.SessionID,
    }
    
    if err := h.db.CreatePenaltyShot(shot); err != nil {
        c.JSON(500, gin.H{"error": "Failed to record shot"})
        return
    }
    
    // Update leaderboard in real-time
    h.broadcastScoreUpdate(shot)
    c.JSON(200, gin.H{"success": true, "shot_id": shot.ID})
}
```

## Troubleshooting

### Common Issues
1. **Issue**: Buttons not responding to clicks
   **Solution**: Check JavaScript console for errors, ensure contestant is selected, verify network connectivity

2. **Issue**: Contestant information not displaying
   **Solution**: Verify contestant is properly scanned and active, check database connection, refresh the interface

3. **Issue**: Scores not updating in real-time
   **Solution**: Check WebSocket connection, verify browser compatibility, restart the scoring service

4. **Issue**: Interface appears too small on touch devices
   **Solution**: Enable touch mode in configuration, adjust button size settings, use browser zoom if needed

### Debugging
- Enable debug mode: `DEBUG=true go run cmd/server/main.go`
- Check browser developer tools for JavaScript errors
- Monitor network requests in browser dev tools
- Review application logs: `tail -f logs/scoring.log`
- Test database connectivity: `go run cmd/dbtest/main.go`

## Contributing
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/scoring-enhancement`
3. Follow Go coding standards and include tests
4. Ensure all existing tests pass: `go test ./...`
5. Submit a pull request with detailed description
6. Code review required before merging

## License
This component is part of the Gola Score Keeper application and follows the same licensing terms as the main project.

## Support
- **Technical Issues**: Create an issue in the GitHub repository
- **Feature Requests**: Submit enhancement requests through GitHub issues
- **Documentation**: Refer to the main Gola Score Keeper documentation
- **Emergency Support**: Contact the development team during event hours

---

**Component Version**: 1.0.0  
**Last Updated**: 2024-01-15  
**Compatibility**: Gola Score Keeper v1.0+  
**Browser Support**: Chrome 90+, Firefox 88+, Safari 14+