package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/controllers"

	"icecreambash/tgup_backend/internal/services"
)

// GetMe godoc
// @Summary      Проверка авторизации
// @Description  Получения claims и проверка авторизации
// @Tags         auth
// @Accept       json
// @Produce      json
// @Router       /api_gateway/auth/me [get]
// @Security ApiKeyAuth
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
