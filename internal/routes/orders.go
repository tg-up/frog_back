package routes

import (
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/controllers"
	"icecreambash/tgup_backend/internal/repositories"
	"icecreambash/tgup_backend/internal/services"
	"icecreambash/tgup_backend/pkg/database"
)

func InitOrdersRoutes(router *gin.RouterGroup) {

	repositoryPlatform := repositories.NewPlatformRepository(database.DB)
	repositoryService := repositories.NewPlatformServiceRepository(database.DB)

	platformService := services.NewPlatformService(repositoryPlatform, repositoryService)

	order := controllers.NewOrderController(*platformService)

	router.GET("/", order.ShowAll)
	router.POST("/", order.Create)
}
