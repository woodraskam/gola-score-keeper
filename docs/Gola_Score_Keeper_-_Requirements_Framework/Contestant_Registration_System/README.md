# Contestant Registration System - README Documentation

## Overview
Comprehensive README documentation for Contestant Registration System, providing essential information for developers, users, and stakeholders.

## Project Description
Contestant Registration System is a core functional component that enables automated registration of trade show participants through badge scanning technology, assigns unique tracking identifiers, and prevents duplicate registrations within the Gola Score Keeper application.

## Features
- **Core Functionality**: Automated badge scanning and contestant data extraction using OCR technology
- **Key Capabilities**: 
  - Real-time badge information parsing (name, company, contact details)
  - Duplicate detection and prevention algorithms
  - Unique identifier generation and assignment
  - Data validation and error handling
  - Local database persistence with SQLite
- **Integration Points**: Integrates with Camera Interface Module, OCR Processing Engine, Score Tracking Engine, and Web Interface Layer

## Getting Started

### Prerequisites
- macOS Monterey 12.0+ (Apple Silicon optimized)
- Go 1.21 or higher
- Tesseract OCR 5.0+
- SQLite 3.40+
- Xcode Command Line Tools
- Active camera device (MacBook Air built-in camera)

### Installation
1. Clone the Gola Score Keeper repository and navigate to the project directory
2. Install OCR dependencies: `brew install tesseract`
3. Initialize the contestant database: `go run cmd/setup/main.go --module=contestants`
4. Verify camera permissions in macOS System Preferences

### Quick Start
```go
// Initialize the registration system
registrationSystem := contestant.NewRegistrationSystem()

// Register a new contestant from badge scan
result, err := registrationSystem.RegisterFromBadge(badgeImageData)
if err != nil {
    log.Printf("Registration failed: %v", err)
    return
}

fmt.Printf("Contestant registered: ID=%s, Name=%s", result.ID, result.Name)
```

## Usage

### Basic Usage
1. **Start the registration interface**: Launch the web dashboard at `http://localhost:8080/registration`
2. **Position badge for scanning**: Place contestant badge within camera frame
3. **Initiate scan**: Click "Scan Badge" button or use automatic detection
4. **Review extracted data**: Verify parsed information (name, company, email)
5. **Complete registration**: System assigns unique ID and saves to database
6. **Handle duplicates**: System alerts if contestant already registered

### Advanced Configuration
```go
// Configure OCR processing parameters
ocrConfig := ocr.Config{
    Language:     "eng",
    PSM:          6, // Uniform block of text
    Confidence:   75, // Minimum confidence threshold
    Preprocessing: true, // Enable image enhancement
}

// Configure duplicate detection sensitivity
duplicateConfig := contestant.DuplicateConfig{
    NameSimilarity:    0.85, // 85% similarity threshold
    EmailExactMatch:   true, // Require exact email match
    CompanyFuzzyMatch: 0.80, // 80% company name similarity
}

// Initialize with custom configuration
registrationSystem := contestant.NewRegistrationSystem(
    contestant.WithOCRConfig(ocrConfig),
    contestant.WithDuplicateConfig(duplicateConfig),
)
```

### API Reference
- `RegisterFromBadge(imageData []byte) (*Contestant, error)`: Register contestant from badge image
- `RegisterManual(contestantData ContestantData) (*Contestant, error)`: Manual registration fallback
- `FindByID(id string) (*Contestant, error)`: Retrieve contestant by unique ID
- `CheckDuplicate(contestantData ContestantData) (bool, *Contestant, error)`: Check for existing registration
- `GetRegistrationStats() (*RegistrationStats, error)`: Get registration statistics
- `ExportContestants(format string) ([]byte, error)`: Export contestant data

## Configuration

### Environment Variables
- `CAMERA_DEVICE_ID`: Camera device identifier (default: 0)
- `OCR_LANGUAGE`: Tesseract language pack (default: "eng")
- `DB_PATH`: SQLite database file path (default: "./data/contestants.db")
- `DUPLICATE_THRESHOLD`: Similarity threshold for duplicate detection (default: 0.85)
- `REGISTRATION_TIMEOUT`: Maximum time for registration process (default: 30s)

