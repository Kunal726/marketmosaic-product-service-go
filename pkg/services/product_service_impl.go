package services

import (
	"fmt"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/dtos"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/errors"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/repositories"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/utils"
	"go.uber.org/zap"
	"net/http"

	commonDto "github.com/Kunal726/market-mosaic-common-lib-go/pkg/dtos"
)

type productService struct {
	productRepo  repositories.ProductRepository
	userRepo     repositories.UserRepository
	productUtils *utils.ProductUtils
}

// NewProductService creates a new instance of ProductService
func NewProductService(
	productRepo repositories.ProductRepository,
	userRepo repositories.UserRepository,
	productUtils *utils.ProductUtils,
) ProductService {
	return &productService{
		productRepo:  productRepo,
		userRepo:     userRepo,
		productUtils: productUtils,
	}
}

func (s *productService) AddProduct(logger *zap.Logger, productDTO *dtos.ProductDetailsDTO) (*commonDto.BaseResponseDTO, error) {
	if err := utils.ValidateRequest(logger, productDTO, "Product details"); err != nil {
		return nil, err
	}

	logger.Info("Mapping product entity")
	product, err := s.productUtils.MapProductEntity(logger, productDTO)
	if err != nil {
		logger.Error("Failed to map product entity", zap.Error(err))
		return nil, errors.NewProductError(http.StatusInternalServerError, err.Error())
	}

	logger.Info("Saving product to database")
	if err := s.productRepo.Save(logger, product); err != nil {
		logger.Error("Failed to save product", zap.Error(err))
		return nil, errors.NewProductError(http.StatusInternalServerError, "Failed to save product")
	}

	logger.Info("Product added successfully",
		zap.String("product_name", product.ProductName),
		zap.Uint("product_id", product.ProductId))
	return &commonDto.BaseResponseDTO{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Product Added",
	}, nil
}

func (s *productService) AddProducts(logger *zap.Logger, products []dtos.ProductDetailsDTO) (*commonDto.BaseResponseDTO, error) {
	if len(products) == 0 {
		logger.Error("Product list cannot be empty")
		return nil, errors.NewProductError(http.StatusBadRequest, "Product list cannot be empty")
	}

	logger.Info("Processing products", zap.Int("count", len(products)))
	for _, dto := range products {
		if _, err := s.AddProduct(logger, &dto); err != nil {
			return nil, err
		}
	}

	logger.Info("Products added successfully", zap.Int("count", len(products)))
	return &commonDto.BaseResponseDTO{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Products Added",
	}, nil
}

func (s *productService) UpdateProduct(logger *zap.Logger, productID string, updateDTO *dtos.UpdateProductRequestDTO) (*commonDto.BaseResponseDTO, error) {
	if err := utils.ValidateRequest(logger, updateDTO, "Update request"); err != nil {
		return nil, err
	}

	id, err := utils.ValidateID(logger, productID, "Product")
	if err != nil {
		return nil, err
	}

	logger.Info("Finding product", zap.Uint("product_id", id))
	product, err := s.productRepo.FindByID(logger, id)
	if err != nil {
		logger.Error("Product not found",
			zap.Error(err),
			zap.Uint("product_id", id))
		return nil, errors.NewProductError(http.StatusNotFound, "Product not found")
	}

	logger.Info("Updating product fields", zap.Uint("product_id", id))
	if updateDTO.Price != nil {
		product.Price = *updateDTO.Price
	}
	if updateDTO.Stock != nil {
		product.StockQuantity = *updateDTO.Stock
	}

	if err := s.productRepo.Save(logger, product); err != nil {
		logger.Error("Failed to update product",
			zap.Error(err),
			zap.Uint("product_id", id))
		return nil, errors.NewProductError(http.StatusInternalServerError, "Failed to update product")
	}

	logger.Info("Product updated successfully", zap.Uint("product_id", id))
	return &commonDto.BaseResponseDTO{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Product Updated",
	}, nil
}

