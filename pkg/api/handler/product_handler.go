package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/dtos"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/errors"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/services"
	"github.com/Kunal726/market-mosaic-common-lib-go/pkg/utils"
	"net/http"

	"github.com/Kunal726/market-mosaic-common-lib-go/pkg/middleware/auth"
	"go.uber.org/zap"
)

// ProductHandler handles product-related HTTP requests
type ProductHandler struct {
	productService services.ProductService
}

// NewProductHandler creates a new ProductHandler instance
func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// AddProduct handles the creation of a single product
// POST /products
func (h *ProductHandler) AddProduct(c *gin.Context) {
	logger := utils.GetLoggerFromContext(c)
	logger.Info("Received add product request")

	// Get user info from context
	_, exists := auth.GetUserFromContext(c)
	if !exists {
		logger.Error("User context not found")
		c.Error(errors.NewAuthError("Unauthorized access"))
		return
	}

	var productDTO dtos.ProductDetailsDTO
	if err := c.ShouldBindJSON(&productDTO); err != nil {
		logger.Error("Failed to bind request body", zap.Error(err))
		c.Error(err)
		return
	}

	logger.Debug("Product details", zap.Any("product", productDTO))
	response, err := h.productService.AddProduct(logger, &productDTO)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// AddProducts handles bulk product creation
// POST /products/bulk
func (h *ProductHandler) AddProducts(c *gin.Context) {
	logger := utils.GetLoggerFromContext(c)
	logger.Info("Received add products request")

	// Get user info from context
	_, exists := auth.GetUserFromContext(c)
	if !exists {
		logger.Error("User context not found")
		c.Error(errors.NewAuthError("Unauthorized access"))
		return
	}

	var products []dtos.ProductDetailsDTO
	if err := c.ShouldBindJSON(&products); err != nil {
		logger.Error("Failed to bind request body",
			zap.Error(err))
		c.Error(err)
		return
	}

	logger.Debug("Products to add",
		zap.Int("count", len(products)))

	response, err := h.productService.AddProducts(logger, products)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProduct handles updating a single product
// PUT /products/:productId
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	logger := utils.GetLoggerFromContext(c)
	logger.Info("Received update product request")

	// Get user info from context
	_, exists := auth.GetUserFromContext(c)
	if !exists {
		logger.Error("User context not found")
		c.Error(errors.NewAuthError("Unauthorized access"))
		return
	}

	productID := c.Param("productId")
	var updateDTO dtos.UpdateProductRequestDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		logger.Error("Failed to bind request body", zap.Error(err))
		c.Error(err)
		return
	}

	logger.Debug("Update details", 
		zap.String("product_id", productID),
		zap.Any("update_data", updateDTO))

	response, err := h.productService.UpdateProduct(logger, productID, &updateDTO)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProducts handles bulk product updates
// POST /products/bulk-update
func (h *ProductHandler) UpdateProducts(c *gin.Context) {
	logger := utils.GetLoggerFromContext(c)
	logger.Info("Received bulk update products request")

	// Get user info from context
	_, exists := auth.GetUserFromContext(c)
	if !exists {
		logger.Error("User context not found")
		c.Error(errors.NewAuthError("Unauthorized access"))
		return
	}

	var updates map[string]dtos.UpdateProductRequestDTO
	if err := c.ShouldBindJSON(&updates); err != nil {
		logger.Error("Failed to bind request body",
			zap.Error(err))
		c.Error(err)
		return
	}

	logger.Debug("Updates to apply",
		zap.Int("count", len(updates)))

	response, err := h.productService.UpdateProducts(logger, updates)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteProduct handles deleting a single product
// DELETE /products/:productId
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	logger := utils.GetLoggerFromContext(c)
	
	// Get user info from context
	_, exists := auth.GetUserFromContext(c)
	if !exists {
		logger.Error("User context not found")
		c.Error(errors.NewAuthError("Unauthorized access"))
		return
	}

	productID := c.Param("productId")
	deactivate := c.DefaultQuery("deactivate", "false") == "true"

	logger.Info("Received delete product request",
		zap.String("product_id", productID),
		zap.Bool("deactivate", deactivate))

	response, err := h.productService.DeleteProduct(logger, productID, deactivate)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetProduct handles retrieving a single product
// GET /products/:productId
func (h *ProductHandler) GetProduct(c *gin.Context) {
	logger := utils.GetLoggerFromContext(c)
	productID := c.Param("productId")

	logger.Info("Received get product request",
		zap.String("product_id", productID))

	response, err := h.productService.GetProduct(logger, productID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetProductList handles retrieving filtered products
// GET /products
func (h *ProductHandler) GetProductList(c *gin.Context) {
	logger := utils.GetLoggerFromContext(c)
	logger.Info("Received product list request",
		zap.String("raw_query", c.Request.URL.RawQuery))

	var filter dtos.ProductFilterDTO
	if err := c.ShouldBindQuery(&filter); err != nil {
		logger.Error("Failed to bind query parameters",
			zap.Error(err))
		c.Error(err)
		return
	}

	logger.Debug("Filter values",
		zap.Any("filter", filter))

	response, err := h.productService.GetProductList(logger, &filter)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetSuggestions handles product name suggestions
// GET /products/suggestions
func (h *ProductHandler) GetProductSuggestions(c *gin.Context) {
	logger := utils.GetLoggerFromContext(c)
	query := c.Query("q")

	logger.Info("Received product suggestions request",
		zap.String("query", query))

	if query == "" {
		logger.Error("Query parameter 'q' is required")
		c.Error(errors.NewProductError(http.StatusBadRequest, "Query parameter 'q' is required"))
		return
	}

	suggestions, err := h.productService.GetProductSuggestions(logger, query)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, suggestions)
}
