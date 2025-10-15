# Gola Score Keeper - Requirements Framework - Design Specification

## Design Overview
The Gola Score Keeper system is designed as a modular, event-driven application optimized for trade show environments. The architecture emphasizes real-time badge scanning, score tracking, and leaderboard management with offline-first capabilities. The system leverages Go's concurrency features for camera processing, OCR operations, and WebSocket communications while maintaining data persistence through SQLite for reliable booth operations.

## Architecture

### System Architecture
```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Camera Layer  │    │   Web Interface  │    │  Export Layer   │
│                 │    │                  │    │                 │
│ • AVFoundation  │    │ • Gin Framework  │    │ • CSV/JSON      │
│ • Image Capture │    │ • WebSocket      │    │ • Report Gen    │
│ • Preprocessing │    │ • Real-time UI   │    │ • Analytics     │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Core Application Layer                       │
│                                                                 │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────────┐ │
│  │ OCR Service │  │Score Service│  │   Contestant Service    │ │
│  │             │  │             │  │                         │ │
│  │• Tesseract  │  │• Tracking   │  │ • Registration          │ │
│  │• Parsing    │  │• Validation │  │ • Validation            │ │
│  │• Confidence │  │• Analytics  │  │ • Duplicate Detection   │ │
│  └─────────────┘  └─────────────┘  └─────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
                 ┌─────────────────────────┐
                 │    Data Persistence     │
                 │                         │
                 │ • SQLite Database       │
                 │ • Local File Storage    │
                 │ • Backup Mechanisms     │
                 └─────────────────────────┘
```

### Component Architecture
- **Camera Interface Component**: Manages AVFoundation integration, image capture optimization for Apple Silicon, and real-time preview functionality for booth operators
- **OCR Processing Engine**: Handles Tesseract integration, badge format recognition, text extraction with confidence scoring, and fallback manual entry capabilities
- **Score Management System**: Implements penalty shot tracking, real-time leaderboard calculations, performance analytics, and historical data management
- **Web Interface Layer**: Provides responsive operator dashboard, WebSocket-based real-time updates, touch-friendly controls, and booth display optimization
- **Data Persistence Layer**: Manages SQLite operations, data validation, backup mechanisms, and offline-first data synchronization

### Data Flow
```
Badge Scan Request → Camera Capture → Image Preprocessing → 
OCR Text Extraction → Badge Data Parsing → Contestant Validation → 
Database Storage → Score Entry → Real-time Leaderboard Update → 
WebSocket Broadcast → UI Refresh → Export Generation
```

## User Interface Design

### Wireframes
**Main Dashboard**: Split-screen layout with camera preview (left 40%), contestant information panel (center 35%), and live leaderboard (right 25%). Large touch-friendly buttons for score entry positioned below camera preview.

**Badge Scanning Interface**: Full-screen camera preview with overlay guides for optimal badge positioning, confidence indicator, and manual override button for failed scans.

**Leaderboard Display**: Responsive table with ranking, contestant name, company, attempts, successful shots, and percentage. Auto-refresh every 2 seconds with smooth transition animations.

### User Experience Flow
1. **Badge Scanning** → System captures image, processes OCR, displays contestant info with 3-second confirmation timeout
2. **Score Entry** → Operator selects penalty shot result (goal/miss), system validates and stores with timestamp
3. **Leaderboard Update** → Real-time ranking recalculation, WebSocket broadcast to all connected displays, smooth UI animations

### Design Principles
- Touch-first interaction design optimized for booth environment
- High-contrast visual elements for trade show lighting conditions
- Minimal cognitive load with clear visual hierarchy and status indicators
- Responsive design supporting tablet, laptop, and external display configurations

## Technical Design

### Database Design
**Tables**:
- `contestants`: Contestant information and registration data
  - `id`: INTEGER PRIMARY KEY - Unique contestant identifier
  - `badge_id`: TEXT UNIQUE - Scanned badge identifier
  - `name`: TEXT NOT NULL - Contestant full name
  - `company`: TEXT - Company/organization name
  - `email`: TEXT - Contact email address
  - `phone`: TEXT - Contact phone number
  - `created_at`: DATETIME - Registration timestamp
  - `updated_at`: DATETIME - Last modification timestamp

- `penalty_shots`: Individual penalty shot attempts and results
  - `id`: INTEGER PRIMARY KEY - Unique shot identifier
  - `contestant_id`: INTEGER - Foreign key to contestants table
  - `attempt_number`: INTEGER - Sequential attempt number for contestant
  - `result`: TEXT - Shot result (goal, miss, saved)
  - `timestamp`: DATETIME - Shot attempt timestamp
  - `operator_id`: TEXT - Booth operator identifier

- `leaderboard_cache`: Optimized leaderboard calculations
  - `contestant_id`: INTEGER PRIMARY KEY - Foreign key to contestants
  - `total_attempts`: INTEGER - Total penalty shots taken
  - `successful_shots`: INTEGER - Number of goals scored
  - `success_percentage`: REAL - Calculated success rate
  - `last_updated`: DATETIME - Cache update timestamp

### API Design
**Endpoints**:
- `POST /api/scan-badge`: Initiate badge scanning process
  - Parameters: image_data (base64), confidence_threshold (optional)
  - Response: {contestant_data, confidence_score, validation_errors}

- `POST /api/contestants`: Register new contestant
  - Parameters: {name, company, email, phone, badge_id}
  - Response: {contestant_id, registration_status, validation_errors}

- `POST /api/penalty-shots`: Record penalty shot attempt
  - Parameters: {contestant_id, result, operator_id}
  - Response: {shot_id, updated_stats, leaderboard_position}

