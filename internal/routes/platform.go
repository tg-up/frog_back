package routes

import (
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/controllers"
	"icecreambash/tgup_backend/internal/repositories"
	"icecreambash/tgup_backend/internal/services"
	"icecreambash/tgup_backend/pkg/database"
)

func InitPlatformRoutes(router *gin.RouterGroup) {
	//Load Repositories
	platformRepository := repositories.NewPlatformRepository(database.DB)
	platformServiceRepository := repositories.NewPlatformServiceRepository(database.DB)
	//Load Services
	platformService := services.NewPlatformService(platformRepository, platformServiceRepository)
	//LoadPlatforms
	platformController := controllers.NewPlatformController(*platformService)

	router.GET("/", platformController.GetAllPlatform)

	platformServices := router.Group(":id/services")

	platformServices.GET("/", platformController.GetPlatformServices)

}
