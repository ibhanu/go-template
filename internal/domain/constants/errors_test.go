package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorConstants(t *testing.T) {
	// Test JWT errors
	assert.Equal(t, "unexpected JWT signing method", ErrUnexpectedSigningMethod.Error())
	assert.Equal(t, "invalid refresh token", ErrInvalidRefreshToken.Error())
	assert.Equal(t, "invalid token type", ErrInvalidTokenType.Error())

	// Test Repository errors
	assert.Equal(t, "user not found", ErrUserNotFound.Error())
	assert.Equal(t, "user already exists", ErrUserAlreadyExists.Error())
}

func TestAuthenticationErrors(t *testing.T) {
	tests := []struct {
		name     string
		errFunc  func() ErrorResponse
		wantCode string
		wantMsg  string
	}{
		{
			name:     "Auth header required",
			errFunc:  ErrAuthHeaderRequired,
			wantCode: "AUTH_HEADER_REQUIRED",
			wantMsg:  "Authorization header is required",
		},
		{
			name:     "Invalid auth format",
			errFunc:  ErrInvalidAuthFormat,
			wantCode: "INVALID_AUTH_FORMAT",
			wantMsg:  "Invalid authorization header format",
		},
		{
			name:     "Invalid token",
			errFunc:  ErrInvalidToken,
			wantCode: "INVALID_TOKEN",
			wantMsg:  "Invalid token",
		},
		{
			name:     "Role not found",
			errFunc:  ErrRoleNotFound,
			wantCode: "ROLE_NOT_FOUND",
			wantMsg:  "Role not found in context",
		},
		{
			name:     "Insufficient permissions",
			errFunc:  ErrInsufficientPermissions,
			wantCode: "INSUFFICIENT_PERMISSIONS",
			wantMsg:  "Insufficient permissions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errFunc()
			assert.Equal(t, tt.wantCode, err.Code)
			assert.Equal(t, tt.wantMsg, err.Message)
		})
	}
}

func TestEncryptionErrors(t *testing.T) {
	tests := []struct {
		name     string
		errFunc  func() ErrorResponse
		wantCode string
		wantMsg  string
	}{
		{
			name:     "Request body read error",
			errFunc:  ErrRequestBodyRead,
			wantCode: "REQUEST_BODY_READ_ERROR",
			wantMsg:  "Failed to read request body",
		},
		{
			name:     "Encryption error",
			errFunc:  ErrEncryption,
			wantCode: "ENCRYPTION_ERROR",
			wantMsg:  "Encryption error occurred",
		},
		{
			name:     "Decryption error",
			errFunc:  ErrDecryption,
			wantCode: "DECRYPTION_ERROR",
			wantMsg:  "Decryption error occurred",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errFunc()
			assert.Equal(t, tt.wantCode, err.Code)
			assert.Equal(t, tt.wantMsg, err.Message)
		})
	}
}

func TestGenericErrors(t *testing.T) {
	tests := []struct {
		name     string
		errFunc  func() ErrorResponse
		wantCode string
		wantMsg  string
	}{
		{
			name:     "Internal server error",
			errFunc:  ErrInternalServer,
			wantCode: "INTERNAL_SERVER_ERROR",
			wantMsg:  "Internal server error occurred",
		},
		{
			name:     "Invalid request",
			errFunc:  ErrInvalidRequest,
			wantCode: "INVALID_REQUEST",
			wantMsg:  "Invalid request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errFunc()
			assert.Equal(t, tt.wantCode, err.Code)
			assert.Equal(t, tt.wantMsg, err.Message)
		})
	}
}

func TestErrorResponse_Fields(t *testing.T) {
	err := ErrorResponse{
		Code:    "TEST_ERROR",
		Message: "Test error message",
	}

	assert.Equal(t, "TEST_ERROR", err.Code)
	assert.Equal(t, "Test error message", err.Message)
}