func (s *productService) UpdateProducts(logger *zap.Logger, updates map[string]dtos.UpdateProductRequestDTO) (*commonDto.BaseResponseDTO, error) {
	if len(updates) == 0 {
		logger.Error("Updates map cannot be empty")
		return nil, errors.NewProductError(http.StatusBadRequest, "Updates map cannot be empty")
	}

	logger.Info("Processing bulk update request", zap.Int("update_count", len(updates)))
	for productID, update := range updates {
		if _, err := s.UpdateProduct(logger, productID, &update); err != nil {
			return nil, err
		}
	}

	logger.Info("Products updated successfully", zap.Int("count", len(updates)))
	return &commonDto.BaseResponseDTO{
		Status:  true,
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%d products updated successfully", len(updates)),
	}, nil
}

func (s *productService) DeleteProduct(logger *zap.Logger, productID string, deactivate bool) (*commonDto.BaseResponseDTO, error) {
	id, err := utils.ValidateID(logger, productID, "Product")
	if err != nil {
		return nil, err
	}

	logger.Info("Processing delete request",
		zap.Uint("product_id", id),
		zap.Bool("deactivate", deactivate))

	var action func(*zap.Logger, uint) error
	if deactivate {
		action = s.productRepo.SoftDelete
	} else {
		action = s.productRepo.Delete
	}

	if err := action(logger, id); err != nil {
		actionType := "deactivate"
		if !deactivate {
			actionType = "delete"
		}
		logger.Error(fmt.Sprintf("Failed to %s product", actionType),
			zap.Error(err),
			zap.Uint("product_id", id))
		return nil, errors.NewProductError(http.StatusInternalServerError, fmt.Sprintf("Failed to %s product", actionType))
	}

	logger.Info("Product operation successful",
		zap.Uint("product_id", id),
		zap.Bool("deactivate", deactivate))
	return &commonDto.BaseResponseDTO{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Successfully Deleted",
	}, nil
}

func (s *productService) GetProduct(logger *zap.Logger, productID string) (*dtos.ProductResponseDTO, error) {
	id, err := utils.ValidateID(logger, productID, "Product")
	if err != nil {
		return nil, err
	}

	logger.Info("Finding product", zap.Uint("product_id", id))
	product, err := s.productRepo.FindByID(logger, id)
	if err != nil {
		logger.Error("Product not found",
			zap.Error(err),
			zap.Uint("product_id", id))
		return nil, errors.NewProductError(http.StatusNotFound, "Product not found")
	}

	productDTO := s.productUtils.MapProductDetails(logger, product)
	logger.Info("Product found", zap.Uint("product_id", id))

	return &dtos.ProductResponseDTO{
		BaseResponseDTO: &commonDto.BaseResponseDTO{
			Status: true,
			Code: http.StatusOK,
			Message: "Product Found",
		},
		Product: productDTO,
	}, nil
}

func (s *productService) GetProductList(logger *zap.Logger, filters *dtos.ProductFilterDTO) (*dtos.ProductResponseDTO, error) {
	logger.Info("Finding products with filters", zap.Any("filters", filters))
	products, err := s.productRepo.FindByFilters(logger, filters)
	if err != nil {
		logger.Error("Failed to retrieve products", zap.Error(err))
		return nil, errors.NewProductError(http.StatusInternalServerError, "Failed to retrieve products")
	}

	logger.Info("Mapping product details", zap.Int("product_count", len(products)))
	var productDTOs []*dtos.ProductDetailsDTO
	for _, product := range products {
		productDTO := s.productUtils.MapProductDetails(logger, &product)
		productDTOs = append(productDTOs, productDTO)
	}

	logger.Info("Products retrieved successfully", zap.Int("count", len(productDTOs)))
	return &dtos.ProductResponseDTO{
		BaseResponseDTO: &commonDto.BaseResponseDTO{
			Status:  true,
			Code:    http.StatusOK,
			Message: fmt.Sprintf("Product List : Entries %d", len(productDTOs)),
		},
		ProductList: productDTOs,
	}, nil
}

func (s *productService) GetProductSuggestions(logger *zap.Logger, query string) ([]string, error) {
	logger.Info("Getting product suggestions", zap.String("query", query))
	suggestions, err := s.productRepo.GetProductSuggestions(logger, query)
	if err != nil {
		logger.Error("Failed to get product suggestions",
			zap.Error(err),
			zap.String("query", query))
		return nil, err
	}

	logger.Info("Product suggestions retrieved",
		zap.String("query", query),
		zap.Int("suggestion_count", len(suggestions)))
	return suggestions, nil
}
