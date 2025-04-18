package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/api/handler"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/constants"
	"github.com/Kunal726/market-mosaic-common-lib-go/pkg/middleware/auth"
)

func RegisterRoutes(router *gin.Engine, handler *handler.ProductHandler, authMiddleware *auth.Middleware) {
	// Product routes that require authentication
	productsGroup := router.Group(constants.ProductBasePath)
	productsGroup.Use(authMiddleware.ValidateToken()) // Add auth middleware to all product routes
	{
		// Protected routes (require authentication)
		productsGroup.POST(constants.ProductRootPath, handler.AddProduct)           // POST /products
		productsGroup.POST(constants.ProductAddBulkPath, handler.AddProducts)       // POST /products/bulk
		productsGroup.PUT(constants.ProductSingleIdPath, handler.UpdateProduct)     // PUT /products/:productId
		productsGroup.POST(constants.ProductUpdateBulkPath, handler.UpdateProducts) // POST /products/bulk-update
		productsGroup.DELETE(constants.ProductSingleIdPath, handler.DeleteProduct)  // DELETE /products/:productId
	}

	// Public product routes (no authentication required)
	router.GET(constants.ProductBasePath+constants.ProductRootPath, handler.GetProductList)       // GET /products
	router.GET(constants.ProductBasePath+constants.ProductSingleIdPath, handler.GetProduct)      // GET /products/:productId
	router.GET(constants.ProductBasePath+constants.ProductSuggestionsPath, handler.GetProductSuggestions) // GET /products/suggestions
}
