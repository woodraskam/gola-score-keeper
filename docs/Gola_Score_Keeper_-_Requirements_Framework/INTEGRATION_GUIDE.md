# Goal Score Keeper - Cross-Module Integration Guide

## Overview
This document describes how all modules in the Goal Score Keeper system integrate together, including data flow, API interactions, and shared components.

## System Architecture Integration

### Module Dependencies
```
┌─────────────────────────────────────────────────────────────┐
│                    Goal Score Keeper                        │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────┐ │
│  │ Apple Silicon   │  │ Badge Scanning  │  │ OCR         │ │
│  │ Optimization    │  │ Camera          │  │ Processing  │ │
│  └─────────────────┘  └─────────────────┘  └─────────────┘ │
│           │                     │                     │     │
│           └─────────────────────┼─────────────────────┘     │
│                                 │                         │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────┐ │
│  │ Contestant      │  │ Penalty Shot   │  │ Leaderboard │ │
│  │ Registration    │  │ Scoring        │  │ Display     │ │
│  └─────────────────┘  └─────────────────┘  └─────────────┘ │
│           │                     │                     │     │
│           └─────────────────────┼─────────────────────┘     │
│                                 │                         │
│  ┌─────────────────────────────────────────────────────────┐ │
│  │              Competitor History Database                │ │
│  └─────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

## Data Flow Integration

### 1. Badge Scanning to Contestant Registration
```
Camera Capture → OCR Processing → Data Validation → Contestant Registration
     │                │                │                    │
     ▼                ▼                ▼                    ▼
Badge Image → Text Extraction → Field Validation → Database Storage
```

**Integration Points:**
- Camera module provides image data to OCR module
- OCR module returns structured contestant data
- Registration module validates and stores contestant information
- Database module persists contestant data

### 2. Contestant Registration to Score Recording
```
Contestant Active → Penalty Shot Interface → Score Recording → Database Update
        │                    │                    │                │
        ▼                    ▼                    ▼                ▼
Contestant Info → Shot Result Entry → Score Calculation → Leaderboard Update
```

**Integration Points:**
- Registration module provides active contestant information
- Scoring interface records shot results
- Database module updates penalty shot records
- Leaderboard module recalculates rankings

### 3. Real-time Updates via WebSocket
```
Score Recording → Database Update → WebSocket Broadcast → UI Updates
       │                │                    │                │
       ▼                ▼                    ▼                ▼
Shot Result → Penalty Shot Record → Real-time Event → Leaderboard Refresh
```

## API Integration Matrix

### Core API Endpoints (Standardized)
| Endpoint | Method | Module | Purpose |
|----------|--------|--------|---------|
| `/api/scan-badge` | POST | Badge Scanning | Process badge image |
| `/api/contestants` | POST | Registration | Register new contestant |
| `/api/contestants/{id}` | GET | Registration | Get contestant info |
| `/api/penalty-shots` | POST | Scoring | Record shot result |
| `/api/leaderboard` | GET | Leaderboard | Get current rankings |
| `/api/export/{format}` | GET | Database | Export data |

### WebSocket Endpoints
| Endpoint | Module | Purpose |
|----------|--------|---------|
| `/ws/leaderboard` | Leaderboard | Real-time leaderboard updates |
| `/ws/scoring` | Scoring | Real-time score updates |

## Database Schema Integration

### Shared Tables
All modules use the same database schema defined in `MASTER_REQUIREMENTS.md`:

```sql
-- Core tables used by all modules
contestants (id, badge_id, name, company, email, phone, created_at, updated_at)
penalty_shots (id, contestant_id, shot_result, attempt_number, timestamp, operator_id, session_id)
leaderboard_cache (contestant_id, total_attempts, successful_shots, success_percentage, last_updated)
```

### Module-Specific Data Access
- **Registration Module**: CRUD operations on `contestants` table
- **Scoring Module**: INSERT operations on `penalty_shots` table
- **Leaderboard Module**: READ operations on `leaderboard_cache` table
- **Database Module**: All operations with transaction management

## Shared Components Integration

### 1. Database Connection Pool
```go
// Shared across all modules
type DatabasePool struct {
    connections chan *sql.DB
    maxConns    int
    timeout     time.Duration
}
```

### 2. WebSocket Hub
```go
// Shared real-time communication
type WebSocketHub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}
```

### 3. Configuration Management
```go
// Shared configuration across modules
type Config struct {
    Database   DatabaseConfig
    OCR        OCRConfig
    Camera     CameraConfig
    WebSocket  WebSocketConfig
}
```

## Error Handling Integration

### Error Propagation Chain
```
Module Error → Error Handler → Logging System → User Notification
     │              │              │                │
     ▼              ▼              ▼                ▼
