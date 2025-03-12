package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-server/internal/application/usecase"
	"web-server/internal/infrastructure/config"
	"web-server/internal/infrastructure/middleware"
	"web-server/internal/infrastructure/repository"
	"web-server/internal/interface/handler"

	_ "web-server/docs" // This is required for swagger docs

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Web Server API
// @version         1.0
// @description     A REST API server in Go using Gin framework
// @host            localhost:8080
// @BasePath        /api
// @schemes         http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter the token with the `Bearer ` prefix, e.g. "Bearer abcde12345"

type Server struct {
	router *gin.Engine
	logger *logrus.Logger
}

func NewServer() *Server {
	// Initialize logger
	logger := middleware.InitLogger()

	// Initialize gin in release mode
	gin.SetMode(gin.ReleaseMode)

	server := &Server{
		router: gin.New(), // Use gin.New() instead of Default() to avoid default middleware
		logger: logger,
	}

	// Use our custom logger middleware
	server.router.Use(middleware.LoggerMiddleware(logger))

	// Use recovery middleware
	server.router.Use(gin.Recovery())

	// Swagger documentation endpoint
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize database and repositories
	prismaClient := config.GetPrismaClient()
	userRepo := repository.NewPrismaUserRepository(prismaClient)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUseCase)

	// Apply global middleware
	server.router.Use(middleware.EncryptionMiddleware())

	// Apply rate limiter middleware - 100 requests per minute
	server.router.Use(middleware.RateLimiterMiddleware(100, 1.67))

	// Setup routes
	api := server.router.Group("/api")
	{
		// Public routes
		public := api.Group("/public")
		{
			public.POST("/users/register", userHandler.CreateUser)
			public.POST("/users/login", userHandler.LoginUser)
			public.POST("/users/refresh", userHandler.RefreshToken) // Add refresh token endpoint
		}

		// Private routes (require authentication)
		private := api.Group("/private")
		private.Use(middleware.AuthMiddleware())
		{
			// User routes (require authentication)
			users := private.Group("/users")
			{
				users.GET("/:id", userHandler.GetUser)
				users.PUT("/:id", userHandler.UpdateUser)    // TODO: Implement update handler
				users.DELETE("/:id", userHandler.DeleteUser) // TODO: Implement delete handler

				// Admin only routes
				admin := users.Group("/admin")
				admin.Use(middleware.RoleMiddleware("admin"))
				{
					admin.GET("/", userHandler.ListUsers) // TODO: Implement list users handler
				}
			}
		}
	}

	return server
}

func (s *Server) Start(addr string) error {
	// Log server startup
	s.logger.WithField("address", addr).Info("Starting server")

	// Create a channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	// Create a channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Create http.Server instance
	srv := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	// Start the server in a goroutine
	go func() {
		s.logger.Info("Server is ready to handle requests")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErrors <- err
		}
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		s.logger.WithField("signal", sig.String()).Info("Start shutdown")

		// Create context for graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Attempt to gracefully shutdown the server
		if err := srv.Shutdown(ctx); err != nil {
			// If shutdown times out, force close
			s.logger.WithError(err).Error("Could not stop server gracefully")
			if err := srv.Close(); err != nil {
				return fmt.Errorf("could not stop server: %w", err)
			}
		}

		// Disconnect from database
		s.logger.Info("Disconnecting from database...")
		config.DisconnectDB()

		s.logger.Info("Server stopped gracefully")
		return nil
	}
}

// GetRouter returns the router for testing purposes
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
