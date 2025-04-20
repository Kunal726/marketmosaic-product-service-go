package repositories

import (
	commonRepo "github.com/Kunal726/market-mosaic-common-lib-go/repository"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"
	"gorm.io/gorm"
)

type ProductImageRepository interface {
	commonRepo.BaseRepository[models.ProductImage, uint]
}

type productImageRepository struct {
	commonRepo.BaseRepository[models.ProductImage, uint]
}

func NewProductImageRepository(db *gorm.DB) ProductImageRepository {
	baseRepo := commonRepo.NewBaseRepository[models.ProductImage, uint](db)
	return &productImageRepository{
		baseRepo,
	}
}