package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomBytes(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantErr bool
	}{
		{
			name:    "Generate valid random bytes",
			size:    32,
			wantErr: false,
		},
		{
			name:    "Generate zero bytes",
			size:    0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := generateRandomBytes(tt.size)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, bytes, tt.size)

			// Test randomness by generating another set
			bytes2, _ := generateRandomBytes(tt.size)
			if tt.size > 0 {
				assert.NotEqual(t, bytes, bytes2, "Random bytes should be different")
			}
		})
	}
}

func TestGetConfig_WithEnvFile(t *testing.T) {
	// Setup test environment variables
	envVars := map[string]string{
		"JWT_SECRET":         "test-jwt-secret",
		"JWT_REFRESH_SECRET": "test-refresh-secret",
		"ENCRYPTION_KEY":     "test-encryption-key",
		"ENCRYPTION_NONCE":   "test-encryption-nonce",
	}

	for k, v := range envVars {
		os.Setenv(k, v)
	}
	defer func() {
		for k := range envVars {
			os.Unsetenv(k)
		}
		// Reset singleton for other tests
		configInstance = nil
	}()

	config := GetConfig()
	assert.NotNil(t, config)

	// Test values from environment
	assert.Equal(t, envVars["JWT_SECRET"], string(config.JWTSecret))
	assert.Equal(t, envVars["JWT_REFRESH_SECRET"], string(config.JWTRefreshSecret))
	assert.Equal(t, envVars["ENCRYPTION_KEY"], string(config.EncryptionKey))
	assert.Equal(t, envVars["ENCRYPTION_NONCE"], string(config.EncryptionNonce))

	// Test durations
	assert.Equal(t, 15*time.Minute, config.JWTExpiration)
	assert.Equal(t, 7*24*time.Hour, config.JWTRefreshExpiration)

	// Test singleton behavior
	config2 := GetConfig()
	assert.Same(t, config, config2)
}

func TestGetConfig_WithoutEnvFile(t *testing.T) {
	// Ensure no environment variables are set
	envVars := []string{
		"JWT_SECRET",
		"JWT_REFRESH_SECRET",
		"ENCRYPTION_KEY",
		"ENCRYPTION_NONCE",
	}
	for _, v := range envVars {
		os.Unsetenv(v)
	}
	defer func() {
		// Reset singleton for other tests
		configInstance = nil
	}()

	config := GetConfig()
	assert.NotNil(t, config)

	// Test generated values
	assert.Len(t, config.JWTSecret, jwtKeySize)
	assert.Len(t, config.JWTRefreshSecret, jwtKeySize)
	assert.Len(t, config.EncryptionKey, encryptionKeySize)
	assert.Len(t, config.EncryptionNonce, encryptionNonceSize)

	// Test durations
	assert.Equal(t, 15*time.Minute, config.JWTExpiration)
	assert.Equal(t, 7*24*time.Hour, config.JWTRefreshExpiration)

	// Verify values remain consistent across calls
	config2 := GetConfig()
	assert.Equal(t, config.JWTSecret, config2.JWTSecret)
	assert.Equal(t, config.JWTRefreshSecret, config2.JWTRefreshSecret)
	assert.Equal(t, config.EncryptionKey, config2.EncryptionKey)
	assert.Equal(t, config.EncryptionNonce, config2.EncryptionNonce)
}

func TestConstants(t *testing.T) {
	assert.Equal(t, 32, jwtKeySize)
	assert.Equal(t, 32, encryptionKeySize)
	assert.Equal(t, 12, encryptionNonceSize)
	assert.Equal(t, 15*time.Minute, accessTokenDuration)
	assert.Equal(t, 7*24*time.Hour, refreshTokenDuration)
}
