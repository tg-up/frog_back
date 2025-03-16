package controllers

import (
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/services"
	"net/http"
	"strconv"
)

type PlatformController struct {
	platformService services.PlatformService
}

func NewPlatformController(platformService services.PlatformService) *PlatformController {
	return &PlatformController{platformService: platformService}
}

func (platform *PlatformController) GetAllPlatform(ctx *gin.Context) {
	values, err := platform.platformService.GetAllPlatforms()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	ctx.JSON(200, gin.H{
		"data": values,
	})
}

func (platform *PlatformController) GetPlatformServices(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id not valid format",
		})
	}

	values, err := platform.platformService.GetPlatformServicesByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
	}

	ctx.JSON(200, gin.H{
		"data": values,
	})
}
