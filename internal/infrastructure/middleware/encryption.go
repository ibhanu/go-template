package middleware

import (
"bytes"
"crypto/aes"
"crypto/cipher"
"encoding/base64"
"encoding/json"
"io"
"net/http"
"strings"

"github.com/gin-gonic/gin"
"web-server/internal/domain/constants"
"web-server/internal/infrastructure/config"
)

type encryptedBody struct {
	Data string `json:"data"`
}

func EncryptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()

		// Skip encryption for non-JSON requests
		if !strings.Contains(c.GetHeader("Content-Type"), "application/json") {
			c.Next()
			return
		}

		// Read the request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
		c.JSON(http.StatusBadRequest, constants.ErrRequestBodyRead)
		c.Abort()
		return
		}
		c.Request.Body.Close()

		// Only encrypt if body is not empty
		if len(body) > 0 {
			// Create cipher
			block, err := aes.NewCipher(cfg.EncryptionKey)
			if err != nil {
			c.JSON(http.StatusInternalServerError, constants.ErrEncryption)
			c.Abort()
			return
			}
			
			aesgcm, err := cipher.NewGCM(block)
			if err != nil {
			c.JSON(http.StatusInternalServerError, constants.ErrEncryption)
			c.Abort()
			return
			}

			// Encrypt the body
			encrypted := aesgcm.Seal(nil, cfg.EncryptionNonce, body, nil)
			encodedData := base64.StdEncoding.EncodeToString(encrypted)

			// Replace request body with encrypted data
			encBody := encryptedBody{Data: encodedData}
			newBody, err := json.Marshal(encBody)
			if err != nil {
			c.JSON(http.StatusInternalServerError, constants.ErrEncryption)
			c.Abort()
			return
			}

			c.Request.Body = io.NopCloser(bytes.NewBuffer(newBody))
			c.Request.ContentLength = int64(len(newBody))
		}

		// Create a custom response writer to intercept the response
		writer := &encryptionResponseWriter{
			ResponseWriter: c.Writer,
			cfg:           cfg,
		}
		c.Writer = writer

		c.Next()
	}
}

type encryptionResponseWriter struct {
	gin.ResponseWriter
	cfg *config.Config
}

func (w *encryptionResponseWriter) Write(data []byte) (int, error) {
	// Only encrypt JSON responses
	if !strings.Contains(w.Header().Get("Content-Type"), "application/json") {
		return w.ResponseWriter.Write(data)
	}

	// Create cipher
	block, err := aes.NewCipher(w.cfg.EncryptionKey)
	if err != nil {
		return 0, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return 0, err
	}

	// Encrypt the response
	encrypted := aesgcm.Seal(nil, w.cfg.EncryptionNonce, data, nil)
	encodedData := base64.StdEncoding.EncodeToString(encrypted)

	// Create encrypted response
	encResp := encryptedBody{Data: encodedData}
	newData, err := json.Marshal(encResp)
	if err != nil {
		return 0, err
	}

	return w.ResponseWriter.Write(newData)
}

func DecryptRequestBody(c *gin.Context) ([]byte, error) {
	cfg := config.GetConfig()

	var encBody encryptedBody
	if err := c.ShouldBindJSON(&encBody); err != nil {
		return nil, err
	}

	encrypted, err := base64.StdEncoding.DecodeString(encBody.Data)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(cfg.EncryptionKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	decrypted, err := aesgcm.Open(nil, cfg.EncryptionNonce, encrypted, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}