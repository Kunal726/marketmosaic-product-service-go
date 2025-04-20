package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Kunal726/marketmosaic-product-service-go/pkg/dtos"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"

	"github.com/Kunal726/market-mosaic-common-lib-go/pkg/redis"
	commonRepo "github.com/Kunal726/market-mosaic-common-lib-go/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository interface {
	commonRepo.BaseRepository[models.Product, uint]
	SoftDelete(logger *zap.Logger, productId uint) error
	GetProductSuggestions(logger *zap.Logger, query string) ([]string, error)
	FindByFilters(logger *zap.Logger, filter *dtos.ProductFilterDTO) ([]models.Product, error)
}

type productRepository struct {
	commonRepo.BaseRepository[models.Product, uint]
	redisManager *redis.Manager
	db           *gorm.DB
}

func NewProductRepository(db *gorm.DB, redisManager *redis.Manager) ProductRepository {
	baseRepo := commonRepo.NewBaseRepository[models.Product, uint](db)
	return &productRepository{
		BaseRepository: baseRepo,
		redisManager:   redisManager,
		db:             db,
	}
}

// Helper function to generate cache key
func (r *productRepository) getCacheKey(prefix string, id interface{}) string {
	return fmt.Sprintf("product:%s:%v", prefix, id)
}

func (r *productRepository) FindByID(logger *zap.Logger, id uint) (*models.Product, error) {
	// Try to get from cache first
	cacheKey := r.getCacheKey("id", id)
	var product models.Product
	err := r.redisManager.Get(context.Background(), cacheKey, &product)
	if err == nil {
		logger.Debug("Product retrieved from cache",
			zap.Uint("product_id", id))
		return &product, nil
	}

	// If not in cache, get from database
	logger.Info("Finding product by ID", zap.Uint("product_id", id))
	err = r.db.Preload("Tags").
		Preload("Category").
		Preload("Supplier").
		Preload("Images").
		First(&product, id).Error
	if err != nil {
		logger.Error("Failed to find product by ID",
			zap.Error(err),
			zap.Uint("product_id", id))
		return nil, err
	}

	// Cache the result
	r.redisManager.Set(context.Background(), cacheKey, product, 5*time.Minute)

	logger.Info("Successfully found product by ID", zap.Uint("product_id", id))
	return &product, nil
}

func (r *productRepository) Save(logger *zap.Logger, product *models.Product) error {
	// Save to database
	if err := r.db.Save(product).Error; err != nil {
		logger.Error("Failed to save product", zap.Error(err))
		return err
	}

	// Update cache
	cacheKey := r.getCacheKey("id", product.ProductId)
	r.redisManager.Set(context.Background(), cacheKey, product, 5*time.Minute)

	// Invalidate filter cache
	r.redisManager.Delete(context.Background(), "product:filter:*")

	logger.Info("Successfully saved product", zap.Uint("product_id", product.ProductId))
	return nil
}

func (r *productRepository) Delete(logger *zap.Logger, id uint) error {
	logger.Info("Deleting product and related records",
		zap.Uint("product_id", id))

	return r.db.Transaction(func(tx *gorm.DB) error {
		// First delete related records from product_tags
		logger.Debug("Deleting product tags",
			zap.Uint("product_id", id))
		if err := tx.Exec("DELETE FROM product_tags WHERE product_id = ?", id).Error; err != nil {
			logger.Error("Failed to delete product tags",
				zap.Error(err),
				zap.Uint("product_id", id))
			return err
		}

		// Then delete the product
		logger.Debug("Deleting product",
			zap.Uint("product_id", id))
		if err := tx.Delete(&models.Product{}, "product_id = ?", id).Error; err != nil {
			logger.Error("Failed to delete product",
				zap.Error(err),
				zap.Uint("product_id", id))
			return err
		}

		// Invalidate cache
		r.redisManager.Delete(context.Background(), r.getCacheKey("id", id))
		r.redisManager.Delete(context.Background(), "product:filter:*")

		logger.Info("Successfully deleted product and related records",
			zap.Uint("product_id", id))
		return nil
	})
}

func (r *productRepository) SoftDelete(logger *zap.Logger, productId uint) error {
	logger.Info("Soft deleting product",
		zap.Uint("product_id", productId))

	err := r.db.Model(&models.Product{}).
		Where("product_id = ?", productId).
		Update("is_active", false).
		Error

	if err != nil {
		logger.Error("Failed to soft delete product",
			zap.Error(err),
			zap.Uint("product_id", productId))
		return err
	}

	logger.Info("Successfully soft deleted product",
		zap.Uint("product_id", productId))
	return nil
}

func (r *productRepository) GetProductSuggestions(logger *zap.Logger, query string) ([]string, error) {
	logger.Info("Getting product suggestions",
		zap.String("query", query))

	var suggestions []string
	err := r.db.Model(&models.Product{}).
		Select("DISTINCT product_name").
		Where("LOWER(product_name) LIKE LOWER(?)", "%"+query+"%").
		Find(&suggestions).
		Error

	if err != nil {
		logger.Error("Failed to get product suggestions",
			zap.Error(err),
			zap.String("query", query))
		return nil, err
	}

	logger.Info("Successfully retrieved product suggestions",
		zap.String("query", query),
		zap.Int("suggestion_count", len(suggestions)))
	return suggestions, err
}

