package repositories

import (
	commonRepo "github.com/Kunal726/market-mosaic-common-lib-go/repository"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	commonRepo.BaseRepository[models.Category, uint]
}

type categoryRepository struct {
	commonRepo.BaseRepository[models.Category, uint]
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	baseRepo := commonRepo.NewBaseRepository[models.Category, uint](db)
	return &categoryRepository{
		baseRepo,
	}
}
