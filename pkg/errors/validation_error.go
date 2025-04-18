package errors

// ValidationError represents validation related errors
type ValidationError struct {
	Errors []ValidationFieldError
}

// ValidationFieldError represents a field-specific validation error
type ValidationFieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidationError(errors []ValidationFieldError) *ValidationError {
	return &ValidationError{
		Errors: errors,
	}
}

func (e *ValidationError) Error() string {
	return "Validation failed"
}
