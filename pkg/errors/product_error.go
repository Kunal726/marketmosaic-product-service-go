package errors

// ProductError represents product-related errors
type ProductError struct {
	Code    int
	Message string
}

// NewProductError creates a new ProductError instance
func NewProductError(code int, message string) *ProductError {
	return &ProductError{
		Code:    code,
		Message: message,
	}
}

// Error implements the error interface for ProductError
func (e *ProductError) Error() string {
	return e.Message
}
