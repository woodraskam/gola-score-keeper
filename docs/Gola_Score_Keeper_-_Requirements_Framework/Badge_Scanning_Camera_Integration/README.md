# Badge Scanning Camera Integration - README Documentation

## Overview
Comprehensive README documentation for Badge Scanning Camera Integration, providing essential information for developers, users, and stakeholders.

## Project Description
Badge Scanning Camera Integration is a core functional component that implements camera functionality to capture contestant badges using the MacBook Air's built-in camera with real-time preview and image capture capabilities for the Gola Score Keeper application.

## Features
- **Core Functionality**: Real-time camera access and badge image capture using MacBook Air's built-in camera
- **Key Capabilities**: 
  - Live camera preview with optimized frame rates
  - High-quality image capture for badge scanning
  - Local image storage and management
  - Apple Silicon performance optimization
  - Error handling and graceful degradation
- **Integration Points**: Seamlessly integrates with OCR Processing Engine, Contestant Management System, and Web Interface Layer

## Getting Started

### Prerequisites
- macOS Monterey 12.0+ with Apple Silicon support
- Go 1.21+ with CGO enabled
- Xcode Command Line Tools installed
- AVFoundation framework access permissions
- Camera hardware permissions granted

### Installation
1. Install system dependencies via Homebrew:
   ```bash
   brew install pkg-config opencv
   ```
2. Clone the repository and navigate to camera module:
   ```bash
   git clone https://github.com/yourorg/gola-score-keeper.git
   cd gola-score-keeper/internal/camera
   ```
3. Install Go dependencies:
   ```bash
   go mod download
   go get -u github.com/blackjack/webcam
   ```

### Quick Start
```go
package main

import (
    "github.com/yourorg/gola-score-keeper/internal/camera"
    "log"
)

func main() {
    cam, err := camera.NewBadgeScanner()
    if err != nil {
        log.Fatal("Failed to initialize camera:", err)
    }
    defer cam.Close()
    
    // Start preview
    if err := cam.StartPreview(); err != nil {
        log.Fatal("Failed to start preview:", err)
    }
    
    // Capture badge image
    image, err := cam.CaptureImage()
    if err != nil {
        log.Fatal("Failed to capture image:", err)
    }
    
    // Save locally
    cam.SaveImage(image, "badge_001.jpg")
}
```

## Usage

### Basic Usage
1. Initialize the camera scanner with default settings
2. Request camera permissions if not already granted
3. Start the live preview stream for operator feedback
4. Position badge within the camera frame
5. Trigger image capture when badge is properly aligned
6. Save captured image to local storage for OCR processing

### Advanced Configuration
```go
config := camera.Config{
    Resolution:    camera.Resolution1080p,
    FrameRate:     30,
    Quality:       camera.QualityHigh,
    AutoFocus:     true,
    StoragePath:   "./captured_badges/",
    PreviewWindow: true,
    Timeout:       30 * time.Second,
}

scanner := camera.NewBadgeScannerWithConfig(config)
```

### API Reference

#### Core Methods
- `NewBadgeScanner() (*BadgeScanner, error)` - Initialize camera with default settings
- `StartPreview() error` - Begin real-time camera preview
- `StopPreview() error` - Stop camera preview stream
- `CaptureImage() (*Image, error)` - Capture high-quality badge image
- `SaveImage(image *Image, filename string) error` - Save image to local storage
- `Close() error` - Release camera resources

#### Configuration Methods
- `SetResolution(res Resolution) error` - Configure capture resolution
- `SetQuality(quality Quality) error` - Set image quality settings
- `SetStoragePath(path string) error` - Configure local storage location

## Configuration

### Environment Variables
- `CAMERA_DEVICE_ID`: Camera device identifier (default: 0 for built-in camera)
- `BADGE_STORAGE_PATH`: Local storage path for captured images (default: ./badges/)
- `CAMERA_TIMEOUT`: Maximum time for camera operations in seconds (default: 30)
- `PREVIEW_ENABLED`: Enable/disable live preview (default: true)

### Settings
```yaml
camera:
  resolution: "1920x1080"
  frame_rate: 30
  quality: "high"
  auto_focus: true
  storage_path: "./captured_badges/"
  preview_window: true
  capture_timeout: 30
  max_file_size: "10MB"
  supported_formats: ["jpg", "png"]
```

## Examples

### Example 1: Basic Badge Capture
```go
func captureBadge() error {
    scanner, err := camera.NewBadgeScanner()
    if err != nil {
        return fmt.Errorf("camera initialization failed: %w", err)
    }
    defer scanner.Close()
    
    // Start preview for operator
    if err := scanner.StartPreview(); err != nil {
        return fmt.Errorf("preview start failed: %w", err)
    }
    
    // Wait for operator signal or auto-trigger
    time.Sleep(2 * time.Second)
    
    // Capture badge image
    image, err := scanner.CaptureImage()
    if err != nil {
        return fmt.Errorf("image capture failed: %w", err)
    }
    
    // Save with timestamp
    filename := fmt.Sprintf("badge_%d.jpg", time.Now().Unix())
    return scanner.SaveImage(image, filename)
}
```

### Example 2: Continuous Scanning Mode
```go
func continuousScanning(ctx context.Context) error {
    scanner, err := camera.NewBadgeScanner()
    if err != nil {
        return err
    }
    defer scanner.Close()
    
    scanner.StartPreview()
    defer scanner.StopPreview()
    
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case trigger := <-scanner.TriggerChannel():
            if trigger.Ready {
                image, err := scanner.CaptureImage()
                if err != nil {
                    log.Printf("Capture failed: %v", err)
                    continue
                }
                
                go processImage(image) // Async processing
            }
        }
    }
}
```

## Troubleshooting

### Common Issues
1. **Issue**: Camera not detected or permission denied
   **Solution**: 
   - Check System Preferences > Security & Privacy > Camera
   - Grant camera access to the application
   - Verify camera is not in use by another application

2. **Issue**: Poor image quality or blurry captures
   **Solution**:
   - Ensure adequate lighting conditions
   - Check auto-focus settings are enabled
   - Verify camera lens is clean
   - Adjust capture resolution settings

3. **Issue**: Preview window not displaying
   **Solution**:
   - Verify display permissions in macOS settings
   - Check if preview_enabled configuration is true
   - Restart application with administrator privileges

4. **Issue**: High CPU usage during preview
   **Solution**:
   - Reduce frame rate in configuration (15-20 fps)
   - Lower preview resolution
   - Enable hardware acceleration if available

### Debugging
- Enable debug logging: `export CAMERA_DEBUG=true`
- Check camera device list: `camera.ListAvailableDevices()`
- Monitor memory usage during continuous operation
- Use Activity Monitor to check camera process resource usage
- Verify file permissions for storage directory

## Contributing
1. Fork the repository and create a feature branch
2. Follow Go coding standards and include comprehensive tests
3. Test camera functionality on multiple macOS versions
4. Update documentation for any API changes
5. Submit pull request with detailed description of changes
6. Ensure all CI/CD checks pass before requesting review

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Support
- **Technical Issues**: Create GitHub issue with system information and error logs
- **Feature Requests**: Submit enhancement requests through GitHub discussions
- **Documentation**: Contribute improvements via pull requests
- **Emergency Support**: Contact development team at dev-support@company.com
- **Performance Issues**: Include system specifications and performance metrics in reports