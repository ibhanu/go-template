package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	tokens        float64    // current number of tokens
	capacity      float64    // maximum number of tokens
	refillRate    float64    // tokens per second to refill
	lastTimestamp time.Time  // last time tokens were refilled
	mutex         sync.Mutex // mutex for thread safety
}

func NewRateLimiter(capacity float64, refillRate float64) *RateLimiter {
	return &RateLimiter{
		tokens:        capacity,
		capacity:      capacity,
		refillRate:    refillRate,
		lastTimestamp: time.Now(),
	}
}

func (rl *RateLimiter) refill() {
	now := time.Now()
	duration := now.Sub(rl.lastTimestamp).Seconds()
	newTokens := duration * rl.refillRate
	rl.tokens = min(rl.capacity, rl.tokens+newTokens)
	rl.lastTimestamp = now
}

func (rl *RateLimiter) tryConsume(tokens float64) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	rl.refill()
	if rl.tokens >= tokens {
		rl.tokens -= tokens
		return true
	}
	return false
}

// min returns the smaller of two float64 numbers
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// RateLimiterMiddleware creates a middleware that limits requests using the token bucket algorithm
// capacity: maximum number of tokens (requests)
// refillRate: number of tokens (requests) to refill per second
func RateLimiterMiddleware(capacity, refillRate float64) gin.HandlerFunc {
	limiter := NewRateLimiter(capacity, refillRate)

	return func(c *gin.Context) {
		if !limiter.tryConsume(1.0) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Please try again later.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
