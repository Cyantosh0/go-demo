package repository

import (
	"github/Cyantosh0/demo/infrastructure"
	"github/Cyantosh0/demo/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	*gorm.DB
}

func NewUserRepository(db infrastructure.Database) UserRepo {
	return UserRepo{db.DB}
}

func (ur *UserRepo) CreateUser(user models.User) error {
	return ur.Table("users").Save(&user).Error
}

func (ur *UserRepo) GetAllUsers() (users []models.User, err error) {
	return users, ur.Find(&users).Error
}
