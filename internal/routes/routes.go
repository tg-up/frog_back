package routes

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/middlewares"
	"net/http"
)

func InjectRoutes(r *gin.Engine) {

	r.Use(
		requestid.New(
			requestid.WithCustomHeaderStrKey("x-ray-id"),
		),
	)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":  "Not found, please check docs",
			"x-ray-id": requestid.Get(ctx),
		})
	})

	r.Use(middlewares.CORSMiddleware())

	api := r.Group("api_gateway")
	{
		platforms := api.Group("platforms")
		InitPlatformRoutes(platforms)
	}
}
