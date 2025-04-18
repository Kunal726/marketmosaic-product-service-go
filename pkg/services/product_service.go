package services

import (
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/dtos"
	commonDto "github.com/Kunal726/market-mosaic-common-lib-go/pkg/dtos"
	"go.uber.org/zap"
)

type ProductService interface {
	// AddProduct adds a single product
	AddProduct(logger *zap.Logger, productDTO *dtos.ProductDetailsDTO) (*commonDto.BaseResponseDTO, error)

	// AddProducts adds multiple products
	AddProducts(logger *zap.Logger, products []dtos.ProductDetailsDTO) (*commonDto.BaseResponseDTO, error)

	// UpdateProduct updates a single product
	UpdateProduct(logger *zap.Logger, productID string, updateDTO *dtos.UpdateProductRequestDTO) (*commonDto.BaseResponseDTO, error)

	// UpdateProducts updates multiple products
	UpdateProducts(logger *zap.Logger, updates map[string]dtos.UpdateProductRequestDTO) (*commonDto.BaseResponseDTO, error)

	// DeleteProduct deletes or deactivates a product
	DeleteProduct(logger *zap.Logger, productID string, deactivate bool) (*commonDto.BaseResponseDTO, error)

	// GetProduct retrieves a single product by ID
	GetProduct(logger *zap.Logger, productID string) (*dtos.ProductResponseDTO, error)

	// GetProductList retrieves products based on filters
	GetProductList(logger *zap.Logger, filters *dtos.ProductFilterDTO) (*dtos.ProductResponseDTO, error)

	// GetProductSuggestions returns product name suggestions based on a query
	GetProductSuggestions(logger *zap.Logger, query string) ([]string, error)
}
