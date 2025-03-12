package config

import (
	"crypto/rand"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

const (
	jwtKeySize           = 32 // Size for JWT secret keys (256 bits)
	encryptionKeySize    = 32 // Size for AES-256 encryption key
	encryptionNonceSize  = 12 // Size for AES-GCM nonce
	accessTokenDuration  = 15 * time.Minute
	refreshTokenDuration = 7 * 24 * time.Hour
)

type Config struct {
	JWTSecret            []byte
	JWTExpiration        time.Duration
	JWTRefreshSecret     []byte
	JWTRefreshExpiration time.Duration
	EncryptionKey        []byte
	EncryptionNonce      []byte
}

func generateRandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// singleton instance
var (
	configInstance *Config
	configOnce     sync.Once
)

// GetConfig returns the singleton config instance.
func GetConfig() *Config {
	configOnce.Do(initConfig)
	return configInstance
}

func initConfig() {
	if godotenv.Load() != nil {
		// If .env file doesn't exist, use generated secrets
		jwtSecret, err := generateRandomBytes(jwtKeySize)
		if err != nil {
			panic(err)
		}

		encKey, err := generateRandomBytes(encryptionKeySize)
		if err != nil {
			panic(err)
		}

		nonce, err := generateRandomBytes(encryptionNonceSize)
		if err != nil {
			panic(err)
		}

		refreshSecret, err := generateRandomBytes(jwtKeySize)
		if err != nil {
			panic(err)
		}

		configInstance = &Config{
			JWTSecret:            jwtSecret,
			JWTExpiration:        accessTokenDuration,
			JWTRefreshSecret:     refreshSecret,
			JWTRefreshExpiration: refreshTokenDuration,
			EncryptionKey:        encKey,
			EncryptionNonce:      nonce,
		}
		return
	}

	// If .env file exists, use values from it
	configInstance = &Config{
		JWTSecret:            []byte(os.Getenv("JWT_SECRET")),
		JWTExpiration:        accessTokenDuration,
		JWTRefreshSecret:     []byte(os.Getenv("JWT_REFRESH_SECRET")),
		JWTRefreshExpiration: refreshTokenDuration,
		EncryptionKey:        []byte(os.Getenv("ENCRYPTION_KEY")),
		EncryptionNonce:      []byte(os.Getenv("ENCRYPTION_NONCE")),
	}
}
