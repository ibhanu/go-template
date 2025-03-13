package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewRateLimiter(t *testing.T) {
	capacity := 10.0
	refillRate := 2.0

	limiter := NewRateLimiter(capacity, refillRate)
	assert.Equal(t, capacity, limiter.tokens)
	assert.Equal(t, capacity, limiter.capacity)
	assert.Equal(t, refillRate, limiter.refillRate)
	assert.False(t, limiter.lastTimestamp.IsZero())
}

func TestRateLimiter_TryConsume(t *testing.T) {
	t.Run("Consume within capacity", func(t *testing.T) {
		limiter := NewRateLimiter(10.0, 2.0)
		assert.True(t, limiter.tryConsume(5.0))
		assert.Equal(t, 5.0, limiter.tokens)
	})

	t.Run("Consume more than available", func(t *testing.T) {
		limiter := NewRateLimiter(10.0, 2.0)
		limiter.tokens = 3.0
		assert.False(t, limiter.tryConsume(5.0))
		assert.Equal(t, 3.0, limiter.tokens)
	})

	t.Run("Consume exact amount", func(t *testing.T) {
		limiter := NewRateLimiter(10.0, 2.0)
		limiter.tokens = 5.0
		assert.True(t, limiter.tryConsume(5.0))
		assert.Equal(t, 0.0, limiter.tokens)
	})
}

func TestRateLimiter_Refill(t *testing.T) {
	t.Run("Partial refill", func(t *testing.T) {
		limiter := NewRateLimiter(10.0, 2.0) // 2 tokens per second
		limiter.tokens = 5.0
		limiter.lastTimestamp = time.Now().Add(-500 * time.Millisecond) // half second ago
		limiter.refill()

		// Should have added 1 token (2 tokens/sec * 0.5 sec)
		assert.InDelta(t, 6.0, limiter.tokens, 0.1)
	})

	t.Run("Full refill", func(t *testing.T) {
		limiter := NewRateLimiter(10.0, 2.0)
		limiter.tokens = 5.0
		limiter.lastTimestamp = time.Now().Add(-10 * time.Second)
		limiter.refill()

		// Should be capped at capacity
		assert.Equal(t, limiter.capacity, limiter.tokens)
	})
}

func TestRateLimiterMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Allow requests within limit", func(t *testing.T) {
		router := gin.New()
		router.Use(RateLimiterMiddleware(5.0, 1.0))
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		for i := 0; i < 5; i++ {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/test", nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		}
	})

	t.Run("Block requests over limit", func(t *testing.T) {
		router := gin.New()
		router.Use(RateLimiterMiddleware(2.0, 1.0))
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		// First two requests should succeed
		for i := 0; i < 2; i++ {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/test", nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		}

		// Third request should be blocked
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})

	t.Run("Allow requests after refill", func(t *testing.T) {
		router := gin.New()
		router.Use(RateLimiterMiddleware(1.0, 2.0)) // 2 tokens per second
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		// First request succeeds
		w1 := httptest.NewRecorder()
		req1 := httptest.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w1, req1)
		assert.Equal(t, http.StatusOK, w1.Code)

		// Second request fails immediately
		w2 := httptest.NewRecorder()
		req2 := httptest.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w2, req2)
		assert.Equal(t, http.StatusTooManyRequests, w2.Code)

		// Wait for refill
		time.Sleep(600 * time.Millisecond) // Should have refilled > 1 token

		// Third request succeeds
		w3 := httptest.NewRecorder()
		req3 := httptest.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w3, req3)
		assert.Equal(t, http.StatusOK, w3.Code)
	})
}

func TestMinFloat64(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{
			name: "First number smaller",
			a:    1.0,
			b:    2.0,
			want: 1.0,
		},
		{
			name: "Second number smaller",
			a:    2.0,
			b:    1.0,
			want: 1.0,
		},
		{
			name: "Equal numbers",
			a:    1.0,
			b:    1.0,
			want: 1.0,
		},
		{
			name: "With negative numbers",
			a:    -1.0,
			b:    1.0,
			want: -1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minFloat64(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
