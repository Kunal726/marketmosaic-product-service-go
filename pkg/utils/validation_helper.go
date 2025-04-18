package utils

import (
	"fmt"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/errors"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/validation"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

// ValidateRequest validates a request DTO and returns appropriate errors
func ValidateRequest(logger *zap.Logger, dto interface{}, entityName string) error {
	if dto == nil {
		logger.Error(fmt.Sprintf("%s cannot be null", entityName))
		return errors.NewProductError(http.StatusBadRequest, fmt.Sprintf("%s cannot be null", entityName))
	}

	if validator, ok := dto.(validation.Validator); ok {
		logger.Info(fmt.Sprintf("Validating %s", entityName))
		if err := validator.Validate(); err != nil {
			logger.Error("Validation error", zap.Error(err))
			if validationErrs, ok := validation.GetValidationErrors(err); ok {
				return validation.NewValidationError(validationErrs)
			}
			return errors.NewProductError(http.StatusBadRequest, err.Error())
		}
	}

	return nil
}

// ValidateID validates and parses a string ID to uint
func ValidateID(logger *zap.Logger, id string, entityName string) (uint, error) {
	if id == "" {
		logger.Error(fmt.Sprintf("%s ID cannot be empty", entityName))
		return 0, errors.NewProductError(http.StatusBadRequest, fmt.Sprintf("%s ID cannot be empty", entityName))
	}

	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logger.Error(fmt.Sprintf("Invalid %s ID", entityName),
			zap.Error(err),
			zap.String("id", id))
		return 0, errors.NewProductError(http.StatusBadRequest, fmt.Sprintf("Invalid %s ID", entityName))
	}

	return uint(parsedID), nil
}

// ValidateEntityExists checks if an entity exists in the database
func ValidateEntityExists[T any](logger *zap.Logger, entity *T, entityName string) error {
	if entity == nil {
		logger.Error(fmt.Sprintf("%s not found", entityName))
		return errors.NewProductError(http.StatusNotFound, fmt.Sprintf("%s not found", entityName))
	}
	return nil
}
