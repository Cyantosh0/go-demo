package controllers

import (
	"github/Cyantosh0/demo/models"
	"github/Cyantosh0/demo/repository"
	"github/Cyantosh0/demo/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	services.UserService
	repository.UserRepo
}

func NewUserController(UserService services.UserService, UserRepo repository.UserRepo) UserController {
	return UserController{UserService, UserRepo}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var inputUser models.User

	err := c.ShouldBindJSON(&inputUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = u.UserService.CreateUser(inputUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": "created"})
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	users, err := u.UserRepo.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": users})
}
