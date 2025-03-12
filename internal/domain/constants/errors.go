package constants

import "errors"

// HTTP Error Messages.
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// JWT errors.
var (
	ErrUnexpectedSigningMethod = errors.New("unexpected JWT signing method")
	ErrInvalidRefreshToken     = errors.New("invalid refresh token")
	ErrInvalidTokenType        = errors.New("invalid token type")
)

// Repository errors.
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

// Authentication Errors.
func ErrAuthHeaderRequired() ErrorResponse {
	return ErrorResponse{
		Code:    "AUTH_HEADER_REQUIRED",
		Message: "Authorization header is required",
	}
}

func ErrInvalidAuthFormat() ErrorResponse {
	return ErrorResponse{
		Code:    "INVALID_AUTH_FORMAT",
		Message: "Invalid authorization header format",
	}
}

func ErrInvalidToken() ErrorResponse {
	return ErrorResponse{
		Code:    "INVALID_TOKEN",
		Message: "Invalid token",
	}
}

func ErrRoleNotFound() ErrorResponse {
	return ErrorResponse{
		Code:    "ROLE_NOT_FOUND",
		Message: "Role not found in context",
	}
}

func ErrInsufficientPermissions() ErrorResponse {
	return ErrorResponse{
		Code:    "INSUFFICIENT_PERMISSIONS",
		Message: "Insufficient permissions",
	}
}

// Encryption Errors.
func ErrRequestBodyRead() ErrorResponse {
	return ErrorResponse{
		Code:    "REQUEST_BODY_READ_ERROR",
		Message: "Failed to read request body",
	}
}

func ErrEncryption() ErrorResponse {
	return ErrorResponse{
		Code:    "ENCRYPTION_ERROR",
		Message: "Encryption error occurred",
	}
}

func ErrDecryption() ErrorResponse {
	return ErrorResponse{
		Code:    "DECRYPTION_ERROR",
		Message: "Decryption error occurred",
	}
}

// Generic Errors.
func ErrInternalServer() ErrorResponse {
	return ErrorResponse{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Internal server error occurred",
	}
}

func ErrInvalidRequest() ErrorResponse {
	return ErrorResponse{
		Code:    "INVALID_REQUEST",
		Message: "Invalid request",
	}
}
