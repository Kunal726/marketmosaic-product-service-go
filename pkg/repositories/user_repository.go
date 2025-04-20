package repositories

import (
	commonRepo "github.com/Kunal726/market-mosaic-common-lib-go/repository"
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	commonRepo.BaseRepository[models.User, uint]
}

type userRepository struct {
	commonRepo.BaseRepository[models.User, uint]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	baseRepo := commonRepo.NewBaseRepository[models.User, uint](db)
	return &userRepository{
		baseRepo,
	}
}