### Settings
```yaml
registration:
  camera:
    resolution: "1920x1080"
    fps: 30
    auto_focus: true
  ocr:
    confidence_threshold: 75
    preprocessing: true
    noise_reduction: true
  validation:
    required_fields: ["name", "company"]
    email_validation: true
    phone_validation: false
  storage:
    batch_size: 100
    backup_interval: "1h"
```

## Examples

### Example 1: Standard Badge Registration
```go
// Initialize camera and registration system
camera := camera.NewInterface()
registration := contestant.NewRegistrationSystem()

// Capture badge image
imageData, err := camera.CaptureImage()
if err != nil {
    log.Fatal("Camera capture failed:", err)
}

// Process registration
contestant, err := registration.RegisterFromBadge(imageData)
if err != nil {
    if errors.Is(err, contestant.ErrDuplicateFound) {
        fmt.Println("Contestant already registered")
        return
    }
    log.Fatal("Registration failed:", err)
}

fmt.Printf("Successfully registered: %+v\n", contestant)
```

### Example 2: Manual Registration with Validation
```go
// Manual registration for damaged/unreadable badges
contestantData := contestant.ContestantData{
    Name:    "John Smith",
    Company: "Tech Corp",
    Email:   "john.smith@techcorp.com",
    Phone:   "+1-555-0123",
}

// Validate data before registration
if err := registration.ValidateData(contestantData); err != nil {
    fmt.Printf("Validation failed: %v\n", err)
    return
}

// Check for duplicates
isDuplicate, existing, err := registration.CheckDuplicate(contestantData)
if err != nil {
    log.Fatal("Duplicate check failed:", err)
}

if isDuplicate {
    fmt.Printf("Duplicate found: %s (ID: %s)\n", existing.Name, existing.ID)
    return
}

// Register manually
contestant, err := registration.RegisterManual(contestantData)
if err != nil {
    log.Fatal("Manual registration failed:", err)
}

fmt.Printf("Manual registration successful: %s\n", contestant.ID)
```

## Troubleshooting

### Common Issues
1. **Issue**: OCR extraction accuracy below 80%
   **Solution**: Improve lighting conditions, clean camera lens, adjust badge positioning, enable image preprocessing in configuration

2. **Issue**: Camera not detected or permission denied
   **Solution**: Check macOS privacy settings for camera access, verify camera device ID in environment variables, restart application with proper permissions

3. **Issue**: Duplicate detection false positives
   **Solution**: Adjust similarity thresholds in duplicate configuration, review name normalization rules, check for special characters in company names

4. **Issue**: Database connection errors
   **Solution**: Verify SQLite database file permissions, check available disk space, ensure database initialization completed successfully

### Debugging
- Enable debug logging: `export LOG_LEVEL=debug`
- OCR confidence analysis: Use `--debug-ocr` flag to output confidence scores
- Camera diagnostics: Run `go run cmd/camera-test/main.go` for camera functionality test
- Database integrity check: `sqlite3 contestants.db ".schema"` to verify table structure
- Performance profiling: Enable pprof endpoint at `/debug/pprof/` for performance analysis

## Contributing
1. Fork the repository and create feature branch: `git checkout -b feature/registration-enhancement`
2. Follow Go coding standards and include unit tests for new functionality
3. Test with various badge formats and lighting conditions
4. Update documentation for any API changes
5. Submit pull request with detailed description of changes
6. Ensure all CI/CD checks pass including security scans

## License
This project is licensed under the MIT License - see the LICENSE file for details. OCR functionality uses Tesseract OCR under Apache 2.0 License.

## Support
- **Technical Issues**: Create GitHub issue with detailed reproduction steps
- **Feature Requests**: Submit enhancement proposals through GitHub discussions
- **Emergency Support**: Contact development team at dev-team@company.com
- **Documentation**: Visit project wiki at https://github.com/yourorg/gola-score-keeper/wiki
- **Community**: Join Slack channel #gola-score-keeper for real-time assistance