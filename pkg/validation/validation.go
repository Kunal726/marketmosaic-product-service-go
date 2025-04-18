package validation

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Error represents a validation error with a custom message
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Errors represents a collection of validation errors
type Errors []Error

// Error implements the error interface for ValidationErrors
func (v Errors) Error() string {
	if len(v) == 0 {
		return ""
	}
	var messages []string
	for _, err := range v {
		messages = append(messages, fmt.Sprintf("%s: %s", err.Field, err.Message))
	}
	return fmt.Sprintf("Validation failed:\n- %s", strings.Join(messages, "\n- "))
}

// IsValidationError checks if an error is a validation error
func IsValidationError(err error) bool {
	_, ok := err.(Errors)
	return ok
}

// GetValidationErrors converts an error to validation errors if possible
func GetValidationErrors(err error) (Errors, bool) {
	if verr, ok := err.(Errors); ok {
		return verr, true
	}
	return nil, false
}

// Validate validates a struct and returns custom validation errors
// It uses struct tags for validation rules and custom error messages
// Example usage:
//
//	type User struct {
//	    Name  string `validate:"required" message:"Name is required"`
//	    Email string `validate:"required,email" message:"Please enter a valid email"`
//	}
func Validate(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	var errors Errors

	// Get the type of the struct for reflection
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Get validation errors
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		if err != nil {
			errors = append(errors, Error{
				Field:   "unknown",
				Message: err.Error(),
			})
		}
		return errors
	}

	for _, err := range validationErrors {
		field, ok := t.FieldByName(err.Field())
		if !ok {
			errors = append(errors, Error{
				Field:   err.Field(),
				Message: fmt.Sprintf("Validation failed on field %s", err.Field()),
			})
			continue
		}

		message := field.Tag.Get("message")
		if message == "" {
			message = fmt.Sprintf("Validation failed on field %s (tag: %s, value: %v)", err.Field(), err.Tag(), err.Value())
		}

		errors = append(errors, Error{
			Field:   err.Field(),
			Message: message,
		})
	}

	if len(errors) == 0 {
		return nil
	}

	return errors
}

// ValidationError represents a custom validation error with error code and message
type ValidationError struct {
	Errors    Errors
	ErrorCode int
}

// NewValidationError creates a new ValidationError instance
func NewValidationError(errors Errors) *ValidationError {
	return &ValidationError{
		Errors:    errors,
		ErrorCode: http.StatusBadRequest,
	}
}

// Error implements the error interface for ValidationError
func (v *ValidationError) Error() string {
	if len(v.Errors) == 0 {
		return ""
	}
	var messages []string
	for _, err := range v.Errors {
		messages = append(messages, fmt.Sprintf("%s: %s", err.Field, err.Message))
	}
	return strings.Join(messages, "\n- ")
}
