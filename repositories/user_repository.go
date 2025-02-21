package repositories

import (
	"sample-project/internal/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}
