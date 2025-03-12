package handler

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"Error message here"`
}

// MessageResponse represents a success message response
type MessageResponse struct {
	Message string `json:"message" example:"Operation completed successfully"`
}
