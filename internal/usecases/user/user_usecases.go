package user

import (
	"sample-project/internal/entities"
	"sample-project/repositories"
)

type UserUseCase struct {
	UserRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: userRepo}
}

func (uc *UserUseCase) CreateUser(user *entities.User) error {
	return uc.UserRepo.Create(user)
}
