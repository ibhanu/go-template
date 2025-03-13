package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"web-server/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(EncryptionMiddleware())
	return r
}

func TestEncryptionMiddleware(t *testing.T) {
	router := setupTestRouter()

	t.Run("Handles non-JSON request", func(t *testing.T) {
		router.POST("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "plain text")
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", bytes.NewBufferString("plain text"))
		req.Header.Set("Content-Type", "text/plain")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "plain text", w.Body.String())
	})

	t.Run("Encrypts JSON request and response", func(t *testing.T) {
		router.POST("/test", func(c *gin.Context) {
			var data map[string]interface{}
			decrypted, err := DecryptRequestBody(c)
			assert.NoError(t, err)

			err = json.Unmarshal(decrypted, &data)
			assert.NoError(t, err)
			assert.Equal(t, "test", data["message"])

			c.JSON(http.StatusOK, map[string]string{"response": "success"})
		})

		reqBody := map[string]string{"message": "test"}
		reqJSON, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Verify response is encrypted
		var encResp encryptedBody
		err := json.Unmarshal(w.Body.Bytes(), &encResp)
		assert.NoError(t, err)
		assert.NotEmpty(t, encResp.Data)
	})

	t.Run("Handles empty JSON body", func(t *testing.T) {
		router.POST("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", bytes.NewBuffer([]byte{}))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Handles invalid JSON request", func(t *testing.T) {
		router.POST("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDecryptRequestBody(t *testing.T) {
	router := setupTestRouter()

	t.Run("Successfully decrypts valid request", func(t *testing.T) {
		router.POST("/test", func(c *gin.Context) {
			decrypted, err := DecryptRequestBody(c)
			assert.NoError(t, err)

			var data map[string]interface{}
			err = json.Unmarshal(decrypted, &data)
			assert.NoError(t, err)
			assert.Equal(t, "test", data["message"])

			c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})

		reqBody := map[string]string{"message": "test"}
		reqJSON, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Handles invalid encrypted data", func(t *testing.T) {
		router.POST("/test", func(c *gin.Context) {
			_, err := DecryptRequestBody(c)
			assert.Error(t, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encrypted data"})
		})

		encBody := encryptedBody{Data: "invalid-base64"}
		reqJSON, _ := json.Marshal(encBody)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestEncryptionResponseWriter(t *testing.T) {
	t.Run("Write non-JSON response", func(t *testing.T) {
		router := setupTestRouter()
		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "plain text")
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "plain text", w.Body.String())
	})

	t.Run("Write JSON response", func(t *testing.T) {
		router := setupTestRouter()
		router.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "test"})
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var encResp encryptedBody
		err := json.Unmarshal(w.Body.Bytes(), &encResp)
		assert.NoError(t, err)
		assert.NotEmpty(t, encResp.Data)
	})
}

func TestEncryptionWithInvalidKey(t *testing.T) {
	// Store original encryption key and nonce
	origKey := os.Getenv("ENCRYPTION_KEY")
	origNonce := os.Getenv("ENCRYPTION_NONCE")
	defer func() {
		// Restore original environment variables after test
		if origKey != "" {
			os.Setenv("ENCRYPTION_KEY", origKey)
		} else {
			os.Unsetenv("ENCRYPTION_KEY")
		}
		if origNonce != "" {
			os.Setenv("ENCRYPTION_NONCE", origNonce)
		} else {
			os.Unsetenv("ENCRYPTION_NONCE")
		}
	}()

	// Set test environment variables with invalid key
	os.Setenv("ENCRYPTION_KEY", "invalid-key")
	os.Setenv("ENCRYPTION_NONCE", string(make([]byte, 12)))

	// Reset config instance to force new config creation with our test values
	config.GetConfig()

	router := setupTestRouter()
	router.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/test", bytes.NewBufferString("{}"))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
