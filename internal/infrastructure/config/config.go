package config

import (
	"crypto/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret       []byte
	JWTExpiration   time.Duration
	EncryptionKey   []byte
	EncryptionNonce []byte
}

var instance *Config

func generateRandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func LoadConfig() (*Config, error) {
	if instance != nil {
		return instance, nil
	}

	if err := godotenv.Load(); err != nil {
		// If .env file doesn't exist, we'll use generated secrets
		// This is fine for development but in production you should use real secrets
		jwtSecret, err := generateRandomBytes(32)
		if err != nil {
			return nil, err
		}

		encKey, err := generateRandomBytes(32) // AES-256 key
		if err != nil {
			return nil, err
		}

		nonce, err := generateRandomBytes(12) // For AES-GCM
		if err != nil {
			return nil, err
		}

		instance = &Config{
			JWTSecret:       jwtSecret,
			JWTExpiration:   24 * time.Hour,
			EncryptionKey:   encKey,
			EncryptionNonce: nonce,
		}
	} else {
		// If .env file exists, use values from it
		instance = &Config{
			JWTSecret:       []byte(os.Getenv("JWT_SECRET")),
			JWTExpiration:   24 * time.Hour,
			EncryptionKey:   []byte(os.Getenv("ENCRYPTION_KEY")),
			EncryptionNonce: []byte(os.Getenv("ENCRYPTION_NONCE")),
		}
	}

	return instance, nil
}

func GetConfig() *Config {
	if instance == nil {
		var err error
		instance, err = LoadConfig()
		if err != nil {
			panic(err)
		}
	}
	return instance
}
