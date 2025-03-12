package constants

// HTTP Error Messages
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	// Authentication Errors
	ErrAuthHeaderRequired = ErrorResponse{
		Code:    "AUTH_HEADER_REQUIRED",
		Message: "Authorization header is required",
	}
	ErrInvalidAuthFormat = ErrorResponse{
		Code:    "INVALID_AUTH_FORMAT",
		Message: "Invalid authorization header format",
	}
	ErrInvalidToken = ErrorResponse{
		Code:    "INVALID_TOKEN",
		Message: "Invalid token",
	}
	ErrRoleNotFound = ErrorResponse{
		Code:    "ROLE_NOT_FOUND",
		Message: "Role not found in context",
	}
	ErrInsufficientPermissions = ErrorResponse{
		Code:    "INSUFFICIENT_PERMISSIONS",
		Message: "Insufficient permissions",
	}

	// Encryption Errors
	ErrRequestBodyRead = ErrorResponse{
		Code:    "REQUEST_BODY_READ_ERROR",
		Message: "Failed to read request body",
	}
	ErrEncryption = ErrorResponse{
		Code:    "ENCRYPTION_ERROR",
		Message: "Encryption error occurred",
	}
	ErrDecryption = ErrorResponse{
		Code:    "DECRYPTION_ERROR",
		Message: "Decryption error occurred",
	}

	// Generic Errors
	ErrInternalServer = ErrorResponse{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Internal server error occurred",
	}
	ErrInvalidRequest = ErrorResponse{
		Code:    "INVALID_REQUEST",
		Message: "Invalid request",
	}
)
