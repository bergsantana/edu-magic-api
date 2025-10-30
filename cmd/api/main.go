package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bergsantana/edu-magic-api/internal/adapters/handlers"
	"github.com/bergsantana/edu-magic-api/internal/adapters/repositories"
	"github.com/bergsantana/edu-magic-api/internal/core/services"
	"github.com/bergsantana/edu-magic-api/internal/infrastructure/config"
	"github.com/bergsantana/edu-magic-api/internal/infrastructure/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// Set up custom logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	log.Println("🚀 Starting EDU Magic API...")

	// Load configuration
	cfg := config.LoadConfig()
	log.Printf("✅ Configuration loaded - Server will run on port %s", cfg.ServerPort)

	// Initialize database connection
	db := database.Connect()
	log.Println("✅ Database connection established")

	// Set up Gin router
	router := gin.New() // Use gin.New() instead of gin.Default() to avoid double middleware

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://edu-magic.vercel.app/", "https://*.vercel.app/"}, // Add your frontend URLs
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Enable custom detailed logging
	router.Use(customLogger())
	router.Use(gin.Recovery())

	// Services
	authService, activityService := initRepositories(db, cfg.JWTSecret)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(*authService)
	activityHandler := handlers.NewActivityHandler(*activityService)

	// Define routes
	router.POST("/signup", authHandler.Signup)
	router.POST("/login", authHandler.Login)
	router.POST("/activities", activityHandler.CreateActivity)
	router.GET("/activities/:userId", activityHandler.GetActivitiesByUser)

	// Start the server
	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}

func initRepositories(db *gorm.DB, jwtSecret string) (*services.AuthService, *services.ActivityService) {
	// Initialize repositories here
	userRepo := repositories.NewUserRepository(db)
	activityRepo := repositories.NewActivityRepository(db)

	return initServices(userRepo, activityRepo, jwtSecret)
	// Optionally, you can return them or use them to initialize services
}

func initServices(userRepo *repositories.UserRepository, activityRepo *repositories.ActivityRepository, jwtSecret string) (*services.AuthService, *services.ActivityService) {
	// Initialize services here
	authService := services.NewAuthService(*userRepo, jwtSecret)
	activityService := services.NewActivityService(*activityRepo)
	return authService, activityService
}

// customLogger returns a gin.HandlerFunc (middleware) that logs requests in a custom format.
func customLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - 📍 [%s] \"%s %s %s\" %d %s \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
