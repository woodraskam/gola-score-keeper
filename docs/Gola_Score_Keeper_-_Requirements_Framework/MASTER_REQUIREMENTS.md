# Gola Score Keeper - Master Requirements Specification

## Overview
This document establishes the single source of truth for all requirements across the Gola Score Keeper system. All module-specific requirements must align with these master specifications.

## Performance Requirements (Standardized)

### Response Time Targets
- **Badge Scanning**: < 5 seconds (baseline requirement)
- **UI Response**: < 2 seconds for all operations
- **Database Queries**: < 100ms for single record lookups, < 2 seconds for complex queries
- **Real-time Updates**: < 1 second propagation via WebSocket

### Throughput Requirements
- **Contestant Capacity**: 200+ contestants per day
- **Concurrent Operations**: Support 50+ concurrent users
- **Shot Recording**: 120 shots per hour per booth
- **Badge Processing**: 200+ scans per day

### Resource Usage (Standardized)
- **Memory Usage**: < 200MB baseline (Apple Silicon optimized)
- **CPU Usage**: < 10% during normal operation
- **Storage**: < 10GB for 1000+ contestants with full history
- **Battery Life**: 8+ hours continuous operation on MacBook Air

## Database Schema (Standardized)

### Core Tables
```sql
-- Contestants table
CREATE TABLE contestants (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    badge_id TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    company TEXT,
    email TEXT,
    phone TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Penalty shots table (standardized naming)
CREATE TABLE penalty_shots (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    contestant_id INTEGER NOT NULL,
    shot_result TEXT NOT NULL CHECK (shot_result IN ('goal', 'miss')),
    attempt_number INTEGER NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    operator_id TEXT,
    session_id TEXT,
    FOREIGN KEY (contestant_id) REFERENCES contestants(id)
);

-- Leaderboard cache table
CREATE TABLE leaderboard_cache (
    contestant_id INTEGER PRIMARY KEY,
    total_attempts INTEGER DEFAULT 0,
    successful_shots INTEGER DEFAULT 0,
    success_percentage REAL DEFAULT 0.0,
    last_updated DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (contestant_id) REFERENCES contestants(id)
);
```

### Indexes (Standardized)
```sql
CREATE INDEX idx_contestants_badge_id ON contestants(badge_id);
CREATE INDEX idx_penalty_shots_contestant_id ON penalty_shots(contestant_id);
CREATE INDEX idx_penalty_shots_timestamp ON penalty_shots(timestamp);
CREATE INDEX idx_leaderboard_success_percentage ON leaderboard_cache(success_percentage DESC);
```

## API Endpoints (Standardized)

### Core Endpoints
- `POST /api/scan-badge` - Badge scanning and OCR processing
- `GET /api/contestants/{id}` - Retrieve contestant by ID
- `POST /api/contestants` - Register new contestant
- `POST /api/penalty-shots` - Record penalty shot result
- `GET /api/leaderboard` - Retrieve current leaderboard
- `GET /api/export/{format}` - Export data (CSV/JSON)

### WebSocket Endpoints
- `ws://localhost:8080/ws/leaderboard` - Real-time leaderboard updates
- `ws://localhost:8080/ws/scoring` - Real-time scoring updates

## OCR Requirements (Standardized)

### Accuracy Targets
- **Baseline OCR Accuracy**: 90% minimum for standard badges
- **Neural Engine Enhanced**: 95%+ when Apple Silicon Neural Engine is available
- **Fallback Threshold**: 85% minimum before manual intervention required

### Processing Requirements
- **Processing Time**: < 5 seconds per badge scan
- **Confidence Threshold**: 0.7 minimum for automatic acceptance
- **Manual Override**: Available for confidence scores below threshold

## Technology Stack (Standardized)

### Core Technologies
- **Language**: Go 1.21+
- **Database**: SQLite 3.40+
- **Web Framework**: Gin v1.9.1
- **OCR Engine**: Tesseract 5.0+ with Apple Vision Framework fallback
- **Platform**: macOS Monterey 12.0+ (Apple Silicon optimized)

### Dependencies
```go
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/mattn/go-sqlite3 v1.14.17
    modernc.org/sqlite v1.25.0
    gocv.io/x/gocv v0.34.0
)
```

## Error Handling (Standardized)

### Error Categories
1. **System Errors**: Database failures, hardware issues
2. **Validation Errors**: Invalid data, missing required fields
3. **Business Logic Errors**: Duplicate contestants, invalid operations
4. **Network Errors**: WebSocket disconnections, API timeouts

### Error Response Format
```json
{
    "error": {
        "code": "ERROR_CODE",
        "message": "Human readable error message",
        "details": "Technical details for debugging",
        "timestamp": "2024-01-15T10:30:00Z"
    }
}
```

## Security Requirements (Standardized)

### Authentication
- Session-based authentication for booth operators
- PIN-based access control for booth environment
- Administrator override capabilities

### Data Protection
- Local data encryption using AES-256
- GDPR compliance for contestant data
- Secure badge image handling with automatic cleanup
- Audit logging for all data modifications

## Integration Requirements (Standardized)

### Module Communication
- **HTTP REST APIs** for synchronous operations
- **WebSocket connections** for real-time updates
- **Shared database** for data persistence
- **Event-driven architecture** for loose coupling

### Data Flow
```
Badge Scan → OCR Processing → Contestant Registration → 
Score Recording → Leaderboard Update → Real-time Display
```

## Testing Requirements (Standardized)

### Test Coverage
- **Unit Tests**: 80% minimum coverage
- **Integration Tests**: All API endpoints and database operations
- **Performance Tests**: Load testing with 200+ concurrent users
- **User Acceptance Tests**: Booth operator workflow validation

### Test Data
- **Sample Badges**: 50+ different badge formats for OCR testing
- **Contestant Data**: 1000+ test contestants for performance testing
- **Score Data**: 5000+ penalty shot attempts for leaderboard testing

## Deployment Requirements (Standardized)

### Environment Configuration
- **Development**: Local SQLite, mock camera, debug logging
- **Staging**: Production-like database, actual camera hardware
- **Production**: Optimized configuration, minimal logging, monitoring

### Build Requirements
- **Native ARM64 compilation** for Apple Silicon
- **CGO enabled** for Apple framework integration
- **Dependency bundling** for standalone deployment
- **Configuration validation** on startup

## Monitoring and Logging (Standardized)

### Metrics Collection
- Badge scanning success rate and processing time
- Contestant registration volume and accuracy
- Penalty shot recording frequency and response time
- Leaderboard update latency and WebSocket connections
- System resource utilization (CPU, memory, disk)

### Logging Levels
- **DEBUG**: Detailed operation tracing
- **INFO**: Normal operation events
- **WARN**: Non-critical issues
- **ERROR**: System errors requiring attention

## Compliance Requirements (Standardized)

### Regulatory Compliance
- **GDPR**: EU data protection compliance
- **Accessibility**: WCAG 2.1 AA standards
- **Data Retention**: Configurable retention policies
- **Audit Trail**: Complete operation logging

### Business Requirements
- **Offline Operation**: Full functionality without network
- **Event Duration**: 8+ hours continuous operation
- **Training Time**: < 15 minutes for booth operators
- **Recovery Time**: < 30 seconds from transient failures

---

**Document Version**: 1.0  
**Last Updated**: 2024-01-15  
**Next Review**: 2024-02-15  
**Approved By**: Project Lead
