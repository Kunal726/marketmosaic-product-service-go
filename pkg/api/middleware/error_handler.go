package middleware

import (
	"fmt"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/errors"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/validation"
	"net/http"
	"strings"

	"github.com/Kunal726/market-mosaic-common-lib-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// ErrorHandler middleware handles all errors in a consistent way
func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		logger := utils.GetLoggerFromContext(c)
		logger.Error("Handling error",
			zap.Error(err),
			zap.String("path", c.Request.URL.Path))

		switch e := err.(type) {
		case validator.ValidationErrors:
			logger.Debug("Handling validator.ValidationErrors",
				zap.Int("error_count", len(e)))
			handleValidatorErrors(c, e)

		case validation.Errors:
			logger.Debug("Handling validation.Errors",
				zap.Int("error_count", len(e)))
			handleCustomValidationErrors(c, e)

		case *validation.ValidationError:
			logger.Debug("Handling ValidationError",
				zap.Int("error_count", len(e.Errors)))
			handleValidationErrorType(c, e)

		case *errors.AuthError:
			logger.Debug("Handling AuthError")
			handleAuthError(c, e)

		case *errors.ProductError:
			logger.Debug("Handling ProductError",
				zap.Int("error_code", e.Code))
			handleProductError(c, e)

		default:
			logger.Error("Unhandled error type",
				zap.String("type", fmt.Sprintf("%T", err)),
				zap.Error(err))
			handleUnknownError(c, err)
		}
	}
}

// handleValidatorErrors handles validator.ValidationErrors
func handleValidatorErrors(c *gin.Context, validationErrors validator.ValidationErrors) {
	validationMap := make(map[string]string)
	for _, err := range validationErrors {
		field := strings.ToLower(err.Field())
		validationMap[field] = fmt.Sprintf(
			"Field validation for '%s' failed on the '%s' tag",
			field,
			err.Tag(),
		)
	}

	c.JSON(http.StatusBadRequest, &errors.BaseErrorResponse{
		Status:  false,
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		AdditionalInfo: map[string]interface{}{
			"errors": validationMap,
		},
	})
}

// handleCustomValidationErrors handles our custom validation.Errors
func handleCustomValidationErrors(c *gin.Context, validationErrors validation.Errors) {
	validationMap := make(map[string]string)
	for _, err := range validationErrors {
		validationMap[err.Field] = err.Message
	}

	c.JSON(http.StatusBadRequest, &errors.BaseErrorResponse{
		Status:  false,
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		AdditionalInfo: map[string]interface{}{
			"errors": validationMap,
		},
	})
}

// handleValidationErrorType handles *validation.ValidationError
func handleValidationErrorType(c *gin.Context, err *validation.ValidationError) {
	validationErrors := make(map[string]string)
	for _, e := range err.Errors {
		validationErrors[e.Field] = e.Message
	}

	c.JSON(err.ErrorCode, &errors.BaseErrorResponse{
		Status:  false,
		Code:    err.ErrorCode,
		Message: "Validation failed",
		AdditionalInfo: map[string]interface{}{
			"errors": validationErrors,
		},
	})
}

// handleProductError handles product-specific errors
func handleProductError(c *gin.Context, err *errors.ProductError) {
	c.JSON(err.Code, &errors.BaseErrorResponse{
		Status:  false,
		Code:    err.Code,
		Message: err.Error(),
	})
}

// handleAuthError handles authentication related errors
func handleAuthError(c *gin.Context, err *errors.AuthError) {
	c.JSON(http.StatusUnauthorized, err.ToResponse())
}

// handleUnknownError handles any other unspecified errors
func handleUnknownError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, &errors.BaseErrorResponse{
		Status:  false,
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	})
}
