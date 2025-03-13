package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
		email    string
		password string
		want     *User
	}{
		{
			name:     "Create new user with valid data",
			username: "testuser",
			email:    "test@example.com",
			password: "password123",
			want: &User{
				ID:       "",
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
				Role:     "user",
			},
		},
		{
			name:     "Create new user with minimum length username",
			username: "abc",
			email:    "abc@example.com",
			password: "password123",
			want: &User{
				ID:       "",
				Username: "abc",
				Email:    "abc@example.com",
				Password: "password123",
				Role:     "user",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUser(tt.username, tt.email, tt.password)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUser_Fields(t *testing.T) {
	// Test struct tags and field assignments
	user := &User{
		ID:       "123",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "admin",
	}

	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "password123", user.Password)
	assert.Equal(t, "admin", user.Role)
}
