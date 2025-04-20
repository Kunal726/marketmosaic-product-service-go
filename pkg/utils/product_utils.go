package utils

import (
	"fmt"
	"time"

	"github.com/Kunal726/marketmosaic-product-service-go/pkg/dtos"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/repositories"

	"go.uber.org/zap"
)

type ProductUtils struct {
	userRepo     repositories.UserRepository
	categoryRepo repositories.CategoryRepository
	tagRepo      repositories.TagRepository
}

func NewProductUtils(
	userRepo repositories.UserRepository,
	categoryRepo repositories.CategoryRepository,
	tagRepo repositories.TagRepository,
) *ProductUtils {
	return &ProductUtils{
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

// MapProductDetails maps a Product entity to ProductDetailsDTO
func (p *ProductUtils) MapProductDetails(logger *zap.Logger, product *models.Product) *dtos.ProductDetailsDTO {
	if product == nil {
		logger.Warn("Attempted to map nil product")
		return nil
	}

	logger.Debug("Mapping product to DTO",
		zap.Uint("product_id", product.ProductId))

	dto := &dtos.ProductDetailsDTO{
		ProductID:     fmt.Sprintf("%d", product.ProductId),
		ProductName:   product.ProductName,
		CategoryID:    fmt.Sprintf("%d", product.CategoryId),
		Price:         product.Price,
		Description:   product.Description,
		IsActive:      &product.IsActive,
		StockQuantity: product.StockQuantity,
		SupplierID:    fmt.Sprintf("%d", product.SupplierId),
	}

	// Map tag IDs if tags exist
	if product.Tags != nil {
		tagIds := make([]string, 0, len(product.Tags))
		for _, tag := range product.Tags {
			tagIds = append(tagIds, fmt.Sprintf("%d", tag.TagId))
		}
		dto.TagIds = tagIds
		logger.Debug("Mapped product tags",
			zap.Uint("product_id", product.ProductId),
			zap.Strings("tag_ids", tagIds))
	}

	logger.Debug("Successfully mapped product to DTO",
		zap.Uint("product_id", product.ProductId))
	return dto
}

// MapProductEntity maps a ProductDetailsDTO to Product entity
func (p *ProductUtils) MapProductEntity(logger *zap.Logger, dto *dtos.ProductDetailsDTO) (*models.Product, error) {
	if dto == nil {
		logger.Error("Product details DTO is nil")
		return nil, fmt.Errorf("product details DTO is nil")
	}

	logger.Debug("Mapping DTO to product entity",
		zap.String("product_name", dto.ProductName))

	// Handle IsActive field
	var isActive bool
	if dto.IsActive != nil {
		isActive = *dto.IsActive
	} else {
		isActive = true // default to true if not specified
	}

	// Parse IDs from string to uint
	supplierID, err := parseUint(dto.SupplierID)
	if err != nil {
		logger.Error("Invalid supplier ID",
			zap.Error(err),
			zap.String("supplier_id", dto.SupplierID))
		return nil, fmt.Errorf("invalid supplier ID: %v", err)
	}

	categoryID, err := parseUint(dto.CategoryID)
	if err != nil {
		logger.Error("Invalid category ID",
			zap.Error(err),
			zap.String("category_id", dto.CategoryID))
		return nil, fmt.Errorf("invalid category ID: %v", err)
	}

	// Verify supplier exists
	logger.Debug("Verifying supplier exists",
		zap.Uint("supplier_id", supplierID))
	_, err = p.userRepo.FindByID(logger, supplierID)
	if err != nil {
		logger.Error("Supplier not found",
			zap.Error(err),
			zap.Uint("supplier_id", supplierID))
		return nil, fmt.Errorf("supplier not found: %v", err)
	}

	// Verify category exists
	logger.Debug("Verifying category exists",
		zap.Uint("category_id", categoryID))
	_, err = p.categoryRepo.FindByID(logger, categoryID)
	if err != nil {
		logger.Error("Category not found",
			zap.Error(err),
			zap.Uint("category_id", categoryID))
		return nil, fmt.Errorf("category not found: %v", err)
	}

	product := &models.Product{
		ProductName:   dto.ProductName,
		CategoryId:    categoryID,
		Price:         dto.Price,
		Description:   dto.Description,
		DateAdded:     time.Now(),
		IsActive:      isActive,
		StockQuantity: dto.StockQuantity,
		SupplierId:    supplierID,
	}

	// Map tags if tag IDs exist
	if dto.TagIds != nil && len(dto.TagIds) > 0 {
		logger.Debug("Mapping product tags",
			zap.Strings("tag_ids", dto.TagIds))

		var tags []models.Tag
		for _, tagIDStr := range dto.TagIds {
			tagID, err := parseUint(tagIDStr)
			if err != nil {
				logger.Warn("Invalid tag ID, skipping",
					zap.Error(err),
					zap.String("tag_id", tagIDStr))
				continue
			}

			_, err = p.tagRepo.FindByID(logger, tagID)
			if err != nil {
				logger.Warn("Tag not found, skipping",
					zap.Error(err),
					zap.Uint("tag_id", tagID))
				continue
			}

			tags = append(tags, models.Tag{TagId: tagID})
		}
		product.Tags = tags
		logger.Debug("Successfully mapped tags",
			zap.Int("tag_count", len(tags)))
	}

	logger.Debug("Successfully mapped DTO to product entity",
		zap.String("product_name", product.ProductName))
	return product, nil
}

// Helper function to parse string to uint
func parseUint(s string) (uint, error) {
	var result uint
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}
