package repositories

import (
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"
	"gorm.io/gorm"
)

type TagRepository interface {
	BaseRepository[models.Tag, uint]
}

type tagRepository struct {
	BaseRepositoryImpl[models.Tag, uint]
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{
		BaseRepositoryImpl: BaseRepositoryImpl[models.Tag, uint]{
			db:     db,
		},
	}
}
