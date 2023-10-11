package database

import (
	"github.com/allanmaral/go-expert/09-apis/internal/entity"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

func (ur *GormUserRepository) Create(user *entity.User) error {
	return ur.DB.Create(user).Error
}

func (ur *GormUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := ur.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
