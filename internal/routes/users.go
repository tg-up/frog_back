package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/controllers"

	"icecreambash/tgup_backend/internal/services"
)

func InitUsersRoutes(router *gin.RouterGroup) {

	userService := services.NewUserService()

	userController := controllers.NewUserController(*userService)

	auth := router.Group("auth")

	auth.POST("login", userController.Login)
	auth.POST("register", userController.Register)

	props := auth.Use(JWTAuthMiddleware.MiddlewareFunc())

	props.GET("me", func(context *gin.Context) {
		claims := jwt.ExtractClaims(context)
		context.JSON(200, claims)
	})
}
