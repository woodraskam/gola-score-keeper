# Badge Information Extraction - README Documentation

## Overview
Comprehensive README documentation for Badge Information Extraction, providing essential information for developers, users, and stakeholders.

## Project Description
Badge Information Extraction is a core processing component that provides OCR and image processing functionality to automatically extract contestant information from captured badge images, including name, company, and other relevant details for the Gola Score Keeper application.

## Features
- **OCR Processing**: Advanced optical character recognition with 90% baseline accuracy (95%+ with Neural Engine)
- **Multi-format Support**: Handles various badge formats and layouts commonly used at trade shows
- **Real-time Processing**: Fast image processing with results delivered in under 5 seconds
- **Data Validation**: Intelligent validation and error correction for extracted information
- **Integration Points**: Seamless integration with camera interface, contestant database, and web interface components

## Getting Started

### Prerequisites
- macOS Monterey 12.0+ (Apple Silicon optimized)
- Go 1.21 or higher
- Tesseract OCR 5.0+
- Xcode Command Line Tools
- Camera access permissions

### Installation
1. Install Tesseract OCR engine via Homebrew:
   ```bash
   brew install tesseract
   ```
2. Install Go dependencies:
   ```bash
   go mod download
   ```
3. Configure camera permissions in macOS System Preferences

### Quick Start
```go
package main

import (
    "github.com/yourorg/gola-score-keeper/internal/ocr"
    "log"
)

func main() {
    processor := ocr.NewBadgeProcessor()
    result, err := processor.ExtractBadgeInfo("path/to/badge/image.jpg")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Extracted: Name=%s, Company=%s", result.Name, result.Company)
}
```

## Usage

### Basic Usage
1. Initialize the badge processor with default configuration
2. Pass captured image file path or image data to the extraction method
3. Receive structured contestant information with confidence scores
4. Validate and store the extracted data in the contestant database

### Advanced Configuration
```go
config := ocr.Config{
    Language:        "eng",
    MinConfidence:   0.7,
    PreprocessImage: true,
    CustomPatterns:  []string{"company", "name", "title"},
    Timeout:         5 * time.Second,
}

processor := ocr.NewBadgeProcessorWithConfig(config)
```

### API Reference
- `ExtractBadgeInfo(imagePath string) (*BadgeInfo, error)`: Extract information from image file
- `ExtractFromBytes(imageData []byte) (*BadgeInfo, error)`: Extract from image byte data
- `ValidateExtraction(info *BadgeInfo) ValidationResult`: Validate extracted information
- `SetConfidenceThreshold(threshold float64)`: Adjust accuracy requirements

## Configuration

### Environment Variables
- `OCR_LANGUAGE`: Language model for OCR processing (default: "eng")
- `OCR_CONFIDENCE_THRESHOLD`: Minimum confidence score for acceptance (default: 0.7)
- `IMAGE_PREPROCESSING`: Enable image enhancement before OCR (default: true)
- `OCR_TIMEOUT`: Maximum processing time in seconds (default: 5)

### Settings
- **Language Models**: Configure primary and fallback language detection
- **Image Enhancement**: Adjust contrast, brightness, and noise reduction parameters
- **Pattern Recognition**: Define custom regex patterns for specific badge formats
- **Validation Rules**: Set field validation criteria and format requirements

## Examples

### Example 1: Basic Badge Processing
```go
processor := ocr.NewBadgeProcessor()
badgeInfo, err := processor.ExtractBadgeInfo("captured_badge.jpg")
if err != nil {
    log.Printf("OCR processing failed: %v", err)
    return
}

fmt.Printf("Name: %s\n", badgeInfo.Name)
fmt.Printf("Company: %s\n", badgeInfo.Company)
fmt.Printf("Title: %s\n", badgeInfo.Title)
fmt.Printf("Confidence: %.2f\n", badgeInfo.Confidence)
```

### Example 2: Advanced Processing with Validation
```go
processor := ocr.NewBadgeProcessor()
processor.SetConfidenceThreshold(0.85)

badgeInfo, err := processor.ExtractBadgeInfo("badge_image.jpg")
if err != nil {
    return handleOCRError(err)
}

validation := processor.ValidateExtraction(badgeInfo)
if !validation.IsValid {
    return handleValidationErrors(validation.Errors)
}

// Store in database
contestant := models.Contestant{
    Name:    badgeInfo.Name,
    Company: badgeInfo.Company,
    Title:   badgeInfo.Title,
}
database.SaveContestant(contestant)
```

## Troubleshooting

### Common Issues
1. **Issue**: OCR accuracy below 90% threshold (85% minimum for manual intervention)
   **Solution**: 
   - Ensure proper lighting during image capture
   - Check camera focus and image quality
   - Enable image preprocessing options
   - Verify badge is fully visible and unobstructed
   - Consider Neural Engine acceleration on Apple Silicon

2. **Issue**: Processing timeout errors
   **Solution**: 
   - Increase OCR_TIMEOUT environment variable
   - Optimize image size before processing
   - Check system resource availability
   - Consider reducing image resolution for faster processing

3. **Issue**: Incorrect company name extraction
   **Solution**: 
   - Add custom regex patterns for known company formats
   - Adjust confidence threshold for company field
   - Implement manual correction interface for operators

### Debugging
- Enable debug logging with `OCR_DEBUG=true` environment variable
- Use `processor.GetLastProcessingStats()` to analyze performance metrics
- Check intermediate image processing steps with `SaveDebugImages=true`
- Monitor memory usage during batch processing operations

## Contributing
1. Fork the repository and create a feature branch
2. Follow Go coding standards and include comprehensive tests
3. Ensure OCR accuracy meets or exceeds 90% requirement
4. Add documentation for new configuration options
5. Submit pull request with detailed description of changes
6. All contributions must pass automated testing and code review

## License
This project is licensed under the MIT License. See LICENSE file for details.

## Support
- **Technical Issues**: Create GitHub issue with detailed reproduction steps
- **Feature Requests**: Submit enhancement proposals through GitHub issues
- **Documentation**: Contribute improvements via pull requests
- **Emergency Support**: Contact development team during event operations
- **Performance Issues**: Include system specifications and processing metrics in reports