Local Error → Centralized Handler → Structured Log → UI Feedback
```

### Shared Error Types
```go
type SystemError struct {
    Code    string
    Message string
    Module  string
    Details map[string]interface{}
}
```

## Performance Integration

### Shared Performance Metrics
- **Response Time**: < 2 seconds for all operations
- **Memory Usage**: < 200MB baseline (Apple Silicon optimized)
- **Database Queries**: < 100ms for single record lookups
- **WebSocket Latency**: < 1 second for real-time updates

### Monitoring Integration
```go
type PerformanceMonitor struct {
    ResponseTime time.Duration
    MemoryUsage  uint64
    CPUUsage     float64
    DBQueries    int
    WSClients    int
}
```

## Security Integration

### Authentication Flow
```
User Login → Session Creation → Module Access → Permission Check
     │              │                │              │
     ▼              ▼                ▼              ▼
Credentials → JWT Token → Module Request → Authorization
```

### Shared Security Components
- **Session Management**: Centralized session handling
- **Permission System**: Role-based access control
- **Data Encryption**: AES-256 for sensitive data
- **Audit Logging**: Complete operation tracking

## Testing Integration

### Integration Test Scenarios
1. **End-to-End Badge Scanning**: Camera → OCR → Registration → Database
2. **Score Recording Flow**: Contestant → Shot → Database → Leaderboard
3. **Real-time Updates**: Score Change → WebSocket → UI Update
4. **Error Recovery**: Network Failure → Offline Mode → Sync

### Shared Test Utilities
```go
type IntegrationTestSuite struct {
    Database   *TestDatabase
    WebSocket  *TestWebSocket
    Camera     *MockCamera
    OCR        *MockOCR
}
```

## Deployment Integration

### Module Startup Sequence
1. **Database Module**: Initialize SQLite connection
2. **Camera Module**: Initialize AVFoundation
3. **OCR Module**: Load Tesseract models
4. **WebSocket Module**: Start WebSocket server
5. **API Module**: Start HTTP server
6. **UI Module**: Serve static files

### Configuration Integration
```yaml
# Shared configuration file
database:
  path: "./data/score_keeper.db"
  max_connections: 10

camera:
  device_id: 0
  resolution: "1920x1080"

ocr:
  language: "eng"
  confidence_threshold: 0.7

websocket:
  port: 8080
  path: "/ws"
```

## Troubleshooting Integration

### Cross-Module Debugging
1. **Enable Debug Logging**: Set `LOG_LEVEL=debug` for all modules
2. **Monitor WebSocket Connections**: Check `/ws/status` endpoint
3. **Database Health Check**: Use `/api/health/database` endpoint
4. **Performance Monitoring**: Access `/debug/pprof/` for profiling

### Common Integration Issues
1. **Database Lock**: Multiple modules accessing database simultaneously
2. **WebSocket Disconnection**: Network issues affecting real-time updates
3. **Memory Leaks**: Long-running operations without proper cleanup
4. **OCR Timeout**: Image processing taking longer than 5 seconds

## Maintenance Integration

### Update Procedures
1. **Database Migrations**: Run schema updates before module updates
2. **Configuration Changes**: Update shared config file
3. **Module Restart**: Graceful shutdown and restart sequence
4. **Data Backup**: Export data before major updates

### Monitoring Integration
- **Health Checks**: All modules report status to central monitor
- **Performance Metrics**: Shared metrics collection
- **Error Tracking**: Centralized error logging and alerting
- **Resource Usage**: Memory, CPU, and disk usage monitoring

---

**Document Version**: 1.0  
**Last Updated**: 2024-01-15  
**Next Review**: 2024-02-15  
**Maintained By**: Development Team
