package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/verssache/chatgpt-creator/docs"
	"github.com/verssache/chatgpt-creator/internal/api/handlers"
	"github.com/verssache/chatgpt-creator/internal/api/middleware"
	"github.com/verssache/chatgpt-creator/internal/api/routes"
	"github.com/verssache/chatgpt-creator/internal/config"
)

// @title ChatGPT Creator API
// @version 1.0
// @description API for ChatGPT account registration bot with batch processing, email domain management, and real-time progress tracking.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/verssache/chatgpt-creator
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api
// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Set Gin mode BEFORE any Gin operations
	gin.SetMode(gin.ReleaseMode)

	envCfg, err := config.LoadEnv()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	logger := config.NewLogger(&envCfg.Logging)

	db, err := config.NewDatabase(&envCfg.DB, logger)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.RunMigrations("migrations"); err != nil {
		logger.Warnf("Migration warning: %v", err)
	}

	router := gin.New()
	router.Use(middleware.Logger(logger))
	router.Use(middleware.Recovery(logger))
	router.Use(middleware.CORS(envCfg.Frontend.CORSOrigins))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"version": "1.0.0",
		})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		routes.SetupRoutes(api, db.DB, logger, envCfg)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", envCfg.App.Port),
		Handler:      router,
		ReadTimeout:  envCfg.API.ReadTimeout,
		WriteTimeout: envCfg.API.WriteTimeout,
	}

	go func() {
		logger.Infof("Starting server on port %d", envCfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), envCfg.API.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Server forced to shutdown: %v", err)
	}

	logger.Info("Server exited")
}

func init() {
	handlers.Init()
}
