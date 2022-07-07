package services

import (
	"github/Cyantosh0/demo/models"
	"github/Cyantosh0/demo/repository"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(UserRepo repository.UserRepo) UserService {
	return UserService{userRepo: UserRepo}
}

func (u *UserService) CreateUser(user models.User) error {
	return u.userRepo.CreateUser(user)
}
