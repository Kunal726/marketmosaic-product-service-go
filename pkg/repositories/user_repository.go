package repositories

import (
	"github.com/Kunal726/marketmosaic-product-service-go/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[models.User, uint]
}

type userRepository struct {
	BaseRepositoryImpl[models.User, uint]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepositoryImpl: BaseRepositoryImpl[models.User, uint]{
			db:     db,
		},
	}
}