- `GET /api/leaderboard`: Retrieve current leaderboard
  - Parameters: limit (optional), offset (optional)
  - Response: {rankings, total_contestants, last_updated}

- `GET /api/export/{format}`: Generate data export
  - Parameters: format (csv/json), date_range (optional)
  - Response: File download or {export_url, expires_at}

### Service Design
**Services**:
- `CameraService`: Manages camera initialization, image capture, preprocessing, and Apple Silicon optimization with error recovery mechanisms
- `OCRService`: Handles Tesseract integration, badge parsing, confidence scoring, and manual override workflows with multiple format support
- `ContestantService`: Implements registration, validation, duplicate detection, and data persistence with comprehensive error handling
- `ScoreService`: Manages penalty shot recording, validation, analytics calculation, and real-time leaderboard updates
- `WebSocketService`: Provides real-time communication, event broadcasting, connection management, and graceful degradation

## Security Design

### Authentication
Session-based authentication for booth operators with configurable timeout periods. Simple PIN-based access control for booth environment with optional administrator override capabilities.

### Authorization
Role-based access control distinguishing between booth operators (score entry, contestant management) and administrators (export, configuration, system management).

### Data Security
Local data encryption for contestant information using AES-256, secure badge image handling with automatic cleanup, GDPR-compliant data retention policies, and optional data anonymization features.

## Performance Design

### Scalability
Optimized for single-booth deployment with concurrent operator support. Database connection pooling, efficient query optimization, and memory management for extended event operation (8-12 hours continuous use).

### Caching Strategy
In-memory leaderboard caching with 30-second refresh cycles, badge parsing result caching to improve repeat scan performance, and static asset caching for improved UI responsiveness.

### Optimization
Apple Silicon-specific compiler optimizations, goroutine pool management for camera processing, database query optimization with proper indexing, and WebSocket connection management with automatic cleanup.

## Integration Design

### External Integrations
- **Tesseract OCR Engine**: Native Go bindings with custom configuration for badge format optimization and multi-language support
- **AVFoundation Camera**: CGO integration for native macOS camera access with Apple Silicon performance optimization

### Internal Integrations
- **Database Layer**: SQLite integration with connection pooling, transaction management, and automatic backup mechanisms
- **WebSocket Communication**: Real-time event broadcasting between services with connection state management and automatic reconnection

## Error Handling

### Error Types
- **Camera Errors**: Hardware failure detection, graceful degradation to manual entry mode, automatic retry mechanisms with exponential backoff
- **OCR Errors**: Low confidence detection, format parsing failures, fallback to manual data entry with guided input forms
- **Database Errors**: Connection failure recovery, transaction rollback mechanisms, data corruption detection and repair
- **Network Errors**: WebSocket disconnection handling, offline mode activation, data synchronization upon reconnection

### Logging Strategy
Structured JSON logging with configurable levels (DEBUG, INFO, WARN, ERROR), automatic log rotation for extended event operation, and comprehensive error context capture for troubleshooting.

## Testing Strategy

### Unit Testing
Comprehensive test coverage for all service methods with mock implementations for external dependencies. Target 85% code coverage with focus on critical path operations (badge scanning, score recording, leaderboard calculation).

### Integration Testing
End-to-end testing of badge scanning workflow, camera integration testing with mock hardware, database operation validation, and WebSocket communication testing with multiple concurrent connections.

### User Acceptance Testing
Booth operator usability testing with realistic trade show scenarios, performance testing under high-volume contestant loads, and accessibility testing for diverse operator capabilities.

## Deployment Design

### Environment Configuration
- **Development**: Local SQLite database, mock camera interface, debug logging enabled, hot-reload development server
- **Staging**: Production-like database, actual camera hardware testing, performance monitoring enabled, limited data retention
- **Production**: Optimized SQLite configuration, full camera integration, minimal logging, automatic backup scheduling, monitoring dashboard

### Deployment Process
Native macOS binary compilation with Apple Silicon optimization, automated dependency bundling, configuration file validation, database migration scripts, and deployment verification checklist.

## Monitoring and Observability

### Metrics
Badge scanning success rate, average processing time per scan, contestant registration volume, penalty shot attempt frequency, leaderboard update latency, and system resource utilization.

### Alerting
Camera hardware failure detection, OCR accuracy degradation alerts, database performance warnings, memory usage thresholds, and WebSocket connection failure notifications.

### Logging
Structured application logs with request tracing, performance metrics logging, error tracking with stack traces, and user interaction analytics for post-event analysis.

## Design Decisions

### Decision 1: SQLite vs External Database
**Context**: Need for reliable data persistence in trade show environment with potential network connectivity issues
**Decision**: SQLite embedded database for offline-first operation with optional cloud synchronization
**Consequences**: Simplified deployment and reduced external dependencies, but limited to single-node scaling and requires manual backup procedures

### Decision 2: Native Camera Integration vs Web-based
**Context**: Requirement for high-quality badge scanning with optimal performance on Apple Silicon hardware
**Decision**: Native AVFoundation integration through CGO bindings for direct camera access
**Consequences**: Superior performance and image quality control, but increased complexity and platform-specific implementation requirements

### Decision 3: Real-time Updates via WebSockets vs Polling
**Context**: Need for live leaderboard updates and real-time operator feedback during high-volume contestant periods
**Decision**: WebSocket-based real-time communication with fallback to HTTP polling for degraded connections
**Consequences**: Improved user experience with instant updates and reduced server load, but increased connection management complexity and potential scaling challenges