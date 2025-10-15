# Penalty Shot Scoring Interface - Design Specification

## Design Overview
The Penalty Shot Scoring Interface is designed as a responsive, touch-optimized web component that provides booth operators with an intuitive interface for recording penalty shot results in real-time. The design follows a component-based architecture that integrates seamlessly with the existing Gola Score Keeper system, emphasizing simplicity, speed, and reliability in high-pressure booth environments. The interface prioritizes large, accessible controls with immediate visual feedback and real-time score synchronization across all connected displays.

## Architecture

### System Architecture
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Web Browser   │    │  Scoring API    │    │   Database      │
│                 │    │                 │    │                 │
│ ┌─────────────┐ │    │ ┌─────────────┐ │    │ ┌─────────────┐ │
│ │ Scoring UI  │◄┼────┤ │ Shot Record │ ├────┤ │ SQLite      │ │
│ │ Component   │ │    │ │ Handler     │ │    │ │ Storage     │ │
│ └─────────────┘ │    │ └─────────────┘ │    │ └─────────────┘ │
│                 │    │                 │    │                 │
│ ┌─────────────┐ │    │ ┌─────────────┐ │    │ ┌─────────────┐ │
│ │ WebSocket   │◄┼────┤ │ Real-time   │ ├────┤ │ Score       │ │
│ │ Client      │ │    │ │ Broadcast   │ │    │ │ Engine      │ │
│ └─────────────┘ │    │ └─────────────┘ │    │ └─────────────┘ │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### Component Architecture
- **ScoringInterface**: Main UI component managing state, user interactions, and visual feedback with touch-optimized controls
- **ShotRecordingService**: Backend service handling shot validation, database persistence, and business logic enforcement
- **RealTimeUpdateEngine**: WebSocket-based service broadcasting score changes to all connected displays and leaderboards
- **ContestantDisplayComponent**: Subsidiary component showing current contestant information with automatic updates
- **ValidationLayer**: Input validation service ensuring data integrity and preventing duplicate or invalid entries

### Data Flow
```
User Button Press → JavaScript Event Handler → API Request Validation → 
Database Transaction → Score Calculation → WebSocket Broadcast → 
UI Update Confirmation → Leaderboard Refresh → Analytics Logging
```

## User Interface Design

### Wireframes
**Main Scoring Interface Layout**:
- Header: Current contestant name and company (prominent display)
- Center: Two large buttons - "GOAL" (green, 120px height) and "MISS" (red, 120px height)
- Stats Panel: Current score, shot count, goal percentage (real-time updates)
- Footer: Undo button (conditional), status indicators, timestamp

**Responsive Breakpoints**:
- Desktop (1024px+): Side-by-side layout with contestant info panel
- Tablet (768-1023px): Stacked layout optimized for touch interaction
- Mobile (320-767px): Simplified single-column layout with larger buttons

### User Experience Flow
1. **Contestant Activation** → Interface displays contestant information and enables scoring buttons
2. **Shot Recording** → User taps GOAL/MISS → Immediate visual feedback → Database update → Score refresh
3. **Error Handling** → Invalid action → Clear error message → Guidance for resolution → Return to ready state
4. **Session Management** → Automatic contestant detection → Seamless transitions → Session timeout handling

### Design Principles
- **Touch-First Design**: Minimum 44px touch targets with generous spacing for booth environment
- **High Contrast Visuals**: Colors and typography optimized for various lighting conditions
- **Immediate Feedback**: Visual and optional audio confirmation for all user actions
- **Error Prevention**: Disabled states and validation prevent common operator mistakes
- **Accessibility Compliance**: WCAG 2.1 AA standards with keyboard navigation and screen reader support

## Technical Design

### Database Design
**Tables**:
- `shots`: Shot recording and tracking
  - `id`: INTEGER PRIMARY KEY - Unique shot identifier
  - `contestant_id`: INTEGER NOT NULL - Foreign key to contestants table
  - `shot_result`: TEXT NOT NULL - 'goal' or 'miss' enumeration
  - `timestamp`: DATETIME DEFAULT CURRENT_TIMESTAMP - Shot recording time
  - `session_id`: TEXT NOT NULL - Scoring session identifier
  - `operator_id`: TEXT - Booth operator identifier for audit trail
  - `deleted_at`: DATETIME NULL - Soft delete for undo functionality

- `scoring_sessions`: Session management and tracking
  - `id`: TEXT PRIMARY KEY - Session UUID
  - `contestant_id`: INTEGER NOT NULL - Active contestant reference
  - `start_time`: DATETIME NOT NULL - Session initiation timestamp
  - `end_time`: DATETIME NULL - Session completion timestamp
  - `total_shots`: INTEGER DEFAULT 0 - Calculated shot count
  - `goals_made`: INTEGER DEFAULT 0 - Calculated goal count
  - `status`: TEXT DEFAULT 'active' - Session status enumeration

### API Design
**Endpoints**:
- `POST /api/shots/record`: Record penalty shot result
  - Parameters: `contestant_id`, `shot_result`, `session_id`
  - Response: `{"success": true, "shot_id": 123, "updated_score": {...}}`
  
- `GET /api/contestant/current`: Retrieve active contestant information
  - Parameters: None (session-based)
  - Response: `{"contestant": {...}, "current_stats": {...}, "session_id": "..."}`
  
- `PUT /api/shots/{id}/undo`: Undo recent shot within time window
  - Parameters: `shot_id` in URL path
  - Response: `{"success": true, "reverted_score": {...}}`

