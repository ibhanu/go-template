package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// responseWriter is a custom response writer that captures the response status and body
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// LoggerMiddleware creates a middleware for logging requests and responses
func LoggerMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Create a copy of the request body
		var reqBody []byte
		if c.Request.Body != nil {
			reqBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		// Create custom response writer
		w := &responseWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = w

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Create log entry
		entry := log.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"duration":   duration.String(),
			"user_agent": c.Request.UserAgent(),
		})

		// Add request ID if present
		if requestID := c.GetString("RequestID"); requestID != "" {
			entry = entry.WithField("request_id", requestID)
		}

		// Add user ID if authenticated
		if userID, exists := c.Get("userID"); exists {
			entry = entry.WithField("user_id", userID)
		}

		// Log based on status code
		statusCode := c.Writer.Status()
		switch {
		case statusCode >= 500:
			entry.Error("Server error")
		case statusCode >= 400:
			entry.Warn("Client error")
		default:
			entry.Info("Request processed")
		}
	}
}

// InitLogger initializes and configures the logger
func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// Configure logger
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// Set log level based on environment
	if gin.Mode() == gin.DebugMode {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return logger
}
