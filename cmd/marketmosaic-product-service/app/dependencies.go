package app

import (
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/api/handler"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/repositories"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/services"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/utils"

	"github.com/Kunal726/market-mosaic-common-lib-go/pkg/redis"
	"gorm.io/gorm"
)

// Repositories holds all repository instances
type Repositories struct {
	User     repositories.UserRepository
	Product  repositories.ProductRepository
	Category repositories.CategoryRepository
	Tag      repositories.TagRepository
}

// Services holds all service instances
type Services struct {
	Product services.ProductService
}

// Handlers holds all handler instances
type Handlers struct {
	Product *handler.ProductHandler
}	

// NewRepositories initializes all repositories
func NewRepositories(db *gorm.DB, redisManager *redis.Manager) *Repositories {
	return &Repositories{
		User:     repositories.NewUserRepository(db),
		Product:  repositories.NewProductRepository(db, redisManager),
		Category: repositories.NewCategoryRepository(db),
		Tag:      repositories.NewTagRepository(db),
	}
}

// NewServices initializes all services
func NewServices(repos *Repositories) *Services {
	productUtils := utils.NewProductUtils(repos.User, repos.Category, repos.Tag)

	return &Services{
		Product: services.NewProductService(repos.Product, repos.User, productUtils),
	}
}

// NewHandlers initializes all handlers
func NewHandlers(services *Services) *Handlers {
	return &Handlers{
		Product: handler.NewProductHandler(services.Product),
	}
}