- `WebSocket /ws/scoring`: Real-time score updates and notifications
  - Events: `shot_recorded`, `score_updated`, `contestant_changed`, `session_ended`

### Service Design
**Services**:
- `ScoringService`: Core business logic for shot recording, validation, and score calculation with concurrent safety
- `ContestantService`: Active contestant management, session handling, and state transitions
- `ValidationService`: Input validation, business rule enforcement, and duplicate prevention logic
- `BroadcastService`: WebSocket connection management and real-time update distribution
- `AuditService`: Logging and tracking service for operator actions and system events

## Security Design

### Authentication
Session-based authentication using secure HTTP cookies with booth operator credentials. Sessions expire after 8 hours or manual logout. CSRF protection implemented for all state-changing operations.

### Authorization
Role-based access control with booth operator permissions for shot recording and basic administrative functions. Administrative functions (shot deletion, session management) require elevated privileges.

### Data Security
Input sanitization and validation for all user inputs. SQL injection prevention through parameterized queries. XSS protection through output encoding. Rate limiting on API endpoints to prevent abuse.

## Performance Design

### Scalability
Horizontal scaling through stateless API design with session storage in database. Connection pooling for database operations. Efficient WebSocket connection management with automatic cleanup.

### Caching Strategy
In-memory caching of active contestant data and session information. Browser-side caching of static assets with appropriate cache headers. Real-time data bypasses cache for accuracy.

### Optimization
Database query optimization with proper indexing on frequently accessed columns. Lazy loading of historical data. Debounced user inputs to prevent excessive API calls. Efficient DOM updates using virtual DOM concepts.

## Integration Design

### External Integrations
- **Camera/Badge Scanner**: Receives contestant activation events to update interface state
- **Leaderboard Display**: Broadcasts score updates through WebSocket connections for real-time leaderboard updates
- **Export Module**: Provides shot data for post-event reporting and analytics generation

### Internal Integrations
- **Contestant Management System**: Retrieves active contestant information and validates contestant status
- **Score Tracking Engine**: Integrates with core scoring calculations and historical data management
- **Web Interface Framework**: Embedded within main Gola Score Keeper web application structure

## Error Handling

### Error Types
- **Validation Errors**: Invalid shot data or missing contestant information handled with clear user messaging
- **Network Errors**: Connection failures handled with offline queuing and retry mechanisms
- **Database Errors**: Transaction failures handled with rollback and user notification
- **Session Errors**: Expired or invalid sessions handled with automatic re-authentication prompts

### Logging Strategy
Structured logging with JSON format including timestamp, user context, action performed, and outcome. Error logs include stack traces and request context. Performance metrics logged for monitoring and optimization.

## Testing Strategy

### Unit Testing
Comprehensive unit tests for all service methods with mock dependencies. JavaScript unit tests for UI components using Jest framework. Database operation testing with isolated test database. Target 90% code coverage.

### Integration Testing
End-to-end API testing with real database connections. WebSocket connection testing with multiple clients. Cross-browser compatibility testing on target browsers. Touch interface testing on actual tablet devices.

### User Acceptance Testing
Booth operator testing in simulated event environment. Usability testing with time-pressure scenarios. Accessibility testing with keyboard navigation and screen readers. Performance testing under expected load conditions.

## Deployment Design

### Environment Configuration
- **Development**: Local SQLite database, debug logging enabled, hot reload for rapid iteration
- **Staging**: Production-like environment with test data, performance monitoring, integration testing
- **Production**: Optimized builds, production database, error tracking, performance monitoring, backup systems

### Deployment Process
1. Automated testing pipeline with unit and integration tests
2. Build optimization and asset compilation
3. Database migration execution with rollback capability
4. Blue-green deployment with health checks
5. Performance monitoring and alerting activation
6. Rollback procedures for critical failures

## Monitoring and Observability

### Metrics
- **Performance Metrics**: Response times, throughput, error rates, WebSocket connection health
- **Business Metrics**: Shot recording rate, contestant engagement, operator efficiency, system uptime
- **System Metrics**: Database performance, memory usage, CPU utilization, network latency

### Alerting
- **Critical Alerts**: System downtime, database connectivity failures, high error rates
- **Warning Alerts**: Performance degradation, unusual usage patterns, resource constraints
- **Information Alerts**: Deployment completions, scheduled maintenance, usage statistics

### Logging
Centralized logging with log aggregation and searching capabilities. Log retention policy of 90 days for operational logs, 1 year for audit logs. Real-time log monitoring with automated alert triggers.

## Design Decisions

### Decision 1: WebSocket vs Server-Sent Events for Real-time Updates
**Context**: Need for real-time score updates across multiple connected displays and leaderboards
**Decision**: WebSocket implementation for bidirectional communication and lower latency
**Consequences**: More complex connection management but superior real-time performance and user experience

### Decision 2: Touch-Optimized Interface vs Traditional Desktop Interface
**Context**: Primary usage in booth environment with potential touch screen devices and high-stress operation
**Decision**: Touch-first design with large buttons and generous spacing optimized for tablet and touch screen usage
**Consequences**: Slightly less information density but significantly improved usability and reduced operator errors

### Decision 3: Embedded Component vs Standalone Application
**Context**: Integration with existing Gola Score Keeper system architecture and workflow
**Decision**: Embedded component within main application framework sharing authentication and data models
**Consequences**: Tighter integration and consistency but increased complexity in main application codebase