// Helper function to parse string to uint
func parseUint(s string) (uint, error) {
	var result uint
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}

func (r *productRepository) FindByFilters(logger *zap.Logger, filter *dtos.ProductFilterDTO) ([]models.Product, error) {
	// Generate cache key based on filter
	filterKey := fmt.Sprintf("%v", filter)
	cacheKey := r.getCacheKey("filter", filterKey)

	// Try to get from cache first
	var products []models.Product
	err := r.redisManager.Get(context.Background(), cacheKey, &products)
	if err == nil {
		logger.Debug("Products retrieved from cache",
			zap.Any("filter", filter))
		return products, nil
	}

	logger.Info("Finding products by filters",
		zap.Any("filter", filter))

	query := r.db.Model(&models.Product{})

	if filter.CategoryId != "" {
		// First get the category and all its subcategories
		var categoryIds []uint
		categoryId, err := parseUint(filter.CategoryId)
		if err != nil {
			logger.Error("Invalid category ID",
				zap.Error(err),
				zap.String("category_id", filter.CategoryId))
			return nil, fmt.Errorf("invalid category ID: %v", err)
		}

		logger.Debug("Getting category hierarchy",
			zap.Uint("category_id", categoryId))

		// Get all subcategories recursively using a CTE (Common Table Expression)
		subQuery := `
			WITH RECURSIVE category_tree AS (
				SELECT category_id, parent_id, category_name
				FROM marketmosaic_category
				WHERE category_id = ?
				UNION ALL
				SELECT c.category_id, c.parent_id, c.category_name
				FROM marketmosaic_category c
				INNER JOIN category_tree ct ON c.parent_id = ct.category_id
			)
			SELECT category_id FROM category_tree;
		`
		if err := r.db.Raw(subQuery, categoryId).Pluck("category_id", &categoryIds).Error; err != nil {
			logger.Error("Failed to get category hierarchy",
				zap.Error(err),
				zap.Uint("category_id", categoryId))
			return nil, fmt.Errorf("failed to get category hierarchy: %v", err)
		}

		logger.Debug("Found category IDs in hierarchy",
			zap.Uint("parent_category_id", categoryId),
			zap.Any("category_ids", categoryIds))

		if len(categoryIds) > 0 {
			query = query.Where("category_id IN ?", categoryIds)
		} else {
			query = query.Where("category_id = ?", categoryId)
		}
	}

	// Handle price range filtering
	if filter.MinPrice != nil && filter.MaxPrice != nil {
		logger.Debug("Applying price range filter",
			zap.Float64("min_price", *filter.MinPrice),
			zap.Float64("max_price", *filter.MaxPrice))
		query = query.Where("price BETWEEN ? AND ?", filter.MinPrice, filter.MaxPrice)
	} else if filter.MinPrice != nil {
		logger.Debug("Applying minimum price filter",
			zap.Float64("min_price", *filter.MinPrice))
		query = query.Where("price >= ?", filter.MinPrice)
	} else if filter.MaxPrice != nil {
		logger.Debug("Applying maximum price filter",
			zap.Float64("max_price", *filter.MaxPrice))
		query = query.Where("price <= ?", filter.MaxPrice)
	}

	if filter.SearchTerm != "" {
		logger.Debug("Applying search term filter",
			zap.String("search_term", filter.SearchTerm))
		searchTerm := "%" + strings.ToLower(filter.SearchTerm) + "%"
		query = query.Where("LOWER(product_name) LIKE ? OR LOWER(description) LIKE ?", searchTerm, searchTerm)
	}

	if filter.SupplierId != "" {
		logger.Debug("Applying supplier filter",
			zap.String("supplier_id", filter.SupplierId))
		query = query.Where("supplier_id = ?", filter.SupplierId)
	}

	if filter.Tags != nil && len(filter.Tags) > 0 {
		logger.Debug("Applying tag filter",
			zap.Strings("tags", filter.Tags))
		query = query.Joins("JOIN product_tags pt ON marketmosaic_product.product_id = pt.product_id").
			Joins("JOIN marketmosaic_tag t ON pt.tag_id = t.tag_id").
			Where("t.tag_id IN ?", filter.Tags)
	}

	// Add preloading for relationships
	query = query.Preload("Tags").
		Preload("Category").
		Preload("Supplier")

	if err := query.Find(&products).Error; err != nil {
		logger.Error("Failed to find products by filters",
			zap.Error(err),
			zap.Any("filter", filter))
		return nil, err
	}

	// Cache the result
	r.redisManager.Set(context.Background(), cacheKey, products, 5*time.Minute)

	logger.Info("Successfully found products by filters",
		zap.Int("product_count", len(products)))
	return products, nil
}
