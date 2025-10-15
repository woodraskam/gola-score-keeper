# Goal Score Keeper

A real-time soccer penalty shot scoring application designed for trade show environments. The system captures contestant information via badge scanning and tracks penalty shot performance for booth engagement analytics.

## Features

- **Real-time Badge Scanning**: Capture contestant information using MacBook Air camera
- **Score Tracking**: Record and display penalty shot results
- **Leaderboard Management**: Maintain competitive rankings
- **Data Analytics**: Generate engagement reports for post-event analysis
- **User Experience**: Provide intuitive interface for booth operators

## Technology Stack

- **Language**: Go 1.21+
- **Database**: SQLite 3.40+
- **Web Framework**: Gin v1.9.1
- **OCR Engine**: Tesseract 5.0+ with Apple Vision Framework
- **Platform**: macOS Monterey 12.0+ (Apple Silicon optimized)

## Performance Requirements

- **Badge Scanning**: < 5 seconds (baseline requirement)
- **UI Response**: < 2 seconds for all operations
- **Memory Usage**: < 200MB baseline (Apple Silicon optimized)
- **OCR Accuracy**: 90% baseline (95%+ with Neural Engine)

## Quick Start

### Prerequisites

- macOS Monterey 12.0+ (Apple Silicon optimized)
- Go 1.21+
- Xcode Command Line Tools
- Tesseract OCR 5.0+

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/goal-score-keeper.git
cd goal-score-keeper
```

2. Install dependencies:
```bash
brew install tesseract
go mod download
```

3. Initialize the database:
```bash
go run cmd/setup/main.go
```

4. Run the application:
```bash
go run cmd/server/main.go
```

## Project Structure

```
goal-score-keeper/
├── cmd/
│   ├── server/          # Main application entry
│   └── setup/           # Database initialization
├── internal/
│   ├── camera/          # Camera interface
│   ├── ocr/            # OCR processing
│   ├── models/         # Data models
│   ├── handlers/       # HTTP handlers
│   ├── services/       # Business logic
│   ├── database/       # Database operations
│   └── websocket/      # Real-time communication
├── web/
│   ├── static/         # CSS, JS, images
│   └── templates/       # HTML templates
├── configs/            # Configuration files
├── migrations/          # Database migrations
├── tests/              # Test files
├── docs/               # Documentation
└── scripts/            # Build and deployment scripts
```

## Development

### Build

```bash
# Standard build
make build

# Apple Silicon optimized build
make build-arm64

# Run tests
make test

# Clean build artifacts
make clean
```

### Environment Variables

```bash
# Database configuration
export DB_PATH="./data/score_keeper.db"

# OCR configuration
export OCR_LANGUAGE="eng"
export OCR_CONFIDENCE_THRESHOLD="0.7"

# Camera configuration
export CAMERA_DEVICE_ID="0"
export CAMERA_RESOLUTION="1920x1080"

# WebSocket configuration
export WS_PORT="8080"
```

## API Endpoints

- `POST /api/scan-badge` - Badge scanning and OCR processing
- `GET /api/contestants/{id}` - Retrieve contestant by ID
- `POST /api/contestants` - Register new contestant
- `POST /api/penalty-shots` - Record penalty shot result
- `GET /api/leaderboard` - Retrieve current leaderboard
- `GET /api/export/{format}` - Export data (CSV/JSON)

## WebSocket Endpoints

- `ws://localhost:8080/ws/leaderboard` - Real-time leaderboard updates
- `ws://localhost:8080/ws/scoring` - Real-time scoring updates

## Database Schema

### Core Tables

- `contestants` - Contestant information and registration data
- `penalty_shots` - Individual penalty shot attempts and results
- `leaderboard_cache` - Optimized leaderboard calculations

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: See `/docs` directory
- **Issues**: Report bugs via GitHub issues
- **Email**: Contact project maintainers
- **Slack**: Join #goal-score-keeper channel for real-time support

---

**Version**: 1.0.0  
**Last Updated**: 2024-01-15  
**Compatibility**: Apple Silicon M1/M2, macOS 12.0+
