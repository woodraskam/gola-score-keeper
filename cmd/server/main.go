package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

func main() {
	log.Printf("Starting Goal Score Keeper v%s (build: %s, commit: %s)", Version, BuildTime, GitCommit)

	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"version":   Version,
			"buildTime": BuildTime,
			"gitCommit": GitCommit,
		})
	})

	// API routes (to be implemented)
	api := router.Group("/api")
	{
		// Contestant routes
		api.GET("/contestants", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Contestants endpoint - to be implemented"})
		})

		// Penalty shot routes
		api.POST("/penalty-shots", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Penalty shots endpoint - to be implemented"})
		})

		// Leaderboard routes
		api.GET("/leaderboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Leaderboard endpoint - to be implemented"})
		})

		// Badge scanning routes
		api.POST("/scan-badge", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Badge scanning endpoint - to be implemented"})
		})
	}

	// WebSocket routes (to be implemented)
	router.GET("/ws/leaderboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "WebSocket leaderboard endpoint - to be implemented"})
	})

	router.GET("/ws/scoring", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "WebSocket scoring endpoint - to be implemented"})
	})

	// Static file serving
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")

	// Main dashboard route
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "dashboard.html", gin.H{
			"title": "Goal Score Keeper",
		})
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}
