package repositories

import (
	commonRepo "github.com/Kunal726/market-mosaic-common-lib-go/repository"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"
	"gorm.io/gorm"
)

type TagRepository interface {
	commonRepo.BaseRepository[models.Tag, uint]
}

type tagRepository struct {
	commonRepo.BaseRepository[models.Tag, uint]
}

func NewTagRepository(db *gorm.DB) TagRepository {
	baseRepo := commonRepo.NewBaseRepository[models.Tag, uint](db)
	return &tagRepository{
		baseRepo,
	}
}
