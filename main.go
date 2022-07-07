package main

import (
	"github/Cyantosh0/demo/api/controllers"
	"github/Cyantosh0/demo/api/routes"
	"github/Cyantosh0/demo/infrastructure"
	"github/Cyantosh0/demo/lib"
	"github/Cyantosh0/demo/repository"
	"github/Cyantosh0/demo/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			lib.Module,
			infrastructure.Module,
			routes.Module,
			controllers.Module,
			services.Module,
			repository.Module,
		),
		fx.Invoke(startApp),
	).Run()
}

func startApp(
	env *lib.Env,
	r *gin.Engine,
	routes routes.Routes,
	migrator infrastructure.Migrations,
) {
	migrator.Migrate()
	routes.SetupRoutes()
	r.Run(env.SeverPort)
}
