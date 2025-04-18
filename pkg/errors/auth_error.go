package errors

import "net/http"

// AuthError represents authentication related errors
type AuthError struct {
	Message string
	Code    int
}

// NewAuthError creates a new AuthError instance
func NewAuthError(message string) *AuthError {
	return &AuthError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func (e *AuthError) Error() string {
	return e.Message
}

// ToResponse converts the error to a BaseErrorResponse
func (e *AuthError) ToResponse() *BaseErrorResponse {
	return &BaseErrorResponse{
		Status:  false,
		Code:    e.Code,
		Message: e.Message,
	}
}
