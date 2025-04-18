package dtos

import (
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/validation"
)

// UpdateProductRequestDTO represents the update product request structure
type UpdateProductRequestDTO struct {
	Price  *float64 `json:"price,omitempty" validate:"omitempty,gt=0" message:"Price must be positive"`
	Stock  *int     `json:"stock,omitempty" validate:"omitempty,min=0" message:"Stock must be greater than or equal to 0"`
	Rating *string  `json:"rating,omitempty" validate:"omitempty,max=5" message:"Rating should be at most 5 characters"`
}

// Validate validates the UpdateProductRequestDTO
func (u *UpdateProductRequestDTO) Validate() error {
	return validation.Validate(u)
}

// GetValidationMessage returns the custom validation message for a field
func (u *UpdateProductRequestDTO) GetValidationMessage(field string) string {
	switch field {
	case "Price":
		return "Price must be greater than 0"
	case "Stock":
		return "Stock must be greater than or equal to 0"
	case "Rating":
		return "Rating should be at most 5 characters"
	default:
		return ""
	}
}
