package routes

import (
	"github/Cyantosh0/demo/api/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Routes struct {
	*gin.Engine
	controllers.UserController
}

func NewRoutes(router *gin.Engine, UserController controllers.UserController) Routes {
	return Routes{router, UserController}
}

func (r *Routes) SetupRoutes() {
	r.POST("/user", r.UserController.CreateUser)

	r.GET("/users", r.UserController.GetAllUsers)
}

var Module = fx.Options(
	fx.Provide(NewRoutes),
)
