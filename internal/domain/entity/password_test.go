package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		shouldError bool
	}{
		{
			name:        "Valid password hash",
			password:    "mypassword123",
			shouldError: false,
		},
		{
			name:        "Empty password hash",
			password:    "",
			shouldError: false,
		},
		{
			name:        "Long password hash",
			password:    "verylongpasswordthatishardtoguess123!@#",
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)

			if tt.shouldError {
				assert.Error(t, err)
				assert.Empty(t, hash)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, hash)
				assert.NotEqual(t, tt.password, hash)
				// Verify the hash is valid by checking it
				assert.True(t, CheckPassword(tt.password, hash))
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	password := "mypassword123"
	hash, err := HashPassword(password)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{
			name:     "Correct password",
			password: password,
			hash:     hash,
			want:     true,
		},
		{
			name:     "Incorrect password",
			password: "wrongpassword",
			hash:     hash,
			want:     false,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash,
			want:     false,
		},
		{
			name:     "Invalid hash format",
			password: password,
			hash:     "invalid_hash",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPassword(tt.password, tt.hash)
			assert.Equal(t, tt.want, got)
		})
	}
}
