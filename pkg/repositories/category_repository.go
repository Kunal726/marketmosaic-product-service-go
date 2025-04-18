package repositories

import (
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	BaseRepository[models.Category, uint]
}

type categoryRepository struct {
	BaseRepositoryImpl[models.Category, uint]
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		BaseRepositoryImpl: BaseRepositoryImpl[models.Category, uint]{
			db:     db,
		},
	}
}
