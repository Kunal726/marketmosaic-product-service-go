package dtos

import (
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/validation"
)

// ProductDetailsDTO represents the data transfer object for product details
type ProductDetailsDTO struct {
	ProductID     string   `json:"productId" validate:"omitempty" message:"Invalid product ID"`
	ProductName   string   `json:"productName" validate:"required" message:"Product name is required"`
	CategoryID    string   `json:"categoryId" validate:"required" message:"Category ID is required"`
	Price         float64  `json:"price" validate:"required,gt=0" message:"Price must be greater than 0"`
	Description   string   `json:"description" validate:"required" message:"Description is required"`
	ImageURL      string   `json:"imageUrl" validate:"required,url" message:"Invalid image URL"`
	IsActive      *bool    `json:"isActive,omitempty"`
	StockQuantity int      `json:"stockQuantity" validate:"required,gte=0" message:"Stock quantity must be greater than or equal to 0"`
	SupplierID    string   `json:"supplierId" validate:"required" message:"Supplier ID is required"`
	TagIds        []string `json:"tagIds,omitempty" validate:"omitempty,dive,required" message:"Invalid tag ID"`
}

// Validate validates the ProductDetailsDTO
func (p *ProductDetailsDTO) Validate() error {
	return validation.Validate(p)
}

// GetValidationMessage returns the custom validation message for a field
func (p *ProductDetailsDTO) GetValidationMessage(field string) string {
	switch field {
	case "ProductName":
		return "Product name cannot be blank and must be between 3 and 100 characters"
	case "Description":
		return "Description cannot exceed 500 characters"
	case "Price":
		return "Price cannot be null and must be positive"
	case "CategoryID":
		return "Category ID cannot be null"
	case "StockQuantity":
		return "Stock quantity must be greater than or equal to 0"
	case "ImageURL":
		return "Image URL cannot be null"
	case "SupplierID":
		return "Supplier ID cannot be null"
	default:
		return ""
	}
}
